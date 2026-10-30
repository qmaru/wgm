package ui

import (
	"slices"
	"strconv"
	"strings"

	myui "github.com/egoist/mygo/ui"
	qrcode "github.com/skip2/go-qrcode"

	"wgm/backend"
)

var configSteps = [3]string{"选择 Interface", "选择 Peers", "确认配置"}

func (a *App) configsPage(c *myui.Context) {
	pageBody(c, func() {
		pageHeader(c, "CONFIGURATION BUILDER", "配置", "按步骤生成 WireGuard 配置并导出", nil)
		a.stepper(c)
		a.stepSummary(c)
		a.stepNav(c)
		switch a.step {
		case 0:
			a.stepInterface(c)
		case 1:
			a.stepPeers(c)
		case 2:
			pageScroll(c, func() { a.stepResult(c) })
		}
	})
}

func (a *App) stepper(c *myui.Context) {
	t := c.Theme()
	myui.Row(c).Gap(4).Margin(0, 0, 10, 0).Children(func() {
		for i, label := range configSteps {
			cell := myui.Column(c).Grow(1).AlignItems(myui.Center).Gap(2).Padding(8).Radius(7)
			switch {
			case i == a.step:
				cell.Background(t.Accent.Alpha(0.14))
			case i < a.step:
				cell.Background(t.Surface)
			}
			cell.Children(func() {
				myui.Textf(c, "0%d", i+1).Font("monospace").FontSize(10).TextColor(t.TextMuted)
				text := myui.Text(c, label).FontSize(12).SingleLine()
				if i == a.step {
					text.Bold().TextColor(t.Accent)
				} else if i < a.step {
					text.TextColor(t.Success)
				}
			})
		}
	})
}

func (a *App) stepSummary(c *myui.Context) {
	t := c.Theme()
	iface := "未选择"
	if p := a.peerByKey(a.iface); p != nil {
		iface = p.Username + " · " + p.PrivateAddr
	}
	peers := a.selectedPeersSummary()

	title := "配置已生成"
	switch a.step {
	case 0:
		title = "选择 Interface"
	case 1:
		title = "选择 Peers 和路由"
	}
	myui.Column(c).Gap(4).MinWidth(0).Padding(10).Radius(8).Border(1, t.Border).Background(t.Surface).
		Margin(0, 0, 10, 0).Children(func() {
		myui.Text(c, title).FontSize(13).Bold()
		myui.Row(c).Gap(6).MinWidth(0).Children(func() {
			myui.Text(c, "Interface").FontSize(12).TextColor(t.TextMuted).Width(58)
			myui.Text(c, iface).FontSize(12).Wrap().MinWidth(0)
		})
		myui.Row(c).Gap(6).MinWidth(0).Children(func() {
			myui.Text(c, "Peers").FontSize(12).TextColor(t.TextMuted).Width(58)
			myui.Text(c, peers).FontSize(12).Wrap().MinWidth(0)
		})
	})
}

func (a *App) stepNav(c *myui.Context) {
	myui.Row(c).Justify(myui.SpaceBetween).Margin(0, 0, 12, 0).Children(func() {
		if myui.Button(c, "上一步").Disabled(a.step == 0).Clicked() {
			a.step--
		}
		label := "下一步"
		if a.step > 0 {
			label = "完成"
		}
		if myui.PrimaryButton(c, label).Clicked() {
			switch a.step {
			case 0:
				if a.iface == "" {
					c.Toast("请选择 Interface")
					return
				}
				a.step = 1
			case 1:
				if msg := a.buildConfig(); msg != "" {
					c.Toast(msg)
				}
			case 2:
				a.resetConfig()
			}
		}
	})
}

func (a *App) stepInterface(c *myui.Context) {
	myui.Column(c).Gap(6).Grow(1).MinWidth(0).MinHeight(0).Children(func() {
		myui.Text(c, "选择本地接口 Interface").FontSize(13).Bold()
		myui.GridView(c, &a.interfaceGrid, len(a.peers), 260, 96, func(i int) {
			a.interfaceChoice(c, &a.peers[i])
		}).Grow(1).MinWidth(0).MinHeight(0).Gap(4).
			ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
			Padding(0, gridRightInset, gridBottomInset, 0)
	})
}

func (a *App) interfaceChoice(c *myui.Context, p *backend.Peer) {
	t := c.Theme()
	key := peerKey(p.ID)
	selectInterface := func() {
		if a.iface == key {
			return
		}
		a.iface = key
		a.clearPeerSelection(key)
	}
	card := myui.Column(c).Fill().Gap(8).MinWidth(0).Padding(12).Radius(8).Border(1, t.Border).Background(t.Surface).
		OnClick(selectInterface)
	card.Children(func() {
		if myui.Radio(c, &a.iface, key, p.Username).Changed() {
			selectInterface()
		}
		myui.Row(c).Gap(8).AlignItems(myui.Center).MinWidth(0).Children(func() {
			myui.Text(c, "LAN").FontSize(11).TextColor(t.TextMuted).Width(32)
			myui.Text(c, p.PrivateAddr).Font("monospace").FontSize(12).Wrap().MinWidth(0)
		})
		if p.PublicAddr != "" {
			myui.Row(c).Gap(8).AlignItems(myui.Center).MinWidth(0).Children(func() {
				myui.Text(c, "WAN").FontSize(11).TextColor(t.TextMuted).Width(32)
				myui.Text(c, endpoint(p.PublicAddr, p.Port)).Font("monospace").FontSize(12).Wrap().MinWidth(0)
			})
		}
	})
	background, border := t.Surface, t.Border
	if a.iface == key {
		background = t.Accent.Alpha(0.12)
		border = t.Accent
	} else if card.Hovered() {
		background = t.SurfaceHover
	}
	card.Background(background).BorderColor(border).Transition(cardStateTransition)
}

func (a *App) stepPeers(c *myui.Context) {
	peerCount := len(a.peers)
	if a.peerByKey(a.iface) != nil {
		peerCount--
	}
	myui.Column(c).Gap(8).Grow(1).MinWidth(0).MinHeight(0).Children(func() {
		myui.Text(c, "选择对端 Peer（可多选）").FontSize(13).Bold()
		myui.GridView(c, &a.wizardGrid, peerCount, 280, 172, func(i int) {
			a.peerChoice(c, a.peerWithoutInterface(i))
		}).Grow(1).MinWidth(0).MinHeight(0).Gap(8).
			ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
			Padding(0, gridRightInset, gridBottomInset, 0)
	})
}

func (a *App) peerChoice(c *myui.Context, p *backend.Peer) {
	t := c.Theme()
	key := peerKey(p.ID)
	checked := a.selected[key]
	selectPeer := func() {
		if checked {
			return
		}
		checked = true
		a.togglePeer(key, true)
	}
	card := myui.Column(c).Gap(6).MinWidth(0).Padding(10).Radius(8).Border(1, t.Border).Background(t.Surface)
	card.Children(func() {
		myui.Column(c).Gap(6).MinWidth(0).OnClick(selectPeer).Children(func() {
			if myui.Checkbox(c, &checked, p.Username).Changed() {
				a.togglePeer(key, checked)
			}
			myui.Text(c, p.PrivateAddr).Font("monospace").FontSize(12).TextColor(t.TextMuted).Wrap().MinWidth(0)
		})
		myui.Column(c).Gap(6).MinWidth(0).Children(func() {
			myui.Divider(c)
			myui.Text(c, "路由（可多选）").FontSize(12).TextColor(t.TextMuted)
			myui.Row(c).Gap(4).Wrap().MinWidth(0).Children(func() {
				if len(a.routes) == 0 {
					myui.Text(c, "暂无可选路由").FontSize(12).TextColor(t.TextMuted)
					return
				}
				for _, r := range a.routes {
					on := slices.Contains(a.routeMap[key], r.CIDR)
					b := myui.Button(c, r.CIDR).FontSize(12).Padding(3, 7).Disabled(!checked)
					if on {
						b.Background(t.Accent).TextColor(t.AccentText)
					}
					if b.Clicked() {
						a.toggleRoute(key, r.CIDR, on)
					}
				}
			})
		})
	})
	background, border := t.Surface, t.Border
	if checked {
		background = t.Accent.Alpha(0.12)
		border = t.Accent
	} else if card.Hovered() {
		background = t.SurfaceHover
	}
	card.Background(background).BorderColor(border).Transition(cardStateTransition)
}

func (a *App) stepResult(c *myui.Context) {
	myui.Column(c).Gap(8).MinWidth(0).Children(func() {
		for i, block := range a.blocks {
			a.resultCard(c, block, "block-"+strconv.Itoa(i), "配置块", nil)
		}
		a.resultCard(c, a.config, "all", "完整配置", func() {
			if a.qr != nil {
				myui.Image(c, a.qr).Size(configQRDisplaySize, configQRDisplaySize)
			}
		})
	})
}

// resultCard is a block of the generated config: a click copies it.
func (a *App) resultCard(c *myui.Context, text, key, label string, footer func()) {
	t := c.Theme()
	hint := "点击此处复制此" + label
	if a.copied == key {
		hint += " · 已复制"
	}
	card := cardShell(c, false, func() {
		myui.Text(c, label).FontSize(13).Bold().TextColor(t.Accent)
	}, func() {
		buildCode := func() {
			myui.Column(c).Gap(2).MinWidth(0).Children(func() {
				for _, line := range strings.Split(text, "\n") {
					if strings.TrimSpace(line) == "" {
						line = " "
					}
					myui.Text(c, line).Font("monospace").FontSize(13).TextColor(t.Text).Wrap().MinWidth(0)
				}
			})
		}
		if footer == nil {
			buildCode()
			return
		}
		myui.Row(c).AlignItems(myui.Start).Gap(12).MinWidth(0).Children(func() {
			myui.Column(c).Grow(1).MinWidth(0).Children(buildCode)
			footer()
		})
	}, func() {
		myui.Text(c, hint).FontSize(12).TextColor(t.TextMuted)
	})
	if a.copied == key {
		card.BorderColor(t.Accent)
	}
	card.OnClick(func() {
		c.WriteClipboard(text)
		a.copied = key
		c.Toast(label + "已复制")
	})
}

// chosenPeers are the peers ticked at step 1, in the order they were
// ticked, without the interface.
func (a *App) chosenPeers() []*backend.Peer {
	chosen := make([]*backend.Peer, 0, len(a.order))
	for _, key := range a.order {
		if key == a.iface {
			continue
		}
		if p := a.peerByKey(key); p != nil {
			chosen = append(chosen, p)
		}
	}
	return chosen
}

func (a *App) clearPeerSelection(key string) {
	delete(a.selected, key)
	a.order = slices.DeleteFunc(a.order, func(item string) bool { return item == key })
	delete(a.routeMap, key)
}

func (a *App) togglePeer(key string, on bool) {
	if !on {
		delete(a.selected, key)
		a.order = slices.DeleteFunc(a.order, func(item string) bool { return item == key })
		return
	}
	if !a.selected[key] {
		a.order = append(a.order, key)
	}
	a.selected[key] = true
}

func (a *App) toggleRoute(key, cidr string, on bool) {
	current := a.routeMap[key]
	if on {
		a.routeMap[key] = slices.DeleteFunc(current, func(item string) bool { return item == cidr })
		return
	}
	a.routeMap[key] = append(current, cidr)
}

func (a *App) resetConfig() {
	a.step = 0
	a.iface = ""
	a.selected = map[string]bool{}
	a.order = nil
	a.routeMap = map[string][]string{}
	a.config = ""
	a.blocks = nil
	a.copied = ""
	a.qr = nil
	a.qrFor = ""
}

// buildConfig builds the WireGuard configuration the web UI built, and
// returns why it cannot, or "".
func (a *App) buildConfig() string {
	if a.iface == "" {
		return "请选择 Interface"
	}
	chosen := a.chosenPeers()
	if len(chosen) == 0 {
		return "请选择 Peers"
	}
	local := a.peerByKey(a.iface)
	if local == nil {
		return "本地接口不存在"
	}

	lines := []string{
		"[Interface]",
		"# " + local.Username,
		"PrivateKey = " + local.Prikey,
		"Address = " + local.PrivateAddr + "/24",
	}
	if local.Port != 0 {
		lines = append(lines, "ListenPort = "+strconv.Itoa(local.Port))
	}
	if local.MTU != 0 {
		lines = append(lines, "MTU = "+strconv.Itoa(local.MTU))
	}
	if local.DNS != "" {
		lines = append(lines, "DNS = "+local.DNS)
	}

	blocks := []string{strings.Join(lines, "\n")}
	for _, peer := range chosen {
		ips := append([]string{peer.AllowedIPs}, a.routeMap[peerKey(peer.ID)]...)
		kept := make([]string, 0, len(ips))
		for _, ip := range ips {
			if ip != "" {
				kept = append(kept, ip)
			}
		}
		peerLines := []string{
			"[Peer]",
			"# " + peer.Username,
			"PublicKey = " + peer.Pubkey,
			"AllowedIPs = " + strings.Join(kept, ", "),
		}
		if peer.PublicAddr != "" {
			peerLines = append(peerLines, "Endpoint = "+peer.PublicAddr+":"+strconv.Itoa(peer.Port))
		}
		if peer.Keepalive != 0 {
			peerLines = append(peerLines, "PersistentKeepalive = "+strconv.Itoa(peer.Keepalive))
		}
		blocks = append(blocks, strings.Join(peerLines, "\n"))
	}

	a.blocks = blocks
	a.config = strings.Join(blocks, "\n\n")
	a.copied = ""
	a.step = 2

	// The QR code of the whole configuration, made once per configuration.
	if a.qrFor != a.config {
		if png, err := qrcode.Encode(a.config, qrcode.Low, configQRSize); err == nil {
			if bitmap, err := myui.DecodeBitmap(png); err == nil {
				a.qr = bitmap
				a.qrFor = a.config
			}
		}
	}
	return ""
}

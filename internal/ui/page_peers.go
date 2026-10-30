package ui

import (
	"fmt"
	"strconv"
	"strings"

	myui "github.com/egoist/mygo/ui"

	"wgm/backend"
)

// peerForm is a peer being added or changed.
type peerForm struct {
	ID          int // 0 表示新增
	UserName    string
	PrivateAddr string
	PublicAddr  string
	Port        string
	AllowedIPs  string
	MTU         string
	DNS         string
	Keepalive   string
}

func peerFormOf(p *backend.Peer) *peerForm {
	f := &peerForm{
		ID:          p.ID,
		UserName:    p.Username,
		PrivateAddr: p.PrivateAddr,
		PublicAddr:  p.PublicAddr,
		AllowedIPs:  p.AllowedIPs,
		MTU:         strconv.Itoa(p.MTU),
		DNS:         p.DNS,
		Keepalive:   strconv.Itoa(p.Keepalive),
	}
	if p.PublicAddr != "" {
		f.Port = strconv.Itoa(p.Port)
	}
	return f
}

func (a *App) peersPage(c *myui.Context) {
	t := c.Theme()
	pageBody(c, func() {
		pageHeader(c, "NETWORK INVENTORY", "节点", "管理本地接口与远程 WireGuard 节点", func() {
			myui.Text(c, fmt.Sprintf("%d 个节点", len(a.peers))).FontSize(13).TextColor(t.TextMuted)
			if myui.PrimaryButton(c, "增加节点").Height(cardButtonHeight).Padding(7, 16).Clicked() {
				a.editPeer = &peerForm{}
			}
		})

		myui.GridView(c, &a.peerGrid, len(a.peers), 340, 380, func(i int) {
			a.peerCard(c, &a.peers[i])
		}).FillWidth().Grow(1).MinWidth(0).MinHeight(0).MaxWidth(1120).AlignSelf(myui.Center).
			Gap(16).ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
			Padding(0, gridRightInset, gridBottomInset, 0).Margin(14, 0, 0, 0)
	})
}

func (a *App) peerCard(c *myui.Context, p *backend.Peer) {
	t := c.Theme()
	cardFrame(c, func() {
		myui.Row(c).AlignItems(myui.Center).MinWidth(0).Gap(8).Children(func() {
			myui.Column(c).Grow(1).MinWidth(0).Gap(4).Children(func() {
				myui.Text(c, "Peer "+strconv.Itoa(p.ID)).FontSize(12).TextColor(t.TextMuted)
				myui.Text(c, p.Username).FontSize(18).Bold().MinWidth(0)
			})
			status := "仅内网"
			if p.PublicAddr != "" {
				status = "已配置公网"
			}
			myui.Badge(c, status)
		})
	}, func() {
		displayEndpoint := "未设置公网"
		if p.PublicAddr != "" {
			displayEndpoint = endpoint(p.PublicAddr, p.Port)
		}
		myui.Grid(c).Columns(1).Gap(10).MinWidth(0).Children(func() {
			fact(c, "内网地址", p.PrivateAddr)
			fact(c, "公网端点", displayEndpoint)
		})
		myui.Column(c).Gap(6).Padding(12).Radius(8).MinWidth(0).Background(t.Background).Children(func() {
			myui.Text(c, "默认路由").FontSize(12).TextColor(t.TextMuted)
			myui.Text(c, orUnset(p.AllowedIPs)).Font("monospace").FontSize(13).Wrap().MinWidth(0)
		})
		myui.Grid(c).Columns(3).Gap(10).MinWidth(0).Children(func() {
			smallFact(c, "DNS", orUnset(p.DNS))
			smallFact(c, "MTU", orUnset(strconv.Itoa(p.MTU)))
			smallFact(c, "Keepalive", strconv.Itoa(p.Keepalive))
		})
	}, func() {
		if cardButton(c, "修改").Clicked() {
			a.editPeer = peerFormOf(p)
		}
		if dangerButton(c, "删除").Clicked() {
			deleted := *p
			a.delPeer = &deleted
		}
	})
}
func fact(c *myui.Context, label, value string) {
	myui.Column(c).Gap(5).MinWidth(0).Children(func() {
		myui.Text(c, label).FontSize(12).TextColor(c.Theme().TextMuted)
		myui.Text(c, value).Font("monospace").FontSize(13).Wrap().MinWidth(0)
	})
}

func smallFact(c *myui.Context, label, value string) {
	t := c.Theme()
	myui.Column(c).AlignItems(myui.Center).Gap(5).MinWidth(0).Padding(10).Radius(7).Background(t.Background).Children(func() {
		myui.Text(c, label).FontSize(11).TextColor(t.TextMuted)
		myui.Text(c, value).FontSize(12).Wrap().MinWidth(0)
	})
}

// orUnset shows 未设置 where the web UI showed an empty value.
func orUnset(value string) string {
	if value == "" || value == "0" {
		return "未设置"
	}
	return value
}

func (a *App) savePeer(c *myui.Context) {
	f := a.editPeer
	add := f.ID == 0

	privateAddr := strings.TrimSpace(f.PrivateAddr)
	publicAddr := strings.TrimSpace(f.PublicAddr)
	port := strings.TrimSpace(f.Port)
	allowedIPs := strings.TrimSpace(f.AllowedIPs)
	mtu := strings.TrimSpace(f.MTU)
	dns := strings.TrimSpace(f.DNS)
	keepalive := strings.TrimSpace(f.Keepalive)

	switch {
	case add && strings.TrimSpace(f.UserName) == "":
		c.Toast("请选择用户")
		return
	case privateAddr == "":
		c.Toast("请输入内网地址")
		return
	case !isIPv4(privateAddr):
		c.Toast("请输入正确的内网地址")
		return
	case publicAddr != "" && !isHost(publicAddr):
		c.Toast("请输入正确的公网地址或域名")
		return
	case publicAddr != "" && port == "":
		c.Toast("请输入公网地址和端口")
		return
	}

	portValue, err := parseOptionalInt(port)
	if err != nil || portValue < 0 || portValue > maxPort {
		c.Toast("请输入正确的端口")
		return
	}
	mtuValue, err := parseOptionalInt(mtu)
	if err != nil || (mtu != "" && (mtuValue < minMTU || mtuValue > maxMTU)) {
		c.Toast("请输入正确的MTU")
		return
	}
	keepaliveValue, err := parseOptionalInt(keepalive)
	if err != nil || keepaliveValue < 0 {
		c.Toast("请输入正确的Keepalive")
		return
	}

	in := backend.PeerInput{
		UserID:      a.userID(f.UserName),
		PrivateAddr: privateAddr,
		PublicAddr:  publicAddr,
		Port:        portValue,
		AllowedIPs:  allowedIPs,
		MTU:         mtuValue,
		DNS:         dns,
		Keepalive:   keepaliveValue,
	}

	if add {
		err = a.db.AddPeer(in)
	} else {
		err = a.db.UpdatePeer(f.ID, in)
	}
	if err != nil {
		if add {
			c.Toast("添加节点失败: " + err.Error())
		} else {
			c.Toast("更新节点失败: " + err.Error())
		}
		return
	}
	a.editPeer = nil
	a.loadPeers(c)
	if add {
		c.Toast("添加节点成功")
	} else {
		c.Toast("更新节点成功")
	}
}

func (a *App) deletePeer(c *myui.Context) {
	if err := a.db.DeletePeer(a.delPeer.ID); err != nil {
		c.Toast("删除节点失败: " + err.Error())
		return
	}
	a.delPeer = nil
	a.loadPeers(c)
	c.Toast("删除节点成功")
}

func (a *App) peerDialogs(c *myui.Context) {
	if f := a.editPeer; f != nil {
		open := true
		myui.Modal(c, &open, func() {
			title := "增加节点"
			if f.ID != 0 {
				title = "修改节点"
			}
			myui.Text(c, title).FontSize(16).Bold()
			submit := false
			myui.Grid(c).Columns(2).Gap(12).Width(600).Children(func() {
				if f.ID == 0 {
					myui.Field(c, "用户名", func() {
						myui.Select(c, &f.UserName, a.userNames())
					})
				} else {
					myui.Field(c, "用户名", func() {
						myui.TextInput(c, &f.UserName).ReadOnly(true)
					})
				}
				var privateAddr myui.Element
				myui.Field(c, "内网地址", func() {
					privateAddr = myui.TextInput(c, &f.PrivateAddr).AutoFocus()
				})
				if privateAddr.Submitted() {
					submit = true
				}
				myui.Field(c, "公网地址", func() { myui.TextInput(c, &f.PublicAddr) })
				myui.Field(c, "监听端口", func() { myui.TextInput(c, &f.Port) })
				myui.Field(c, "默认路由", func() { myui.TextInput(c, &f.AllowedIPs) })
				myui.Field(c, "MTU", func() { myui.TextInput(c, &f.MTU) })
				myui.Field(c, "DNS", func() { myui.TextInput(c, &f.DNS) })
				myui.Field(c, "Keepalive", func() { myui.TextInput(c, &f.Keepalive) })
			})
			myui.Row(c).Justify(myui.End).Gap(8).Margin(10, 0, 0, 0).Children(func() {
				if myui.Button(c, "取消").Clicked() {
					a.editPeer = nil
				}
				if myui.PrimaryButton(c, "提交").Clicked() {
					a.savePeer(c)
				}
			})
			if submit {
				a.savePeer(c)
			}
		})
		if !open {
			a.editPeer = nil
		}
	}

	if a.delPeer != nil {
		open := true
		deleteConfirmation(c, &open, "确认删除节点", func() { a.deletePeer(c) })
		if !open {
			a.delPeer = nil
		}
	}
}

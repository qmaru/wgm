package ui

import (
	"strconv"
	"strings"
	"time"

	"github.com/egoist/mygo"
	myui "github.com/egoist/mygo/ui"

	"github.com/qmaru/wgm/backend"
	"github.com/qmaru/wgm/common/utils"
)

const (
	themeLight = iota
	themeDark
	themeAuto
)

type loadScope uint8

const (
	loadUsersScope loadScope = 1 << iota
	loadPeersScope
	loadRoutesScope
	loadAllScope = loadUsersScope | loadPeersScope | loadRoutesScope
)

type appLoadResult struct {
	scope    loadScope
	users    []backend.User
	peers    []backend.Peer
	routes   []backend.Route
	userErr  error
	peerErr  error
	routeErr error
}

// App is everything the window draws.
type App struct {
	db      backend.Store
	version string
	tab     int
	theme   int
	loaded  bool
	loading bool

	loadDone        chan appLoadResult
	reloadRequested loadScope

	users    []backend.User
	userGrid myui.GridState
	userName string
	editUser *backend.User
	delUser  *backend.User

	peers    []backend.Peer
	peerGrid myui.GridState
	editPeer *peerForm
	delPeer  *backend.Peer

	routes    []backend.Route
	routeGrid myui.GridState
	cidr      string
	editRoute *backend.Route
	delRoute  *backend.Route

	// 配置向导
	step          int
	interfaceGrid myui.GridState
	wizardGrid    myui.GridState
	iface         string
	selected      map[string]bool
	order         []string
	routeMap      map[string][]string
	config        string
	blocks        []string
	copied        string
	qr            *myui.Bitmap
	qrFor         string
}

func NewApp(db backend.Store) *App {
	a := &App{
		db:       db,
		version:  utils.AppVersion,
		theme:    themeAuto,
		selected: map[string]bool{},
		routeMap: map[string][]string{},
		loadDone: make(chan appLoadResult, 1),
	}
	return a
}

const (
	pageMaxWidth        = 1180
	pagePaddingX        = 20
	pagePaddingY        = 16
	cardHeadHeight      = 64
	cardActionHeight    = 58
	cardButtonHeight    = 36
	cardContentGap      = 8
	gridGap             = 8
	gridRightInset      = 12
	gridBottomInset     = 16
	configQRSize        = 512
	configQRDisplaySize = 240
)

var (
	cardStateTransition = myui.ElementTransition{
		Colors:   true,
		Duration: 120 * time.Millisecond,
		Ease:     myui.EaseOut,
	}
	pageTransition = myui.ElementTransition{
		Duration: 160 * time.Millisecond,
		Ease:     myui.EaseOut,
		Enter:    &myui.Motion{X: 12, Opacity: 0, Scale: 0.995},
	}
)

func applyWireGuardTheme(c *myui.Context) {
	applyWireGuardThemeMode(c, themeAuto)
}

func setApplicationTheme(mode int) {
	switch mode {
	case themeLight:
		mygo.Theme.SetSource(mygo.ThemeLight)
	case themeDark:
		mygo.Theme.SetSource(mygo.ThemeDark)
	default:
		mygo.Theme.SetSource(mygo.ThemeSystem)
	}
}

func applyWireGuardThemeMode(c *myui.Context, mode int) {
	t := *c.Theme()
	switch mode {
	case themeLight:
		t.Dark = false
	case themeDark:
		t.Dark = true
	}
	t.Radius = 9
	t.Spacing = 4
	t.FontSize = 14
	if t.Dark {
		t.Background = myui.Hex("#191617")
		t.Surface = myui.Hex("#292324")
		t.SurfaceHover = myui.Hex("#382C2E")
		t.SurfacePressed = myui.Hex("#473639")
		t.Border = myui.Hex("#4B3B3D")
		t.Text = myui.Hex("#F4EEEE")
		t.TextMuted = myui.Hex("#B7A8AA")
		t.Accent = myui.Hex("#EA6269")
		t.AccentHover = myui.Hex("#F47A80")
		t.AccentPressed = myui.Hex("#CF4C53")
		t.AccentText = myui.Hex("#241012")
		t.Danger = myui.Hex("#FFC16A")
		t.Warning = myui.Hex("#F2C56D")
		t.Inverse = myui.Hex("#F4EEEE")
		t.InverseText = myui.Hex("#271D1F")
	} else {
		t.Background = myui.Hex("#FBF8F8")
		t.Surface = myui.Hex("#F1ECEC")
		t.SurfaceHover = myui.Hex("#E8DEDF")
		t.SurfacePressed = myui.Hex("#DCCFD0")
		t.Border = myui.Hex("#D9CDCE")
		t.Text = myui.Hex("#271D1F")
		t.TextMuted = myui.Hex("#746568")
		t.Accent = myui.Hex("#981C23")
		t.AccentHover = myui.Hex("#81171D")
		t.AccentPressed = myui.Hex("#6C1218")
		t.AccentText = myui.Hex("#FFFFFF")
		t.Danger = myui.Hex("#A85B00")
		t.Warning = myui.Hex("#946200")
		t.Inverse = myui.Hex("#271D1F")
		t.InverseText = myui.Hex("#F4EEEE")
	}
	// Keep wheel and keyboard scrolling available without overlaying cards.
	t.Scrollbar = myui.Transparent
	t.Focus = t.Accent
	t.Selection = t.Accent.Alpha(0.28)
	c.SetTheme(&t)
}

// deleteConfirmation keeps destructive actions amber so they do not compete
// with the WireGuard red accent.
func deleteConfirmation(c *myui.Context, open *bool, title string, confirm func()) {
	myui.Modal(c, open, func() {
		myui.Text(c, title).FontSize(15).Bold()
		myui.Text(c, "删除后不可恢复。").FontSize(13).TextColor(c.Theme().TextMuted)
		myui.Row(c).Justify(myui.End).Gap(8).Margin(8, 0, 0, 0).Children(func() {
			if myui.Button(c, "取消").Clicked() {
				*open = false
			}
			if dangerButton(c, "确认删除").Clicked() {
				confirm()
			}
		})
	})
}

func cardButton(c *myui.Context, label string) myui.Element {
	return myui.Button(c, label).Height(cardButtonHeight).Padding(5, 12)
}

func dangerButton(c *myui.Context, label string) myui.Element {
	danger := c.Theme().Danger
	button := myui.Button(c, label).Height(cardButtonHeight).Padding(5, 12).Border(1, danger).TextColor(danger)
	if button.Hovered() {
		button.Background(danger.Alpha(0.12))
	}
	if button.Pressed() {
		button.Background(danger.Alpha(0.22))
	}
	return button
}

// view builds the window: the version, the tabs, the page chosen and the
// dialogs over it.
func (a *App) View(c *myui.Context) {
	applyWireGuardThemeMode(c, a.theme)
	a.applyPendingLoad(c)
	if !a.loaded {
		a.requestLoad(c, loadAllScope)
	}

	myui.Column(c).Fill().MinWidth(0).Children(func() {
		myui.Row(c).Justify(myui.SpaceBetween).AlignItems(myui.Center).Padding(10, pagePaddingX, 0, pagePaddingX).Shrink(0).Children(func() {
			myui.Text(c, "v"+a.version).FontSize(12).TextColor(c.Theme().TextMuted)
			if myui.Segmented(c, &a.theme, "Light", "Dark", "Auto").Label("Theme").Changed() {
				setApplicationTheme(a.theme)
			}
		})
		myui.Row(c).Justify(myui.Center).Padding(0, pagePaddingX).Shrink(0).Children(func() {
			myui.Tabs(c, &a.tab, "配置", "节点", "用户", "路由")
		})
		page := myui.Column(c).Key(a.tab).Grow(1).MinWidth(0).MinHeight(0)
		page.Transition(pageTransition).Children(func() {
			switch a.tab {
			case 0:
				a.configsPage(c)
			case 1:
				a.peersPage(c)
			case 2:
				a.usersPage(c)
			case 3:
				a.routesPage(c)
			}
		})
		a.dialogs(c)
	})
}

// dialogs builds the dialogs of every page; only the open ones draw.
func (a *App) dialogs(c *myui.Context) {
	a.userDialogs(c)
	a.peerDialogs(c)
	a.routeDialogs(c)
}

func pageScroll(c *myui.Context, fn func()) {
	myui.Scroll(c).Grow(1).MinWidth(0).MinHeight(0).
		ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
		Padding(0, gridRightInset, gridBottomInset, 0).Children(fn)
}

// pageBody is the centered column every page's content sits in.
func pageBody(c *myui.Context, fn func()) {
	myui.Row(c).Grow(1).MinWidth(0).MinHeight(0).Justify(myui.Center).AlignItems(myui.Stretch).Children(func() {
		myui.Column(c).MaxWidth(pageMaxWidth).Grow(1).MinWidth(0).MinHeight(0).
			Padding(pagePaddingY, pagePaddingX).Children(fn)
	})
}

// pageHeader is a page's title block, with what the page shows beside it.
func pageHeader(c *myui.Context, eyebrow, title, subtitle string, beside func()) {
	t := c.Theme()
	myui.Row(c).AlignItems(myui.End).Gap(12).Wrap().MinWidth(0).Shrink(0).Children(func() {
		myui.Column(c).Grow(1).MinWidth(0).Gap(3).Children(func() {
			myui.Text(c, eyebrow).FontSize(11).Bold().TextColor(t.Accent)
			myui.Text(c, title).FontSize(22).Bold()
			myui.Text(c, subtitle).FontSize(13).TextColor(t.TextMuted).Wrap()
		})
		if beside != nil {
			myui.Row(c).Gap(8).Shrink(0).AlignItems(myui.Center).Children(beside)
		}
	})
	myui.Divider(c).Shrink(0).Margin(12, 0, 16, 0)
}

func (a *App) requestLoad(c *myui.Context, scope loadScope) {
	if a.loading {
		a.reloadRequested |= scope
		return
	}
	if a.loadDone == nil {
		a.loadDone = make(chan appLoadResult, 1)
	}
	a.loading = true
	services := c.Services()
	go func() {
		result := appLoadResult{scope: scope}
		if scope&loadUsersScope != 0 {
			result.users, result.userErr = a.db.Users()
		}
		if scope&loadPeersScope != 0 {
			result.peers, result.peerErr = a.db.Peers()
		}
		if scope&loadRoutesScope != 0 {
			result.routes, result.routeErr = a.db.Routes()
		}
		a.loadDone <- result
		services.Invalidate()
	}()
}

func (a *App) applyPendingLoad(c *myui.Context) {
	if a.loadDone == nil {
		return
	}
	select {
	case result := <-a.loadDone:
		a.loading = false
		a.loaded = true
		if result.scope&loadUsersScope != 0 {
			if result.userErr == nil {
				a.users = result.users
			} else {
				c.Toast("用户列表获取失败: " + result.userErr.Error())
			}
		}
		if result.scope&loadPeersScope != 0 {
			if result.peerErr == nil {
				a.peers = result.peers
			} else {
				c.Toast("节点列表获取失败: " + result.peerErr.Error())
			}
		}
		if result.scope&loadRoutesScope != 0 {
			if result.routeErr == nil {
				a.routes = result.routes
			} else {
				c.Toast("路由列表获取失败: " + result.routeErr.Error())
			}
		}
		if a.reloadRequested != 0 {
			scope := a.reloadRequested
			a.reloadRequested = 0
			a.requestLoad(c, scope)
		}
	default:
	}
}

func (a *App) refreshData(c *myui.Context, scope loadScope) {
	a.requestLoad(c, scope)
}

// userNames are the users a new peer may belong to.
func (a *App) userNames() []string {
	names := make([]string, 0, len(a.users))
	for _, u := range a.users {
		names = append(names, u.Username)
	}
	return names
}

func (a *App) userID(name string) int {
	for _, u := range a.users {
		if u.Username == name {
			return u.ID
		}
	}
	return 0
}

func peerKey(id int) string { return strconv.Itoa(id) }

func endpoint(addr string, port int) string {
	if addr == "" || port == 0 || strings.Contains(addr, ":") {
		return addr
	}
	return addr + ":" + strconv.Itoa(port)
}

func (a *App) peerByKey(key string) *backend.Peer {
	for i := range a.peers {
		if peerKey(a.peers[i].ID) == key {
			return &a.peers[i]
		}
	}
	return nil
}

func (a *App) peerWithoutInterface(index int) *backend.Peer {
	for i := range a.peers {
		if peerKey(a.peers[i].ID) == a.iface {
			continue
		}
		if index == 0 {
			return &a.peers[i]
		}
		index--
	}
	return nil
}

func (a *App) selectedPeersSummary() string {
	var summary strings.Builder
	for _, key := range a.order {
		if key == a.iface {
			continue
		}
		p := a.peerByKey(key)
		if p == nil {
			continue
		}
		if summary.Len() > 0 {
			summary.WriteString("、")
		}
		summary.WriteString(p.Username)
		summary.WriteString(" · ")
		summary.WriteString(p.PrivateAddr)
	}
	if summary.Len() == 0 {
		return "未选择"
	}
	return summary.String()
}

func cardFrame(c *myui.Context, head, content, actions func()) {
	cardShell(c, true, head, content, actions)
}

func cardShell(c *myui.Context, fillHeight bool, head, content, actions func()) myui.Element {
	t := c.Theme()
	card := myui.Column(c).MinWidth(0)
	if fillHeight {
		card.Fill()
	} else {
		card.FillWidth()
	}
	card.Radius(10).Border(1, t.Border).Background(t.Surface).Children(func() {
		myui.Column(c).Height(cardHeadHeight).MinWidth(0).Padding(14, 18).Children(head)
		myui.Divider(c)
		myui.Column(c).Grow(1).MinWidth(0).MinHeight(0).Gap(cardContentGap).Padding(14, 18).Children(content)
		myui.Divider(c)
		myui.Row(c).Height(cardActionHeight).Gap(8).Padding(10, 18).
			AlignItems(myui.Center).Wrap().Children(actions)
	})
	background, border := t.Surface, t.Border
	if card.Hovered() {
		background = t.SurfaceHover
		border = t.Accent.Alpha(0.45)
	}
	card.Background(background).BorderColor(border).Transition(cardStateTransition)
	return card
}

// keyBlock shows a key in a code block that copies on click, as the web
// UI's copyable previews do.
func keyBlock(c *myui.Context, text, tip string) bool {
	t := c.Theme()
	block := myui.Column(c).FillWidth().MinWidth(0).Padding(10).Radius(7).Border(1, t.Border).
		Background(t.Background).Tooltip(tip)
	block.Children(func() {
		myui.Text(c, text).Font("monospace").FontSize(13).SingleLine().MinWidth(0)
	})
	return block.Clicked()
}

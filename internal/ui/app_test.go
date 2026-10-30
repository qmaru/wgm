package ui

import (
	"strconv"
	"strings"
	"testing"

	myui "github.com/egoist/mygo/ui"

	"github.com/qmaru/wgm/backend"
)

func TestBuildConfig(t *testing.T) {
	a := &App{
		peers: []backend.Peer{
			{ID: 1, Username: "local", Prikey: "PRIVATE", PrivateAddr: "10.0.0.1", Port: 51820, MTU: 1420, DNS: "1.1.1.1"},
			{ID: 2, Username: "remote", Pubkey: "PUBLIC", PrivateAddr: "10.0.0.2", PublicAddr: "vpn.example.com", Port: 51821, AllowedIPs: "10.0.0.2/32", Keepalive: 25},
		},
		iface:    "1",
		order:    []string{"2"},
		routeMap: map[string][]string{"2": {"192.168.1.0/24"}},
	}

	if msg := a.buildConfig(); msg != "" {
		t.Fatal(msg)
	}
	want := strings.Join([]string{
		"[Interface]\n# local\nPrivateKey = PRIVATE\nAddress = 10.0.0.1/24\nListenPort = 51820\nMTU = 1420\nDNS = 1.1.1.1",
		"[Peer]\n# remote\nPublicKey = PUBLIC\nAllowedIPs = 10.0.0.2/32, 192.168.1.0/24\nEndpoint = vpn.example.com:51821\nPersistentKeepalive = 25",
	}, "\n\n")
	if a.config != want {
		t.Fatalf("config = %q, want %q", a.config, want)
	}
	if a.step != 2 || len(a.blocks) != 2 {
		t.Fatalf("step=%d blocks=%d, want step 2 and two blocks", a.step, len(a.blocks))
	}
	if a.qr == nil {
		t.Fatal("expected a QR code for the generated configuration")
	}
	if width, height := myui.ImageSize(a.qr); width != configQRSize || height != configQRSize {
		t.Fatalf("QR size = %gx%g, want %dx%d", width, height, configQRSize, configQRSize)
	}
}

func TestBuildConfigRequiresInterfaceAndPeer(t *testing.T) {
	a := &App{}
	if got := a.buildConfig(); got != "请选择 Interface" {
		t.Fatalf("message = %q", got)
	}
	a.iface = "1"
	if got := a.buildConfig(); got != "请选择 Peers" {
		t.Fatalf("message = %q", got)
	}
}

func TestListsReachLastCardAfterScrolling(t *testing.T) {
	users := make([]backend.User, 6)
	peers := make([]backend.Peer, 6)
	routes := make([]backend.Route, 6)
	for i := range users {
		id := i + 1
		users[i] = backend.User{ID: id, Username: "user-" + string(rune('a'+i))}
		peers[i] = backend.Peer{ID: id, Username: users[i].Username, PrivateAddr: "10.0.0." + strconv.Itoa(id)}
		routes[i] = backend.Route{ID: id, CIDR: "10.0." + strconv.Itoa(i) + ".0/24"}
	}

	cases := []struct {
		name string
		tab  int
		last string
	}{
		{"users", 2, "user-f"},
		{"peers", 1, "user-f"},
		{"routes", 3, "10.0.5.0/24"},
	}
	for _, tc := range cases {
		for _, size := range []struct {
			name          string
			width, height int
		}{
			{name: "desktop", width: 1280, height: 720},
			{name: "compact", width: 860, height: 520},
		} {
			t.Run(tc.name+"/"+size.name, func(t *testing.T) {
				a := &App{loaded: true, version: "test", users: users, peers: peers, routes: routes}
				a.tab = tc.tab
				tester := myui.NewTester(a.View, size.width, size.height)
				tester.Scroll(float32(size.width/2), float32(size.height-40), 0, 1200)
				if !tester.HasText(tc.last) {
					t.Fatalf("last %s item is not visible after scrolling: %q", tc.name, tester.Texts())
				}
			})
		}
	}
}

func TestPageHeadersStayFixedWhileContentScrolls(t *testing.T) {
	users := make([]backend.User, 6)
	peers := make([]backend.Peer, 6)
	routes := make([]backend.Route, 6)
	for i := range users {
		id := i + 1
		users[i] = backend.User{ID: id, Username: "user-" + strconv.Itoa(id)}
		peers[i] = backend.Peer{ID: id, Username: users[i].Username, PrivateAddr: "10.0.0." + strconv.Itoa(id)}
		routes[i] = backend.Route{ID: id, CIDR: "10.0." + strconv.Itoa(i) + ".0/24"}
	}

	cases := []struct {
		name, eyebrow, last string
		tab                 int
	}{
		{"configs", "CONFIGURATION BUILDER", "user-6", 0},
		{"peers", "NETWORK INVENTORY", "user-6", 1},
		{"users", "IDENTITY VAULT", "user-6", 2},
		{"routes", "ROUTING TABLE", "10.0.5.0/24", 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := &App{loaded: true, version: "test", users: users, peers: peers, routes: routes}
			a.tab = tc.tab
			tester := myui.NewTester(a.View, 860, 520)
			before, ok := tester.Find(tc.eyebrow)
			if !ok {
				t.Fatalf("page header %q is not visible", tc.eyebrow)
			}
			tester.Scroll(430, 460, 0, 1200)
			after, ok := tester.Find(tc.eyebrow)
			if !ok {
				t.Fatalf("page header %q disappeared after scrolling", tc.eyebrow)
			}
			if before.Y != after.Y {
				t.Fatalf("page header moved from y=%v to y=%v", before.Y, after.Y)
			}
			if !tester.HasText(tc.last) {
				t.Fatalf("last %s item is not visible after inner scrolling: %q", tc.name, tester.Texts())
			}
		})
	}
}

func TestConfigurationCardsHandleClicksByRegion(t *testing.T) {
	a := &App{
		loaded:   true,
		version:  "test",
		peers:    []backend.Peer{{ID: 1, Username: "local", PrivateAddr: "10.0.0.1"}, {ID: 2, Username: "remote", PrivateAddr: "10.0.0.2"}},
		routes:   []backend.Route{{ID: 1, CIDR: "192.168.1.0/24"}},
		selected: map[string]bool{},
		routeMap: map[string][]string{},
	}
	tester := myui.NewTester(a.View, 860, 520)
	if err := tester.Click("10.0.0.2"); err != nil {
		t.Fatal(err)
	}
	if a.iface != "2" {
		t.Fatalf("interface = %q, want 2", a.iface)
	}

	a.iface = "1"
	a.step = 1
	tester.Frame()
	if err := tester.Click("10.0.0.2"); err != nil {
		t.Fatal(err)
	}
	if !a.selected["2"] {
		t.Fatal("clicking the peer content should select the peer")
	}
	if err := tester.Click("192.168.1.0/24"); err != nil {
		t.Fatal(err)
	}
	if got := a.routeMap["2"]; len(got) != 1 || got[0] != "192.168.1.0/24" {
		t.Fatalf("routes = %v, want the route button to toggle independently", got)
	}
}

func TestParseOptionalInt(t *testing.T) {
	if got, err := parseOptionalInt(""); err != nil || got != 0 {
		t.Fatalf("empty value = %d, %v", got, err)
	}
	if got, err := parseOptionalInt("51820"); err != nil || got != 51820 {
		t.Fatalf("numeric value = %d, %v", got, err)
	}
	if _, err := parseOptionalInt("not a number"); err == nil {
		t.Fatal("expected invalid number to fail")
	}
}

func TestPeerAddressAndKeyFormatting(t *testing.T) {
	for _, ip := range []string{"10.0.0.1", "192.168.1.255"} {
		if !isIPv4(ip) {
			t.Errorf("isIPv4(%q) = false", ip)
		}
	}
	for _, ip := range []string{"10.0.0", "10.0.0.256", "host.example"} {
		if isIPv4(ip) {
			t.Errorf("isIPv4(%q) = true", ip)
		}
	}
	if !isHost("vpn.example.com") || isHost("-invalid.example") {
		t.Fatal("host validation did not match expected values")
	}
	if got := displayKey("abcdefghijklmnopqrstuvwxyz"); got != "abcd...uvwxyz" {
		t.Fatalf("displayKey = %q", got)
	}
}

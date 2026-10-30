package ui

import (
	"testing"

	myui "github.com/egoist/mygo/ui"
)

func TestWireGuardThemeUsesDifferentBrandAndDangerColors(t *testing.T) {
	for _, dark := range []bool{false, true} {
		var got myui.Theme
		tester := myui.NewTester(func(c *myui.Context) {
			applyWireGuardTheme(c)
			got = *c.Theme()
		}, 320, 240)
		tester.SetDark(dark)
		tester.Frame()

		if got.Dark != dark {
			t.Errorf("theme dark = %v, want %v", got.Dark, dark)
		}
		if got.Accent == got.Danger {
			t.Errorf("brand accent and destructive danger colors should differ in dark=%v", dark)
		}
		if got.Scrollbar != myui.Transparent {
			t.Errorf("scrollbars should stay hidden in dark=%v", dark)
		}
	}
}

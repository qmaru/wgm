package ui

import (
	"fmt"
	"strings"

	myui "github.com/egoist/mygo/ui"

	"github.com/qmaru/wgm/backend"
)

func (a *App) routesPage(c *myui.Context) {
	t := c.Theme()
	pageBody(c, func() {
		pageHeader(c, "ROUTING TABLE", "路由", "维护可分发到节点的 CIDR 路由规则", func() {
			myui.Text(c, fmt.Sprintf("%d 条规则", len(a.routes))).FontSize(13).TextColor(t.TextMuted)
		})

		myui.Row(c).Gap(0).AlignItems(myui.End).Wrap().MaxWidth(420).Shrink(0).Children(func() {
			myui.Field(c, "路由", func() {
				myui.TextInput(c, &a.cidr).Placeholder("路由").Height(cardButtonHeight).
					Radius(8, 0, 0, 8).Grow(1).MinWidth(0)
			}).Grow(1).MinWidth(160)
			if myui.PrimaryButton(c, "提交").Height(cardButtonHeight).Padding(5, 16).Radius(0, 8, 8, 0).Clicked() {
				a.addRoute(c)
			}
		})

		myui.GridView(c, &a.routeGrid, len(a.routes), 240, 180, func(i int) {
			a.routeCard(c, &a.routes[i])
		}).FillWidth().Grow(1).MinWidth(0).MinHeight(0).Gap(gridGap).
			ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
			Padding(0, gridRightInset, gridBottomInset, 0).Margin(14, 0, 0, 0)
	})
}

func (a *App) routeCard(c *myui.Context, r *backend.Route) {
	t := c.Theme()
	cardFrame(c, func() {
		myui.Text(c, "CIDR rule").FontSize(11).Bold().TextColor(t.Accent)
	}, func() {
		myui.Text(c, r.CIDR).Font("monospace").FontSize(17).Bold().Wrap().MinWidth(0)
	}, func() {
		if cardButton(c, "修改").Clicked() {
			edited := *r
			a.editRoute = &edited
		}
		if dangerButton(c, "删除").Clicked() {
			deleted := *r
			a.delRoute = &deleted
		}
	})
}

func (a *App) addRoute(c *myui.Context) {
	cidr := strings.TrimSpace(a.cidr)
	if cidr == "" {
		c.Toast("请输入路由")
		return
	}
	if err := a.db.AddRoute(cidr); err != nil {
		c.Toast("添加路由失败: " + err.Error())
		return
	}
	a.cidr = ""
	a.refreshData(c, loadRoutesScope)
	c.Toast("添加路由成功")
}

func (a *App) updateRoute(c *myui.Context) {
	cidr := strings.TrimSpace(a.editRoute.CIDR)
	if cidr == "" {
		c.Toast("请输入路由")
		return
	}
	if err := a.db.UpdateRoute(a.editRoute.ID, cidr); err != nil {
		c.Toast("更新路由失败: " + err.Error())
		return
	}
	a.editRoute = nil
	a.refreshData(c, loadRoutesScope)
	c.Toast("更新路由成功")
}

func (a *App) deleteRoute(c *myui.Context) {
	if err := a.db.DeleteRoute(a.delRoute.ID); err != nil {
		c.Toast("删除路由失败: " + err.Error())
		return
	}
	a.delRoute = nil
	a.refreshData(c, loadRoutesScope)
	c.Toast("删除路由成功")
}

func (a *App) routeDialogs(c *myui.Context) {
	if r := a.editRoute; r != nil {
		open := true
		myui.Modal(c, &open, func() {
			myui.Text(c, "修改路由").FontSize(16).Bold()
			var input myui.Element
			myui.Field(c, "路由", func() {
				input = myui.TextInput(c, &r.CIDR).Width(240).AutoFocus()
			})
			if input.Submitted() {
				a.updateRoute(c)
			}
			myui.Row(c).Justify(myui.End).Gap(8).Margin(8, 0, 0, 0).Children(func() {
				if myui.Button(c, "取消").Clicked() {
					a.editRoute = nil
				}
				if myui.PrimaryButton(c, "提交").Clicked() {
					a.updateRoute(c)
				}
			})
		})
		if !open {
			a.editRoute = nil
		}
	}

	if a.delRoute != nil {
		open := true
		deleteConfirmation(c, &open, "确认删除路由", func() { a.deleteRoute(c) })
		if !open {
			a.delRoute = nil
		}
	}
}

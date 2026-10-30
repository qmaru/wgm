package ui

import (
	"fmt"
	"strings"

	myui "github.com/egoist/mygo/ui"

	"wgm/backend"
)

func (a *App) usersPage(c *myui.Context) {
	t := c.Theme()
	pageBody(c, func() {
		pageHeader(c, "IDENTITY VAULT", "用户", "管理用户身份与 WireGuard 密钥", func() {
			myui.Text(c, fmt.Sprintf("%d 个用户", len(a.users))).FontSize(13).TextColor(t.TextMuted)
		})

		myui.Row(c).Gap(0).AlignItems(myui.End).Wrap().MaxWidth(420).Shrink(0).Children(func() {
			myui.Field(c, "用户名", func() {
				myui.TextInput(c, &a.userName).Placeholder("用户名").Height(cardButtonHeight).
					Radius(8, 0, 0, 8).Grow(1).MinWidth(0)
			}).Grow(1).MinWidth(160)
			if myui.PrimaryButton(c, "提交").Height(cardButtonHeight).Padding(5, 16).Radius(0, 8, 8, 0).Clicked() {
				a.addUser(c)
			}
		})

		myui.GridView(c, &a.userGrid, len(a.users), 260, 276, func(i int) {
			a.userCard(c, &a.users[i])
		}).FillWidth().Grow(1).MinWidth(0).MinHeight(0).Gap(gridGap).
			ScrollbarInsets(0, gridRightInset, gridBottomInset, 0).
			Padding(0, gridRightInset, gridBottomInset, 0).Margin(14, 0, 0, 0)
	})
}

func (a *App) userCard(c *myui.Context, u *backend.User) {
	t := c.Theme()
	cardFrame(c, func() {
		myui.Column(c).Gap(4).Children(func() {
			myui.Text(c, fmt.Sprintf("USER %d", u.ID)).FontSize(12).TextColor(t.TextMuted)
			myui.Text(c, u.Username).FontSize(18).Bold().MinWidth(0)
		})
	}, func() {
		myui.Text(c, "Private key").FontSize(12).Bold().TextColor(t.Accent)
		if keyBlock(c, displayKey(u.PrivateKey), "点击复制完整 Private key") {
			a.copyKey(c, u.PrivateKey)
		}
		myui.Text(c, "Public key").FontSize(12).Bold().TextColor(t.Accent)
		if keyBlock(c, displayKey(u.PublicKey), "点击复制完整 Public key") {
			a.copyKey(c, u.PublicKey)
		}
	}, func() {
		if cardButton(c, "修改").Clicked() {
			edited := *u
			a.editUser = &edited
		}
		if dangerButton(c, "删除").Clicked() {
			deleted := *u
			a.delUser = &deleted
		}
	})
}

func (a *App) copyKey(c *myui.Context, key string) {
	c.WriteClipboard(key)
	c.Toast("密钥已复制")
}

func (a *App) addUser(c *myui.Context) {
	name := strings.TrimSpace(a.userName)
	if name == "" {
		c.Toast("请输入用户名")
		return
	}
	if err := a.db.AddUser(name); err != nil {
		c.Toast("添加用户失败: " + err.Error())
		return
	}
	a.userName = ""
	a.refreshData(c, loadUsersScope)
	c.Toast("添加用户成功")
}

func (a *App) updateUser(c *myui.Context) {
	name := strings.TrimSpace(a.editUser.Username)
	if name == "" {
		c.Toast("请输入用户名")
		return
	}
	if err := a.db.UpdateUser(a.editUser.ID, name); err != nil {
		c.Toast("更新用户失败: " + err.Error())
		return
	}
	a.editUser = nil
	a.refreshData(c, loadUsersScope)
	c.Toast("更新用户成功")
}

func (a *App) deleteUser(c *myui.Context) {
	if err := a.db.DeleteUser(a.delUser.ID); err != nil {
		c.Toast("删除用户失败: " + err.Error())
		return
	}
	a.delUser = nil
	a.refreshData(c, loadUsersScope)
	c.Toast("删除用户成功")
}

func (a *App) userDialogs(c *myui.Context) {
	if u := a.editUser; u != nil {
		open := true
		myui.Modal(c, &open, func() {
			myui.Text(c, "修改用户").FontSize(16).Bold()
			var input myui.Element
			myui.Field(c, "用户名", func() {
				input = myui.TextInput(c, &u.Username).Width(260).AutoFocus()
			})
			if input.Submitted() {
				a.updateUser(c)
			}
			myui.Row(c).Justify(myui.End).Gap(8).Margin(8, 0, 0, 0).Children(func() {
				if myui.Button(c, "取消").Clicked() {
					a.editUser = nil
				}
				if myui.PrimaryButton(c, "提交").Clicked() {
					a.updateUser(c)
				}
			})
		})
		if !open {
			a.editUser = nil
		}
	}

	if a.delUser != nil {
		open := true
		deleteConfirmation(c, &open, "确认删除用户", func() { a.deleteUser(c) })
		if !open {
			a.delUser = nil
		}
	}
}

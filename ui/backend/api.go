package backend

import (
	"context"

	"wgm/dbs/models"
	"wgm/services/data"
	"wgm/services/peers"
	"wgm/services/routes"
	"wgm/services/users"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// User API
func (a *App) UserListAPI() map[string]any {
	data, err := users.UserList()
	if err != nil {
		return JsonData(0, "用户列表获取失败: "+err.Error(), []any{})
	}
	return JsonData(1, "用户列表", data)
}

func (a *App) UserAddAPI(userData models.Users) map[string]any {
	err := users.UserAdd(&userData)
	if err != nil {
		return JsonData(0, "添加用户失败: "+err.Error(), []any{})
	}
	return JsonData(1, "添加用户成功", []any{})
}

func (a *App) UserUpdateAPI(userID string, userData models.Users) map[string]any {
	userIDN, err := IDtoInt(userID)
	if err != nil {
		return JsonData(0, "用户ID错误", []any{})
	}
	userData.CommonModel.ID = userIDN
	err = users.UserUpdate(&userData)
	if err != nil {
		return JsonData(0, "更新用户失败: "+err.Error(), []any{})
	}
	return JsonData(1, "更新用户成功", []any{})
}

func (a *App) UserDeleteAPI(userID string) map[string]any {
	userIDN, err := IDtoInt(userID)
	if err != nil {
		return JsonData(0, "用户ID错误", []any{})
	}

	err = users.UserDelete(userIDN)
	if err != nil {
		return JsonData(0, "删除用户失败: "+err.Error(), []any{})
	}
	return JsonData(1, "删除用户成功", []any{})
}

// Route API
func (a *App) RouteListAPI() map[string]any {
	data, err := routes.RouteList()
	if err != nil {
		return JsonData(0, "路由列表获取失败: "+err.Error(), []any{})
	}
	return JsonData(1, "路由列表", data)
}

func (a *App) RouteAddAPI(routeData models.Routes) map[string]any {
	err := routes.RouteAdd(&routeData)
	if err != nil {
		return JsonData(0, "添加路由失败: "+err.Error(), []any{})
	}
	return JsonData(1, "添加路由成功", []any{})
}

func (a *App) RouteUpdateAPI(routeID string, routeData models.Routes) map[string]any {
	routeIDN, err := IDtoInt(routeID)
	if err != nil {
		return JsonData(0, "路由ID错误", []any{})
	}

	routeData.CommonModel.ID = routeIDN
	err = routes.RouteUpdate(&routeData)
	if err != nil {
		return JsonData(0, "更新路由失败: "+err.Error(), []any{})
	}
	return JsonData(1, "更新路由成功", []any{})
}

func (a *App) RouteDeleteAPI(routeID string) map[string]any {
	routeIDN, err := IDtoInt(routeID)
	if err != nil {
		return JsonData(0, "路由ID错误", []any{})
	}

	err = routes.RouteDelete(routeIDN)
	if err != nil {
		return JsonData(0, "删除路由失败: "+err.Error(), []any{})
	}
	return JsonData(1, "删除路由成功", []any{})
}

// Peer API
func (a *App) PeerListAPI() map[string]any {
	data, err := peers.PeerList()
	if err != nil {
		return JsonData(0, "节点列表获取失败: "+err.Error(), []any{})
	}
	return JsonData(1, "节点列表", data)
}

func (a *App) PeerAddAPI(peerData models.Peers) map[string]any {
	if peerData.UserID < 0 {
		return JsonData(0, "节点ID错误", []any{})
	}

	err := peers.PeerAdd(&peerData)
	if err != nil {
		return JsonData(0, "添加节点失败: "+err.Error(), []any{})
	}
	return JsonData(1, "添加节点成功", []any{})
}

func (a *App) PeerUpdateAPI(peerID string, peerData models.Peers) map[string]any {
	peerIDN, err := IDtoInt(peerID)
	if err != nil {
		return JsonData(0, "节点ID错误", []any{})
	}

	if peerData.UserID < 0 {
		return JsonData(0, "节点ID错误", []any{})
	}

	peerData.CommonModel.ID = peerIDN
	err = peers.PeerUpdate(&peerData)
	if err != nil {
		return JsonData(0, "更新节点失败: "+err.Error(), []any{})
	}
	return JsonData(1, "更新节点成功", []any{})
}

func (a *App) PeerDeleteAPI(peerID string) map[string]any {
	peerIDN, err := IDtoInt(peerID)
	if err != nil {
		return JsonData(0, "节点ID错误", []any{})
	}

	err = peers.PeerDelete(peerIDN)
	if err != nil {
		return JsonData(0, "删除节点失败: "+err.Error(), []any{})
	}
	return JsonData(1, "删除节点成功", []any{})
}

// Data API
func (a *App) DataListAPI() map[string]any {
	data, err := data.PeerData()
	if err != nil {
		return JsonData(0, "节点详情获取失败: "+err.Error(), []any{})
	}
	return JsonData(1, "节点详情", data)
}

// ClipboardText
func (a *App) ClipboardText(text string) bool {
	err := runtime.ClipboardSetText(a.ctx, text)
	return err == nil
}

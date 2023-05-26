package routes

import (
	"wgm/apis/common"
	"wgm/dbs/models"
	"wgm/services/routes"

	"github.com/gin-gonic/gin"
)

func ListRoutes(c *gin.Context) {
	data, err := routes.RouteList()
	if err != nil {
		common.JSONHandler(c, 0, "路由列表获取失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "路由列表", data)
}

func AddRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var routeData models.Routes
	err := c.ShouldBindJSON(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "路由数据错误", []any{})
		return
	}
	err = routes.RouteAdd(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "添加路由失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "添加路由成功", []any{})
}

func UpdateRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	routeID := c.Param("routeID")
	routeIDN, err := common.IDtoInt(routeID)
	if err != nil {
		common.JSONHandler(c, 0, "路由ID错误", []any{})
		return
	}

	var routeData models.Routes
	err = c.ShouldBindJSON(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "路由数据错误", []any{})
		return
	}

	routeData.CommonModel.ID = routeIDN
	err = routes.RouteUpdate(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "更新路由失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "更新路由成功", []any{})
}

func DeleteRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	routeID := c.Param("routeID")
	routeIDN, err := common.IDtoInt(routeID)
	if err != nil {
		common.JSONHandler(c, 0, "路由ID错误", []any{})
		return
	}

	err = routes.RouteDelete(routeIDN)
	if err != nil {
		common.JSONHandler(c, 0, "删除路由失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "删除路由成功", []any{})
}

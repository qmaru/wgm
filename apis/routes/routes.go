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
		common.JSONHandler(c, 0, "Route List failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Route List", data)
}

func AddRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var routeData models.Routes
	err := c.ShouldBindJSON(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "Route data error", []any{})
		return
	}
	err = routes.RouteAdd(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "Add Route failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Add Route succeed", []any{})
}

func UpdateRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	routeID := c.Param("routeID")
	routeIDN, err := common.IDtoInt(routeID)
	if err != nil {
		common.JSONHandler(c, 0, "Route id error", []any{})
		return
	}

	var routeData models.Routes
	err = c.ShouldBindJSON(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "Route data error", []any{})
		return
	}

	routeData.CommonModel.ID = routeIDN
	err = routes.RouteUpdate(&routeData)
	if err != nil {
		common.JSONHandler(c, 0, "Update Route failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Update Route succeed", []any{})
}

func DeleteRoute(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	routeID := c.Param("routeID")
	routeIDN, err := common.IDtoInt(routeID)
	if err != nil {
		common.JSONHandler(c, 0, "Route id error", []any{})
		return
	}

	err = routes.RouteDelete(routeIDN)
	if err != nil {
		common.JSONHandler(c, 0, "Delete Route failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Delete Route succeed", []any{})
}

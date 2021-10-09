package apis

import (
	"wgm/models"
	"wgm/services"

	"github.com/gin-gonic/gin"
)

func ServerInfo(c *gin.Context) {
	serverID := c.Param("serverID")
	data := services.QueryServer(ParseID(serverID))
	if data.Title != "" {
		DataHandler(c, 1, "Server Info", data)
	} else {
		DataHandler(c, 0, "Server Not Found", []interface{}{})
	}
}

func ServerAdd(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var server models.Servers
	c.ShouldBindJSON(&server)

	info := models.Servers{
		Title:      server.Title,
		Address:    server.Address,
		Port:       server.Port,
		LanIP:      server.LanIP,
		LanNetmask: server.LanNetmask,
		MTU:        server.MTU,
		DNS:        server.DNS,
	}
	code := services.CreateServer(info)
	if code == services.ServerCreateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func ServerUpdate(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var server models.Servers
	c.ShouldBindJSON(&server)
	serverID := c.Param("serverID")

	code := services.UpdateServer(ParseID(serverID), server)
	if code == services.ServerUpdateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func ServerDel(c *gin.Context) {
	serverID := c.Param("serverID")
	code := services.DeleteServer(ParseID(serverID))
	if code == services.ServerDeleteSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

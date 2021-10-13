package apis

import (
	"wgm/models"
	"wgm/services"

	"github.com/gin-gonic/gin"
)

func NodeInfo(c *gin.Context) {
	serverID := c.Param("serverID")
	userID := c.Param("userID")
	data := services.QueryUser(ParseID(serverID), ParseID(userID))
	if data.CommonModel.ID != 0 {
		DataHandler(c, 1, "Node Info", data)
	} else {
		DataHandler(c, 0, "Node Not Found", []interface{}{})
	}
}

func NodeAdd(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var user models.Users
	c.ShouldBindJSON(&user)
	serverID := c.Param("serverID")

	info := models.Users{
		ServerID:            ParseID(serverID),
		Username:            user.Username,
		IP:                  user.IP,
		IsAccess:            user.IsAccess,
		IsExtra:             user.IsExtra,
		IsServer:            user.IsServer,
		PersistentKeepalive: user.PersistentKeepalive,
	}
	code := services.CreateUser(info)
	if code == services.UserCreateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func NodeUpdate(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var user models.Users
	c.ShouldBindJSON(&user)

	serverID := c.Param("serverID")
	userID := c.Param("userID")

	code := services.UpdateUser(ParseID(serverID), ParseID(userID), user)
	if code == services.UserUpdateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func NodeUpdateKey(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var user models.Users
	c.ShouldBindJSON(&user)

	serverID := c.Param("serverID")
	userID := c.Param("userID")

	code := services.UpdateUserKey(ParseID(serverID), ParseID(userID))
	if code == services.UserUpdateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func NodeDel(c *gin.Context) {
	serverID := c.Param("serverID")
	userID := c.Param("userID")
	code := services.DeleteUser(ParseID(serverID), ParseID(userID))
	if code == services.UserDeleteSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

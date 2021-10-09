package apis

import (
	"wgm/models"
	"wgm/services"

	"github.com/gin-gonic/gin"
)

func ClientInfo(c *gin.Context) {
	serverID := c.Param("serverID")
	userID := c.Param("userID")
	data := services.QueryUser(ParseID(serverID), ParseID(userID))
	if data.CommonModel.ID != 0 {
		DataHandler(c, 1, "Client Info", data)
	} else {
		DataHandler(c, 0, "Client Not Found", []interface{}{})
	}
}

func ClientAdd(c *gin.Context) {
	c.Header("Content-Typ", "application/json")
	var user models.Users
	c.ShouldBindJSON(&user)
	serverID := c.Param("serverID")

	info := models.Users{
		ServerID:            ParseID(serverID),
		Username:            user.Username,
		IP:                  user.IP,
		IsExtra:             user.IsExtra,
		PersistentKeepalive: user.PersistentKeepalive,
	}
	code := services.CreateUser(info)
	if code == services.UserCreateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func ClientUpdate(c *gin.Context) {
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

func ClientUpdateKey(c *gin.Context) {
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

func ClientDel(c *gin.Context) {
	serverID := c.Param("serverID")
	userID := c.Param("userID")
	code := services.DeleteUser(ParseID(serverID), ParseID(userID))
	if code == services.UserDeleteSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

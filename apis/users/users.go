package users

import (
	"wgm/apis/common"
	"wgm/dbs/models"
	"wgm/services/users"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	data, err := users.UserList()
	if err != nil {
		common.JSONHandler(c, 0, "User List failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "User List", data)
}

func AddUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var userData models.Users
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "User data error", []any{})
		return
	}
	err = users.UserAdd(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "Add User failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Add User succeed", []any{})
}

func UpdateUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	userID := c.Param("userID")
	userIDN, err := common.IDtoInt(userID)
	if err != nil {
		common.JSONHandler(c, 0, "User id error", []any{})
		return
	}

	var userData models.Users
	err = c.ShouldBindJSON(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "User data error", []any{})
		return
	}

	userData.CommonModel.ID = userIDN
	err = users.UserUpdate(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "Update User failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Update User succeed", []any{})
}

func DeleteUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	userID := c.Param("userID")
	userIDN, err := common.IDtoInt(userID)
	if err != nil {
		common.JSONHandler(c, 0, "User id error", []any{})
		return
	}

	err = users.UserDelete(userIDN)
	if err != nil {
		common.JSONHandler(c, 0, "Delete User failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Delete User succeed", []any{})
}

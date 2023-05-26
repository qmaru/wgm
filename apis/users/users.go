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
		common.JSONHandler(c, 0, "用户列表获取失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "用户列表", data)
}

func AddUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var userData models.Users
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "用户数据错误", []any{})
		return
	}
	err = users.UserAdd(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "添加用户失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "添加用户成功", []any{})
}

func UpdateUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	userID := c.Param("userID")
	userIDN, err := common.IDtoInt(userID)
	if err != nil {
		common.JSONHandler(c, 0, "用户ID错误", []any{})
		return
	}

	var userData models.Users
	err = c.ShouldBindJSON(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "用户数据错误", []any{})
		return
	}

	userData.CommonModel.ID = userIDN
	err = users.UserUpdate(&userData)
	if err != nil {
		common.JSONHandler(c, 0, "更新用户失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "更新用户成功", []any{})
}

func DeleteUser(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	userID := c.Param("userID")
	userIDN, err := common.IDtoInt(userID)
	if err != nil {
		common.JSONHandler(c, 0, "用户ID错误", []any{})
		return
	}

	err = users.UserDelete(userIDN)
	if err != nil {
		common.JSONHandler(c, 0, "删除用户失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "删除用户成功", []any{})
}

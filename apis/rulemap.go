package apis

import (
	"wgm/services"

	"github.com/gin-gonic/gin"
)

func RuleMapAdd(c *gin.Context) {
	ruleID := c.Param("ruleID")
	userID := c.Param("userID")
	code := services.CreateUserRule(ParseID(ruleID), ParseID(userID))
	if code == services.RuleMapCreateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func RuleMapDel(c *gin.Context) {
	ruleID := c.Param("ruleID")
	userID := c.Param("userID")
	code := services.DeleteUserRule(ParseID(ruleID), ParseID(userID))
	if code == services.RuleMapDeleteSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

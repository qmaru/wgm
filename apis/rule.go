package apis

import (
	"wgm/services"

	"github.com/gin-gonic/gin"
)

func RuleAdd(c *gin.Context) {
	allowedIP := c.Query("allowed_ip")
	code := services.CreateRule(allowedIP)
	if code == services.RuleCreateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func RuleUpdate(c *gin.Context) {
	ruleID := c.Param("ruleID")
	allowedIP := c.Query("allowed_ip")
	code := services.UpdateRule(ParseID(ruleID), allowedIP)
	if code == services.RuleUpdateSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

func RuleDel(c *gin.Context) {
	ruleID := c.Param("ruleID")
	code := services.DeleteRule(ParseID(ruleID))
	if code == services.RuleDeleteSucceed {
		DataHandler(c, 1, code.String(), []interface{}{})
	} else {
		DataHandler(c, 0, code.String(), []interface{}{})
	}
}

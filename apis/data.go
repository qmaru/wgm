package apis

import (
	"strings"

	"wgm/services"

	"github.com/gin-gonic/gin"
)

func RuleList(c *gin.Context) {
	data := services.AllRules()
	if len(data) != 0 {
		DataHandler(c, 1, "All Rules", data)
	} else {
		DataHandler(c, 0, "No Rule", []interface{}{})
	}
}

func ServerList(c *gin.Context) {
	data := services.AllServers()
	if len(data) != 0 {
		DataHandler(c, 1, "All Servers", data)
	} else {
		DataHandler(c, 0, "No Server", []interface{}{})
	}
}

func NodeList(c *gin.Context) {
	data := services.AllNodes()
	if len(data) != 0 {
		DataHandler(c, 1, "All Nodes", data)
	} else {
		DataHandler(c, 0, "No Node", []interface{}{})
	}
}

func ServerGroup(c *gin.Context) {
	data := services.AllServerUsers()
	if len(data) != 0 {
		DataHandler(c, 1, "All Server and Users", data)
	} else {
		DataHandler(c, 0, "No Data", []interface{}{})
	}
}

func UserRuleSets(c *gin.Context) {
	userID := c.Param("userID")
	rules := c.Query("rules")
	ruleIDs := strings.Split(rules, ",")
	for _, ruleID := range ruleIDs {
		services.CreateUserRule(ParseID(ruleID), ParseID(userID))
	}
	DataHandler(c, 1, "Set User Rule", []interface{}{})
}

func ShowConfig(c *gin.Context) {
	stype := c.Query("type")
	sserver := c.Query("server")
	snode := c.Query("node")

	cfg := services.ShowUserCfg(stype, sserver, snode, true, "")
	DataHandler(c, 1, "Set User Rule", cfg)
}

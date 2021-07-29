package main

import (
	"fmt"
	"testing"

	"wgm/models"
	"wgm/services"
)

func TestDBInit(t *testing.T) {
	models.InitTable()
}

func TestServerAdd(t *testing.T) {
	info := map[string]interface{}{
		"title":       "wg1",
		"address":     "wg.example.com",
		"port":        443,
		"lan_ip":      "10.0.0.1",
		"lan_netmask": "24",
		"mtu":         "",
		"dns":         "",
	}

	code := services.CreateServer(info)
	fmt.Println(code)
}

func TestServerUpdate(t *testing.T) {
	info := map[string]interface{}{
		"title":       "wg2",
		"address":     "wg2.example.com",
		"lan_ip":      "10.0.0.10",
		"lan_netmask": "24",
		"mtu":         "",
	}
	code := services.UpdateServer(info)
	fmt.Println(code)
}

func TestServerDel(t *testing.T) {
	title := "wg3"
	code := services.DeleteServer(title)
	fmt.Println(code)
}

func TestUserAdd(t *testing.T) {
	info := map[string]interface{}{
		"server_id": 1,
		"username":  "server",
		"ip":        "10.0.0.10",
	}
	code := services.CreateUser(info)
	fmt.Println(code)
}
func TestUserUpdate(t *testing.T) {
	info := map[string]interface{}{
		"server_id": 1,
		"username":  "server2",
		"keepalive": 10,
	}
	code := services.UpdateUser(info)
	fmt.Println(code)
}

func TestUserKeyUpdate(t *testing.T) {
	code := services.UpdateUserKey(1, "server2")
	fmt.Println(code)
}

func TestUserDel(t *testing.T) {
	code := services.DeleteUser(1, "server2")
	fmt.Println(code)
}

func TestRuleAdd(t *testing.T) {
	rule := "192.168.0.0/24"
	code := services.CreateRule(rule)
	fmt.Println(code)
}

func TestRuleUpdate(t *testing.T) {
	RuleID := 2
	newRule := "192.168.10.0/24"
	code := services.UpdateRule(RuleID, newRule)
	fmt.Println(code)
}

func TestRuleDel(t *testing.T) {
	ruleID := 1
	code := services.DeleteRule(ruleID)
	fmt.Println(code)
}

func TestRulemapAdd(t *testing.T) {
	ruleID := 2
	userID := 1
	code := services.CreateUserRule(userID, ruleID)
	fmt.Println(code)
}

func TestRulemapDel(t *testing.T) {
	ruleID := 1
	userID := 2
	code := services.DeleteUserRule(userID, ruleID)
	fmt.Println(code)
}

func TestShowUserList(t *testing.T) {
	services.ShowUserList(0)
}

func TestShowServerList(t *testing.T) {
	services.ShowServerList()
}

func TestShowRuleList(t *testing.T) {
	services.ShowRuleList()
}

func TestShowUserCfg(t *testing.T) {
	serverCfg := services.ShowUserCfg(1, 1, "server", "", true, "")
	fmt.Println(serverCfg)

	clientCfg := services.ShowUserCfg(2, 1, "server", "server", true, "")
	fmt.Println(clientCfg)
}

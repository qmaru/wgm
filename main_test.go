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
	info := models.Servers{
		Title:      "wg1",
		Address:    "wg.example.com",
		Port:       443,
		LanIP:      "10.0.0.1",
		LanNetmask: "24",
		MTU:        "",
		DNS:        "",
	}

	code := services.CreateServer(info)
	fmt.Println(code)
}

func TestServerUpdate(t *testing.T) {
	info := models.Servers{
		Title:      "wg2",
		Address:    "wg2.example.com",
		LanIP:      "10.0.0.1",
		LanNetmask: "24",
	}

	code := services.UpdateServer(1, info)
	fmt.Println(code)
}

func TestServerDel(t *testing.T) {
	title := "wg3"
	serverID := services.GetServerID(title)
	code := services.DeleteServer(serverID)
	fmt.Println(code)
}

func TestUserAdd(t *testing.T) {
	info := models.Users{
		ServerID: 1,
		Username: "server",
		IP:       "10.0.0.10",
	}
	code := services.CreateUser(info)
	fmt.Println(code)
}
func TestUserUpdate(t *testing.T) {
	info := models.Users{
		ServerID:            1,
		Username:            "server2",
		PersistentKeepalive: 10,
	}
	code := services.UpdateUser(1, 1, info)
	fmt.Println(code)
}

func TestUserKeyUpdate(t *testing.T) {
	code := services.UpdateUserKey(1, 1)
	fmt.Println(code)
}

func TestUserDel(t *testing.T) {
	code := services.DeleteUser(1, 1)
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
	services.ShowUserList(0, false)
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

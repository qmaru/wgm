package services

import (
	"fmt"
	"sort"

	"wgm/models"
)

func allDataGroupByServer(sData []map[string]interface{}) (dData []map[string]interface{}) {
	i, j := 0, 0
	groupData := make([][]map[string]interface{}, 0)

	sort.SliceStable(sData, func(i, j int) bool {
		return sData[i]["server_id"].(int) < sData[j]["server_id"].(int)
	})

	for {
		if i >= len(sData) {
			break
		}

		for j = i; j < len(sData) && sData[i]["server_id"] == sData[j]["server_id"]; j++ {
		}

		groupData = append(groupData, sData[i:j])
		i = j
	}

	for _, gData := range groupData {
		mData := make(map[string]interface{})
		mDetails := make([]map[string]interface{}, 0)
		for _, gd := range gData {
			mmData := make(map[string]interface{})
			mmData["user_id"] = gd["user_id"]
			mmData["user_name"] = gd["user_name"]
			mmData["user_lan"] = gd["user_lan"]
			mmData["user_is_server"] = gd["user_is_server"]
			mmData["user_prikey"] = gd["user_prikey"]
			mmData["user_pubkey"] = gd["user_pubkey"]
			mmData["user_keepalive"] = gd["user_keepalive"]

			mData["server_id"] = gd["server_id"]
			mData["server_title"] = gd["server_title"]
			mData["server_endpoint"] = gd["server_endpoint"]
			mData["server_lan"] = gd["server_lan"]

			mDetails = append(mDetails, mmData)
		}
		mData["users"] = mDetails
		dData = append(dData, mData)
	}
	return
}

func AllServerUsers() []map[string]interface{} {
	allQuery := `
	SELECT
		servers.id as serverID,
		servers.title as serverTitle,
		servers.address as serverAddress,
		servers.port as serverPort,
		servers.lan_ip as serverLan,
		servers.lan_netmask as serverNM,
		users.id as userID,
		users.username as username,
		users.ip as userLan,
		users.prikey as userPrikey,
		users.pubkey as userPubkey,
		users.is_server as userIsServer,
		users.keepalive as userKeepalive
	FROM
		servers
	LEFT JOIN users on users.server_id = servers.id
	WHERE users.status = 1 and servers.status = 1
	`
	allRows, _ := models.DBQuery(allQuery)
	defer allRows.Close()

	var allData []map[string]interface{}
	for allRows.Next() {
		var serverID int
		var serverTitle string
		var serverAddress string
		var serverPort int
		var serverIP string
		var serverNM string
		var userID int
		var username string
		var userIP string
		var userPrikey string
		var userPubkey string
		var userIsServer int
		var userKeepalive int

		aData := make(map[string]interface{})
		allRows.Scan(&serverID, &serverTitle, &serverAddress, &serverPort, &serverIP, &serverNM, &userID, &username, &userIP, &userPrikey, &userPubkey, &userIsServer, &userKeepalive)
		aData["server_id"] = serverID
		aData["server_title"] = serverTitle
		aData["server_endpoint"] = fmt.Sprintf("%s:%d", serverAddress, serverPort)
		aData["server_lan"] = fmt.Sprintf("%s/%s", serverIP, serverNM)
		aData["user_id"] = userID
		aData["user_name"] = username
		aData["user_lan"] = userIP
		aData["user_prikey"] = userPrikey
		aData["user_pubkey"] = userPubkey
		aData["user_is_server"] = userIsServer
		aData["user_keepalive"] = userKeepalive
		allData = append(allData, aData)
	}
	return allDataGroupByServer(allData)
}

func AllRules() []map[string]interface{} {
	allQuery := `
	SELECT
		id,
		allowedip
	FROM
		rules
	WHERE
		status = 1
	`
	allRows, _ := models.DBQuery(allQuery)
	defer allRows.Close()

	var allRules []map[string]interface{}
	for allRows.Next() {
		var ruleID int
		var ruleIP string
		allRows.Scan(&ruleID, &ruleIP)
		allRules = append(allRules, map[string]interface{}{
			"id": ruleID,
			"ip": ruleIP,
		})
	}
	return allRules
}

func AllServers() []map[string]interface{} {
	allQuery := `
	SELECT
		id,
		title,
		address,
		port,
		lan_ip,
		lan_netmask,
		mtu,
		dns
	FROM
		servers
	WHERE
		status = 1
	`
	allRows, _ := models.DBQuery(allQuery)
	defer allRows.Close()

	var allServers []map[string]interface{}
	for allRows.Next() {
		var serverID int
		var serverTitle string
		var serverAddress string
		var serverPort int
		var serverIP string
		var serverNM string
		var serverMTU string
		var serverDNS string

		allRows.Scan(&serverID, &serverTitle, &serverAddress, &serverPort, &serverIP, &serverNM, &serverMTU, &serverDNS)
		allServers = append(allServers, map[string]interface{}{
			"id":          serverID,
			"title":       serverTitle,
			"address":     serverAddress,
			"port":        serverPort,
			"lan_ip":      serverIP,
			"lan_netmask": serverNM,
			"mtu":         serverMTU,
			"dns":         serverDNS,
		})
	}
	return allServers
}

func AllNodes() []map[string]interface{} {
	allQuery := `
	SELECT
		users.id,
		servers.id as server_id,
		servers.title as server_title,
		users.username,
		users.prikey,
		users.pubkey,
		users.ip,
		users.default_rule,
		users.is_access,
		users.is_extra,
		users.is_server,
		users.keepalive
	FROM
		users
	LEFT JOIN servers on users.server_id = servers.id
	WHERE
		users.status = 1 and servers.status = 1
	`
	allRows, _ := models.DBQuery(allQuery)
	defer allRows.Close()

	var allNodes []map[string]interface{}
	for allRows.Next() {
		var nodeID int
		var serverID int
		var serverTitle string
		var nodeUsername string
		var nodePrikey string
		var nodePubkey string
		var nodeIP string
		var nodeDefaultRule string
		var nodeIsAccess int
		var nodeIsExtra int
		var nodeIsServer int
		var nodeKeepalive int

		allRows.Scan(&nodeID, &serverID, &serverTitle, &nodeUsername, &nodePrikey, &nodePubkey, &nodeIP, &nodeDefaultRule, &nodeIsAccess, &nodeIsExtra, &nodeIsServer, &nodeKeepalive)
		allNodes = append(allNodes, map[string]interface{}{
			"id":           nodeID,
			"server_id":    serverID,
			"server_title": serverTitle,
			"username":     nodeUsername,
			"prikey":       nodePrikey,
			"pubkey":       nodePubkey,
			"ip":           nodeIP,
			"default_rule": nodeDefaultRule,
			"is_access":    nodeIsAccess,
			"is_extra":     nodeIsExtra,
			"is_server":    nodeIsServer,
			"keepalive":    nodeKeepalive,
		})
	}
	return allNodes
}

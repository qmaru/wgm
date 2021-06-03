package services

import (
	"fmt"
	"strconv"
	"time"

	"wgm/models"
)

func ServerExist() bool {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1", models.ServersTable)
	row := models.DBQueryOne(sql)
	var sid int
	row.Scan(&sid)
	return sid != 0
}

func ServerCheck(title string) int {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and title=?", models.ServersTable)
	row := models.DBQueryOne(sql, title)
	var sid int
	row.Scan(&sid)
	return sid
}

// CreateServer 创建服务器
func CreateServer(info map[string]interface{}) statusCode {
	// 收集数据
	serverTitleRaw, ok := info["title"]
	if !ok {
		return serverTitleRequired
	}

	serverTitle := serverTitleRaw.(string)
	serverAddress := info["address"].(string)
	serverPort := info["port"].(int)
	serverLanIP := info["lan_ip"].(string)
	serverLanNetmask := info["lan_netmask"].(string)

	var serverMTU int
	var serverDNS string
	if _, ok := info["mtu"]; ok {
		serverMTU, _ = strconv.Atoi(info["mtu"].(string))
	}

	if _, ok := info["dns"]; ok {
		serverDNS = info["dns"].(string)
	}
	// 检查数据
	if serverID := ServerCheck(serverTitle); serverID != 0 {
		return serverHasExist
	}

	serverLAN := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
	if ok, serverIP, serverNetmask := LanValidator(serverLAN); ok {
		createdat := time.Now().Unix()
		updatedat := time.Now().Unix()

		sqlInsert := fmt.Sprintf("INSERT INTO %s (created_at,updated_at,title,address,port,lan_ip,lan_netmask,mtu,dns) VALUES (?,?,?,?,?,?,?,?,?)", models.ServersTable)
		models.DBExec(
			sqlInsert,
			createdat,
			updatedat,
			serverTitle,
			serverAddress,
			serverPort,
			serverIP,
			serverNetmask,
			serverMTU,
			serverDNS,
		)
		return serverCreateSucceed
	}
	return serverIPError
}

// UpdateServer 更新服务器
func UpdateServer(info map[string]interface{}) statusCode {
	// 生成更新字段
	serverTitleRaw, ok := info["title"]
	if !ok {
		return serverTitleRequired
	}

	serverTitle := serverTitleRaw.(string)
	// 检查数据
	if serverID := ServerCheck(serverTitle); serverID != 0 {
		var serverLanIP string
		var serverLanNetmask string
		var serverLanCheck bool
		_, ok1 := info["lan_ip"]
		_, ok2 := info["lan_netmask"]

		if ok1 && ok2 {
			serverLanIP = info["lan_ip"].(string)
			serverLanNetmask = info["lan_netmask"].(string)
			serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
			serverLanCheck, _, _ = LanValidator(serverLan)
			if !serverLanCheck {
				return serverIPError
			}
		}

		delete(info, "title")
		updateSets := GenUpdate(info)
		updatedat := time.Now().Unix()
		sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,%s WHERE id=? and status=1", models.ServersTable, updateSets)

		models.DBExec(
			sqlUpdate,
			updatedat,
			serverID,
		)
		return serverUpdateSucceed
	}
	return serverNotFound
}

// QueryServer 查找服务器
func QueryServer(sid int) (data map[string]interface{}) {
	var serverID int
	var serverTitle string
	var serverAddress string
	var serverPort int
	var serverLanIP string
	var serverLanNetmask string
	var sreverMTU string
	var srerverDNS string

	serverSelect := fmt.Sprintf("SELECT id,title,address,port,lan_ip,lan_netmask,mtu,dns FROM %s WHERE status=1 and id=?", models.ServersTable)
	row := models.DBQueryOne(serverSelect, sid)
	row.Scan(
		&serverID,
		&serverTitle,
		&serverAddress,
		&serverPort,
		&serverLanIP,
		&serverLanNetmask,
		&sreverMTU,
		&srerverDNS,
	)
	if serverID != 0 {
		data = map[string]interface{}{
			"id":          serverID,
			"title":       serverTitle,
			"address":     serverAddress,
			"port":        serverPort,
			"lan_ip":      serverLanIP,
			"lan_netmask": serverLanNetmask,
			"mtu":         sreverMTU,
			"dns":         srerverDNS,
		}
	}
	return
}

// DeleteServer 删除服务器
func DeleteServer(title string) statusCode {
	if serverID := ServerCheck(title); serverID != 0 {
		sqlDelete := fmt.Sprintf("UPDATE %s SET status=0 WHERE id=? and status=1", models.ServersTable)
		models.DBExec(sqlDelete, serverID)
		return serverDeleteSucceed
	}
	return serverNotFound
}

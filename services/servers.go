package services

import (
	"fmt"
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

// CreateServer 创建服务端
func CreateServer(server models.Servers) statusCode {
	// 必须有服务端标识
	serverTitle := server.Title
	if server.Title == "" {
		return ServerTitleRequired
	}

	serverAddress := server.Address
	serverPort := server.Port
	serverLanIP := server.LanIP
	serverLanNetmask := server.LanNetmask
	serverMTU := server.MTU
	serverDNS := server.DNS

	// 检查数据
	if serverID := GetServerID(serverTitle); serverID != 0 {
		return ServerHasExist
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
		return ServerCreateSucceed
	}
	return ServerIPError
}

// UpdateServer 更新服务器
func UpdateServer(id int, server models.Servers) statusCode {
	if serverID := GetServerID(id); serverID != 0 {
		data := make(map[string]interface{})
		serverLanIP := server.LanIP
		serverLanNetmask := server.LanNetmask

		if serverLanIP != "" && serverLanNetmask != "" {
			serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
			serverLanCheck, _, _ := LanValidator(serverLan)
			if !serverLanCheck {
				return ServerIPError
			} else {
				data["lan_ip"] = serverLanIP
				data["lan_netmask"] = serverLanNetmask
			}
		}

		if server.Title != "" {
			data["title"] = server.Title
		}

		if server.Address != "" {
			data["address"] = server.Address
		}

		if server.Port != 0 {
			data["port"] = server.Port
		}

		data["mtu"] = server.MTU
		data["dns"] = server.DNS

		updateSets := GenUpdate(data)
		updatedat := time.Now().Unix()
		sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,%s WHERE id=? and status=1", models.ServersTable, updateSets)
		models.DBExec(
			sqlUpdate,
			updatedat,
			serverID,
		)
		return ServerUpdateSucceed
	}
	return ServerNotFound
}

// QueryServer 查找服务器
func QueryServer(id int) (data models.Servers) {
	var serverID int
	var serverTitle string
	var serverAddress string
	var serverPort int
	var serverLanIP string
	var serverLanNetmask string
	var sreverMTU string
	var srerverDNS string

	serverSelect := fmt.Sprintf("SELECT id,title,address,port,lan_ip,lan_netmask,mtu,dns FROM %s WHERE status=1 and id=?", models.ServersTable)
	row := models.DBQueryOne(serverSelect, id)
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
		data = models.Servers{
			CommonModel: models.CommonModel{
				ID: serverID,
			},
			Title:      serverTitle,
			Address:    serverAddress,
			Port:       serverPort,
			LanIP:      serverLanIP,
			LanNetmask: serverLanNetmask,
			MTU:        sreverMTU,
			DNS:        srerverDNS,
		}
		return
	}
	return
}

// DeleteServer 删除服务器
func DeleteServer(id int) statusCode {
	if serverID := GetServerID(id); serverID != 0 {
		sqlDelete := fmt.Sprintf("UPDATE %s SET status=0 WHERE id=? and status=1", models.ServersTable)
		models.DBExec(sqlDelete, serverID)
		return ServerDeleteSucceed
	}
	return ServerNotFound
}

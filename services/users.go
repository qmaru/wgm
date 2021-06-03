package services

import (
	"fmt"
	"time"

	"wgm/models"
)

func UserCheck(serverID int, username string) int {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and username=? and server_id=?", models.UsersTable)
	row := models.DBQueryOne(sql, username, serverID)
	var uid int
	row.Scan(&uid)
	return uid
}

// CreateUser 创建用户
func CreateUser(info map[string]interface{}) statusCode {
	if !ServerExist() {
		return serverNotFound
	}

	userServerIDRaw, ok1 := info["server_id"]
	userUsernameRaw, ok2 := info["username"]

	if !ok1 {
		return serverNotFound
	}

	if !ok2 {
		return userNotFound
	}

	userServerID := userServerIDRaw.(int)
	userUsername := userUsernameRaw.(string)

	if UserCheck(userServerID, userUsername) != 0 {
		return userHasExist
	}

	serverData := QueryServer(userServerID)
	if len(serverData) != 0 {
		serverLanIP := serverData["lan_ip"].(string)
		serverLanNetmask := serverData["lan_netmask"].(string)
		serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
		userIP := info["ip"].(string)

		if !IPcheck(userIP, serverLan) {
			return userIPError
		}

		prikey, _ := GeneratePrivateKey()
		pubkey := prikey.PublicKey()

		createdat := time.Now().Unix()
		updatedat := time.Now().Unix()
		userPrikey := prikey.String()
		userPubkey := pubkey.String()
		userIPLan := fmt.Sprintf("%s/%s", userIP, "24")
		userDefaultRule := fmt.Sprintf("%s/%s", userIP, "32")
		userIsExtra := info["is_extra"].(int)

		sqlInsert := fmt.Sprintf("INSERT INTO %s (created_at,updated_at,server_id,username,prikey,pubkey,ip,default_rule,is_extra) VALUES (?,?,?,?,?,?,?,?,?)", models.UsersTable)
		models.DBExec(
			sqlInsert,
			createdat,
			updatedat,
			userServerID,
			userUsername,
			userPrikey,
			userPubkey,
			userIPLan,
			userDefaultRule,
			userIsExtra,
		)
		return userCreateSucceed
	}
	return serverNotFound
}

// UpdateUser 修改用户
func UpdateUser(info map[string]interface{}) statusCode {
	if !ServerExist() {
		return serverNotFound
	}

	userServerIDRaw, ok1 := info["server_id"]
	userUsernameRaw, ok2 := info["username"]

	if !ok1 {
		return serverNotFound
	}

	if !ok2 {
		return userNotFound
	}

	userServerID := userServerIDRaw.(int)
	userUsername := userUsernameRaw.(string)

	if uid := UserCheck(userServerID, userUsername); uid != 0 {
		serverData := QueryServer(userServerID)
		if len(serverData) != 0 {
			var userIP string
			var userIPLan string
			var userDefaultRule string

			if userIPRaw, ok := info["ip"]; ok {
				userIP = userIPRaw.(string)
				serverLanIP := serverData["lan_ip"].(string)
				serverLanNetmask := serverData["lan_netmask"].(string)
				serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
				if !IPcheck(userIP, serverLan) {
					return userIPError
				}
				userIPLan = fmt.Sprintf("%s/%s", userIP, "24")
				userDefaultRule = fmt.Sprintf("%s/%s", userIP, "32")
			}

			updatedat := time.Now().Unix()

			if userIPLan != "" {
				info["ip"] = userIPLan
				info["default_rule"] = userDefaultRule
			}

			delete(info, "server_id")
			delete(info, "username")

			updateSets := GenUpdate(info)
			if updateSets != "" {
				sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,%s WHERE id=? and status=1", models.UsersTable, updateSets)
				models.DBExec(
					sqlUpdate,
					updatedat,
					uid,
				)
			}
			return userUpdateSucceed
		}
	}
	return userNotFound
}

// UpdateUserKey 更换用户密钥
func UpdateUserKey(serverID int, username string) statusCode {
	if uid := UserCheck(serverID, username); uid != 0 {
		updateat := time.Now().Unix()
		prikey, _ := GeneratePrivateKey()
		pubkey := prikey.PublicKey()

		userPrikey := prikey.String()
		userPubkey := pubkey.String()

		sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,prikey=?,pubkey=? WHERE id=? and status=1", models.UsersTable)
		models.DBExec(sqlUpdate, updateat, userPrikey, userPubkey, uid)
		return userUpdateSucceed
	}
	return userNotFound
}

// DeleteUser 删除用户
func DeleteUser(serverID int, username string) statusCode {
	if uid := UserCheck(serverID, username); uid != 0 {
		sqlDelete := fmt.Sprintf("UPDATE %s SET status=0 WHERE id=? and status=1", models.UsersTable)
		models.DBExec(sqlDelete, uid)
		return userDeleteSucceed
	}
	return userNotFound
}

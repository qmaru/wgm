package services

import (
	"fmt"
	"time"

	"wgm/models"
)

func UserCheck(serverID, userID int) int {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and id=? and server_id=?", models.UsersTable)
	row := models.DBQueryOne(sql, userID, serverID)
	var uid int
	row.Scan(&uid)
	return uid
}

func UserIPCheck(serverID int, lanIP string) bool {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and server_id=? and ip=? ", models.UsersTable)
	row := models.DBQueryOne(sql, serverID, lanIP)
	var uid int
	row.Scan(&uid)
	return uid != 0
}

// CreateUser 创建用户
func CreateUser(user models.Users) statusCode {
	if !ServerExist() {
		return ServerNotFound
	}

	userServerID := user.ServerID
	userUsername := user.Username
	userIsServer := user.IsServer

	userID := GetUserID(userUsername)

	if userIsServer == 1 {
		if GetPeerID(userServerID) != 0 {
			return UserCenterExist
		}
	}

	if UserCheck(userServerID, userID) != 0 {
		return UserHasExist
	}

	serverData := QueryServer(userServerID)
	if serverData.CommonModel.ID != 0 {
		serverLanIP := serverData.LanIP
		serverLanNetmask := serverData.LanNetmask
		serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
		userIP := user.IP

		if !IPcheck(userIP, serverLan) {
			return UserIPError
		}

		userIPLan := fmt.Sprintf("%s/%s", userIP, serverLanNetmask)

		if UserIPCheck(userServerID, userIPLan) {
			return UserIPDuplicate
		}

		prikey, _ := GeneratePrivateKey()
		pubkey := prikey.PublicKey()

		createdat := time.Now().Unix()
		updatedat := time.Now().Unix()
		userPrikey := prikey.String()
		userPubkey := pubkey.String()

		userDefaultRule := fmt.Sprintf("%s/%s", userIP, "32")
		userIsAccess := user.IsAccess
		userIsExtra := user.IsExtra
		userKeepalive := user.PersistentKeepalive

		sqlInsert := fmt.Sprintf("INSERT INTO %s (created_at,updated_at,server_id,username,prikey,pubkey,ip,default_rule,is_access,is_extra,is_server,keepalive) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)", models.UsersTable)
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
			userIsAccess,
			userIsExtra,
			userIsServer,
			userKeepalive,
		)
		return UserCreateSucceed
	}
	return ServerNotFound
}

// UpdateUser 修改用户
func UpdateUser(serverID, userID int, user models.Users) statusCode {
	if !ServerExist() {
		return ServerNotFound
	}

	if uid := UserCheck(serverID, userID); uid != 0 {
		serverData := QueryServer(serverID)
		if serverData.CommonModel.ID != 0 {
			data := make(map[string]interface{})
			userIP := user.IP

			var userIPLan string
			var userDefaultRule string

			if userIP != "" {
				serverLanIP := serverData.LanIP
				serverLanNetmask := serverData.LanNetmask
				serverLan := fmt.Sprintf("%s/%s", serverLanIP, serverLanNetmask)
				if !IPcheck(userIP, serverLan) {
					return UserIPError
				}
				userIPLan = fmt.Sprintf("%s/%s", userIP, serverLanNetmask)
				userDefaultRule = fmt.Sprintf("%s/%s", userIP, "32")
			}

			if userIPLan != "" {
				data["ip"] = userIPLan
				data["default_rule"] = userDefaultRule
			}

			if user.PersistentKeepalive != 0 {
				data["keepalive"] = user.PersistentKeepalive
			} else {
				data["keepalive"] = 25
			}

			if user.Username != "" {
				data["username"] = user.Username
			}

			data["is_access"] = user.IsAccess
			data["is_extra"] = user.IsExtra
			data["is_server"] = user.IsServer

			updatedat := time.Now().Unix()
			updateSets := GenUpdate(data)
			if updateSets != "" {
				sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,%s WHERE id=? and status=1", models.UsersTable, updateSets)
				models.DBExec(
					sqlUpdate,
					updatedat,
					uid,
				)
			}
			return UserUpdateSucceed
		}
	}
	return UserNotFound
}

// UpdateUserKey 更换用户密钥
func UpdateUserKey(serverID, userID int) statusCode {
	if uid := UserCheck(serverID, userID); uid != 0 {
		updateat := time.Now().Unix()
		prikey, _ := GeneratePrivateKey()
		pubkey := prikey.PublicKey()

		userPrikey := prikey.String()
		userPubkey := pubkey.String()

		sqlUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,prikey=?,pubkey=? WHERE id=? and status=1", models.UsersTable)
		models.DBExec(sqlUpdate, updateat, userPrikey, userPubkey, uid)
		return UserUpdateSucceed
	}
	return UserNotFound
}

func QueryUser(serverID, userID int) (data models.Users) {
	if uid := UserCheck(serverID, userID); uid != 0 {
		var userName string
		var userPriKey string
		var userPubkey string
		var userIP string
		var userDefaultRule string
		var userIsAccess int
		var userIsExtra int
		var userIsServer int
		var userPersistentKeepalive int

		userSelect := fmt.Sprintf("SELECT id,server_id,username,prikey,pubkey,ip,default_rule,is_access,is_extra,is_server,keepalive FROM %s WHERE status=1 and server_id=? and id=?", models.UsersTable)
		row := models.DBQueryOne(userSelect, serverID, userID)
		row.Scan(
			&userID,
			&serverID,
			&userName,
			&userPriKey,
			&userPubkey,
			&userIP,
			&userDefaultRule,
			&userIsAccess,
			&userIsExtra,
			&userIsServer,
			&userPersistentKeepalive,
		)
		if userID != 0 {
			data = models.Users{
				CommonModel: models.CommonModel{
					ID: userID,
				},
				ServerID:            serverID,
				Username:            userName,
				PriKey:              userPriKey,
				Pubkey:              userPubkey,
				IP:                  userIP,
				DefaultRule:         userDefaultRule,
				IsAccess:            userIsAccess,
				IsExtra:             userIsExtra,
				IsServer:            userIsServer,
				PersistentKeepalive: userPersistentKeepalive,
			}
			return
		}
	}
	return
}

// DeleteUser 删除用户
func DeleteUser(serverID, userID int) statusCode {
	if uid := UserCheck(serverID, userID); uid != 0 {
		sqlDelete := fmt.Sprintf("UPDATE %s SET status=0 WHERE id=? and status=1", models.UsersTable)
		models.DBExec(sqlDelete, uid)
		return UserDeleteSucceed
	}
	return UserNotFound
}

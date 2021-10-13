package services

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"wgm/models"
)

func GetPeerID(serverID int) int {
	var userID int
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and is_server=1 and server_id=?", models.UsersTable)
	row := models.DBQueryOne(sql, serverID)
	row.Scan(&userID)
	return userID
}

func GetServerID(i interface{}) int {
	checkType := reflect.TypeOf(i).String()
	var serverID int
	if checkType == "string" {
		sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and title=?", models.ServersTable)
		row := models.DBQueryOne(sql, i)
		row.Scan(&serverID)
	} else if checkType == "int" {
		sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and id=?", models.ServersTable)
		row := models.DBQueryOne(sql, i)
		row.Scan(&serverID)
	}
	return serverID
}

func GetUserID(i interface{}) int {
	checkType := reflect.TypeOf(i).String()
	var userID int
	if checkType == "string" {
		sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and username=?", models.UsersTable)
		row := models.DBQueryOne(sql, i)
		row.Scan(&userID)
	} else if checkType == "int" {
		sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and id=?", models.UsersTable)
		row := models.DBQueryOne(sql, i)
		row.Scan(&userID)
	}
	return userID
}

// LanValidator 检查 IP 和 掩码 是否符合规范
func LanValidator(cidr string) (bool, string, string) {
	ip, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, "", ""
	}
	ipStr := ip.String()
	subnetStr := subnet.String()
	ipAndNetmask := strings.Split(subnetStr, "/")
	netmask := ipAndNetmask[1]
	return true, ipStr, netmask
}

// IPCheck 检查 IP 是否在划分的范围内
func IPcheck(ip string, cidr string) bool {
	_, subnet, _ := net.ParseCIDR(cidr)
	nip := net.ParseIP(ip)
	return subnet.Contains(nip)
}

func GenUpdate(data map[string]interface{}) string {
	updateStr := []string{}
	for k, v := range data {
		var kv string
		if v != " " {
			switch v.(type) {
			case string:
				kv = fmt.Sprintf(`%s='%s'`, k, v)
			case int:
				kv = fmt.Sprintf(`%s='%d'`, k, v)
			}
		} else {
			kv = fmt.Sprintf(`%s=''`, k)
		}
		updateStr = append(updateStr, kv)
	}
	return strings.Join(updateStr, ",")
}

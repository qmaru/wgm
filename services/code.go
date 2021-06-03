package services

import (
	"fmt"
	"net"
	"strings"
)

type statusCode int

const (
	userCreateSucceed statusCode = iota
	userUpdateSucceed
	userDeleteSucceed
	userNotFound
	userHasExist
	userIPError
	serverCreateSucceed
	serverUpdateSucceed
	serverDeleteSucceed
	serverNotFound
	serverTitleRequired
	serverHasExist
	serverIPError
	ruleCreateSucceed
	ruleUpdateSucceed
	ruleDeleteSucceed
	ruleNotFound
	ruleHasExist
	ruleIPError
	ruleMapCreateSucceed
	ruleMapUpdateSucceed
	ruleMapDeleteSucceed
	ruleMapNotFound
	ruleMapHasExist
)

func (code statusCode) String() string {
	switch code {
	case userCreateSucceed:
		return "Create User Succeed"
	case userUpdateSucceed:
		return "Update User Succeed"
	case userDeleteSucceed:
		return "Delete User Succeed"
	case userNotFound:
		return "User Not Found"
	case userHasExist:
		return "User Has Exist"
	case userIPError:
		return "User IP Error"
	case serverCreateSucceed:
		return "Create Server Succeed"
	case serverUpdateSucceed:
		return "Update Server Succeed"
	case serverDeleteSucceed:
		return "Delete Server Succeed"
	case serverNotFound:
		return "Server Not Found"
	case serverTitleRequired:
		return "Server Title Required"
	case serverHasExist:
		return "Server Has Exist"
	case serverIPError:
		return "Server IP Error"
	case ruleCreateSucceed:
		return "Create Rule Succeed"
	case ruleUpdateSucceed:
		return "Update Rule Succeed"
	case ruleDeleteSucceed:
		return "Delete Rule Succeed"
	case ruleNotFound:
		return "Rule Not Found"
	case ruleHasExist:
		return "Rule Has Exist"
	case ruleIPError:
		return "Rule IP Error"
	case ruleMapCreateSucceed:
		return "Create RuleMap Succeed"
	case ruleMapUpdateSucceed:
		return "Update RuleMap Succeed"
	case ruleMapDeleteSucceed:
		return "Delete RuleMap Succeed"
	case ruleMapNotFound:
		return "RuleMap Not Found"
	case ruleMapHasExist:
		return "RuleMap Has Exist"
	}
	return "codeError"
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

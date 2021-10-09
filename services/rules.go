package services

import (
	"fmt"
	"net"
	"time"

	"wgm/models"
)

func RuleCheck(allowedip string) int {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and allowedip=?", models.RulesTable)
	row := models.DBQueryOne(sql, allowedip)
	var rid int
	row.Scan(&rid)
	return rid
}

func AllowedCheck(allowedip string) (bool, string) {
	_, subnet, err := net.ParseCIDR(allowedip)
	if err != nil {
		return false, ""
	}
	return true, subnet.String()
}

func CreateRule(allowedip string) statusCode {
	if ok, subnet := AllowedCheck(allowedip); ok {
		if rid := RuleCheck(subnet); rid != 0 {
			return RuleHasExist
		}
		createdat := time.Now().Unix()
		updatedat := time.Now().Unix()

		ruleInsert := fmt.Sprintf("INSERT INTO %s (created_at,updated_at,allowedip) VALUES (?,?,?)", models.RulesTable)
		models.DBExec(ruleInsert, createdat, updatedat, subnet)
		return RuleCreateSucceed
	}
	return RuleIPError
}

func UpdateRule(ruleID int, newAllowedip string) statusCode {
	var ok bool
	var newSubnet string

	if ok, newSubnet = AllowedCheck(newAllowedip); !ok {
		return RuleIPError
	}

	if IDCheck(ruleID, models.RulesTable) {
		updatedat := time.Now().Unix()
		ruleUpdate := fmt.Sprintf("UPDATE %s SET updated_at=?,allowedip=? WHERE id=? and status=1", models.RulesTable)
		models.DBExec(ruleUpdate, updatedat, newSubnet, ruleID)
		return RuleUpdateSucceed
	}
	return RuleNotFound
}

func DeleteRule(ruleID int) statusCode {
	if IDCheck(ruleID, models.RulesTable) {
		ruleDelete := fmt.Sprintf("UPDATE %s SET status=0 WHERE id=? and status=1", models.RulesTable)
		models.DBExec(ruleDelete, ruleID)
		return RuleDeleteSucceed
	}
	return RuleNotFound
}

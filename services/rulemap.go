package services

import (
	"fmt"
	"time"

	"wgm/models"
)

func IDCheck(id int, table string) bool {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and id=?", table)
	row := models.DBQueryOne(sql, id)
	var cid int
	row.Scan(&cid)
	return cid != 0
}

func RuleMapCheck(userID, ruleID int) int {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE status=1 and user_id=? and rule_id=?", models.RulemapTable)
	row := models.DBQueryOne(sql, userID, ruleID)
	var mid int
	row.Scan(&mid)
	return mid
}

func CreateUserRule(userID, ruleID int) statusCode {
	uCheck := IDCheck(userID, models.UsersTable)
	rCheck := IDCheck(ruleID, models.RulesTable)

	if !uCheck {
		return userNotFound
	}

	if !rCheck {
		return ruleNotFound
	}

	if rmID := RuleMapCheck(userID, ruleID); rmID != 0 {
		return ruleMapHasExist
	}

	createdat := time.Now().Unix()
	updatedat := time.Now().Unix()

	sqlInsert := fmt.Sprintf("INSERT INTO %s (created_at,updated_at,user_id,rule_id) VALUES (?,?,?,?)", models.RulemapTable)
	models.DBExec(sqlInsert, createdat, updatedat, userID, ruleID)
	return ruleMapCreateSucceed
}

func DeleteUserRule(userID, ruleID int) statusCode {
	if rmID := RuleMapCheck(userID, ruleID); rmID != 0 {
		sqlUpdate := fmt.Sprintf("UPDATE %s SET status=0 WHERE status=1 and id=?", models.RulemapTable)
		models.DBExec(sqlUpdate, rmID)
		return ruleMapDeleteSucceed
	}
	return ruleMapNotFound
}

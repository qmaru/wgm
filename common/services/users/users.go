package users

import (
	"fmt"
	"time"

	"wgm/common/dbs"
	"wgm/common/dbs/models"
)

func UserIDCheck(userID int) (int, error) {
	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE id=? AND state=1`, dbs.UserTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, userID)
	if err != nil {
		return 0, err
	}

	var uid int
	err = row.Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func UserList() ([]map[string]any, error) {
	selectSql := fmt.Sprintf(`SELECT id,username,prikey,pubkey FROM %s WHERE state=1`, dbs.UserTable)
	rows, err := dbs.Sqlite.Query(selectSql)
	if err != nil {
		return nil, err
	}

	var uId int
	var uUsername string
	var uPrikey string
	var uPubkey string

	data := make([]map[string]any, 0)
	for rows.Next() {
		err = rows.Scan(&uId, &uUsername, &uPrikey, &uPubkey)
		if err != nil {
			return nil, err
		}
		tmp := make(map[string]any)
		tmp["id"] = uId
		tmp["username"] = uUsername
		tmp["private_key"] = uPrikey
		tmp["public_key"] = uPubkey
		data = append(data, tmp)
	}
	return data, nil
}

func UserAdd(userData *models.Users) error {
	username := userData.Username

	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE username=? AND state=1`, dbs.UserTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, username)
	if err != nil {
		return err
	}
	var uid int
	err = row.Scan(&uid)
	if err != nil {
		privateKeySuit, err := GeneratePrivateKey()
		if err != nil {
			return err
		}
		publicKeySuit := privateKeySuit.PublicKey()

		prikey := privateKeySuit.String()
		pubkey := publicKeySuit.String()

		createdAt := time.Now().Unix()
		updatedAt := createdAt
		insertSql := fmt.Sprintf(`INSERT INTO %s (created_at,updated_at,username,prikey,pubkey) VALUES (?,?,?,?,?)`, dbs.UserTable)
		_, err = dbs.Sqlite.Exec(insertSql, createdAt, updatedAt, username, prikey, pubkey)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("username already exists")
}

func UserUpdate(userData *models.Users) error {
	userID := userData.CommonModel.ID
	uid, err := UserIDCheck(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	username := userData.Username
	updatedAt := time.Now().Unix()
	updateSql := fmt.Sprintf(`UPDATE %s SET updated_at=?,username=? WHERE id=?`, dbs.UserTable)
	_, err = dbs.Sqlite.Exec(updateSql, updatedAt, username, uid)
	if err != nil {
		return err
	}
	return nil
}

func UserDelete(userID int) error {
	uid, err := UserIDCheck(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	deleteSql := fmt.Sprintf(`UPDATE %s SET state=0 WHERE id=?`, dbs.UserTable)
	_, err = dbs.Sqlite.Exec(deleteSql, uid)
	if err != nil {
		return err
	}
	return nil
}

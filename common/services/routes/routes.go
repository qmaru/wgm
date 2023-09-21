package routes

import (
	"fmt"
	"time"

	"wgm/common/dbs"
	"wgm/common/dbs/models"
	"wgm/common/services/common"
)

func RouteIDCheck(routeID int) (int, error) {
	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE id=? AND state=1`, dbs.RouteTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, routeID)
	if err != nil {
		return 0, err
	}

	var rid int
	err = row.Scan(&rid)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

func RouteList() ([]map[string]any, error) {
	selectSql := fmt.Sprintf(`SELECT id,cidr FROM %s WHERE state=1`, dbs.RouteTable)
	rows, err := dbs.Sqlite.Query(selectSql)
	if err != nil {
		return nil, err
	}

	var rId int
	var rCidr string

	data := make([]map[string]any, 0)
	for rows.Next() {
		err = rows.Scan(&rId, &rCidr)
		if err != nil {
			return nil, err
		}
		tmp := make(map[string]any)
		tmp["id"] = rId
		tmp["cidr"] = rCidr
		data = append(data, tmp)
	}
	return data, nil
}

func RouteAdd(routeData *models.Routes) error {
	cidr := routeData.CIDR
	cidrN, err := common.CIDRValidator(cidr)
	if err != nil {
		return err
	}

	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE cidr=? AND state=1`, dbs.RouteTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, cidrN)
	if err != nil {
		return err
	}

	var rid int
	err = row.Scan(&rid)
	if err != nil {
		createdAt := time.Now().Unix()
		updatedAt := createdAt
		selectSql := fmt.Sprintf(`INSERT INTO %s (created_at,updated_at,cidr) VALUES (?,?,?)`, dbs.RouteTable)
		_, err = dbs.Sqlite.Exec(selectSql, createdAt, updatedAt, cidrN)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("route already exists")
}

func RouteUpdate(routeData *models.Routes) error {
	routeID := routeData.CommonModel.ID
	rid, err := RouteIDCheck(routeID)
	if err != nil {
		return fmt.Errorf("route not found")
	}

	cidr := routeData.CIDR
	cidrN, err := common.CIDRValidator(cidr)
	if err != nil {
		return err
	}

	updatedAt := time.Now().Unix()
	updateSql := fmt.Sprintf(`UPDATE %s SET updated_at=?,cidr=? WHERE id=?`, dbs.RouteTable)
	_, err = dbs.Sqlite.Exec(updateSql, updatedAt, cidrN, rid)
	if err != nil {
		return err
	}
	return nil
}

func RouteDelete(routeID int) error {
	rid, err := RouteIDCheck(routeID)
	if err != nil {
		return fmt.Errorf("route not found")
	}

	deleteSql := fmt.Sprintf(`UPDATE %s SET state=0 WHERE id=?`, dbs.RouteTable)
	_, err = dbs.Sqlite.Exec(deleteSql, rid)
	if err != nil {
		return err
	}
	return nil
}

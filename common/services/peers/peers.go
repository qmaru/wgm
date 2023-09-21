package peers

import (
	"fmt"
	"time"

	"wgm/common/dbs"
	"wgm/common/dbs/models"
	"wgm/common/services/common"
	"wgm/common/services/users"
)

func PeerIDCheck(peerID int) (int, error) {
	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE id=? AND state=1`, dbs.PeerTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, peerID)
	if err != nil {
		return 0, err
	}

	var pid int
	err = row.Scan(&pid)
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func PeerList() ([]map[string]any, error) {
	selectSql := fmt.Sprintf(`
	SELECT
		u.username,
		u.prikey,
		u.pubkey,
		p.id,
		p.public_addr,
		p.private_addr,
		p.port,
		p.allowed_ips,
		p.mtu,
		p.dns,
		p.keepalive
	FROM %s as u
	LEFT JOIN %s as p ON u.id = p.user_id
	WHERE u.state=1 AND p.state=1`, dbs.UserTable, dbs.PeerTable)
	rows, err := dbs.Sqlite.Query(selectSql)
	if err != nil {
		return nil, err
	}

	var pUsername string
	var pPrikey string
	var pPubkey string
	var pId int
	var pPublicAddr string
	var pPrivateAddr string
	var pPort int
	var pAllowedIps string
	var pMtu int
	var pDns string
	var pKeepalive int

	data := make([]map[string]any, 0)
	for rows.Next() {
		err = rows.Scan(&pUsername, &pPrikey, &pPubkey, &pId, &pPublicAddr, &pPrivateAddr, &pPort, &pAllowedIps, &pMtu, &pDns, &pKeepalive)
		if err != nil {
			return nil, err
		}
		tmp := make(map[string]any)
		tmp["username"] = pUsername
		tmp["prikey"] = pPrikey
		tmp["pubkey"] = pPubkey
		tmp["id"] = pId
		tmp["public_addr"] = pPublicAddr
		tmp["private_addr"] = pPrivateAddr
		tmp["port"] = pPort
		tmp["allowed_ips"] = pAllowedIps
		tmp["mtu"] = pMtu
		tmp["dns"] = pDns
		tmp["keepalive"] = pKeepalive
		data = append(data, tmp)
	}
	return data, nil
}

func PeerAdd(peerData *models.Peers) error {
	userID := peerData.UserID
	pubAddr := peerData.PublicAddr
	priAddr := peerData.PrivateAddr
	port := peerData.Port

	uid, err := users.UserIDCheck(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if pubAddr != "" {
		_, err := common.PublicValidator(pubAddr)
		if err != nil {
			return err
		}
	}

	_, err = common.PrivateValidator(priAddr)
	if err != nil {
		return err
	}

	selectSql := fmt.Sprintf(`SELECT id FROM %s WHERE user_id=? AND public_addr=? AND private_addr=? AND port=? AND state=1`, dbs.PeerTable)
	row, err := dbs.Sqlite.QueryOne(selectSql, uid, pubAddr, priAddr, port)
	if err != nil {
		return err
	}
	var pid int
	err = row.Scan(&pid)
	if err != nil {
		allowedIPs := peerData.AllowedIPs
		mtu := peerData.MTU
		dns := peerData.DNS
		keepalive := peerData.PersistentKeepalive
		if keepalive == 0 {
			keepalive = 25
		}
		if allowedIPs == "" {
			allowedIPs = priAddr + "/32"
		}
		createdAt := time.Now().Unix()
		updatedAt := createdAt
		insertSql := fmt.Sprintf(`INSERT INTO %s (created_at,updated_at,user_id,public_addr,private_addr,port,allowed_ips,mtu,dns,keepalive) VALUES (?,?,?,?,?,?,?,?,?,?)`, dbs.PeerTable)
		_, err = dbs.Sqlite.Exec(insertSql, createdAt, updatedAt, uid, pubAddr, priAddr, port, allowedIPs, mtu, dns, keepalive)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("peer already exists")
}

func PeerUpdate(peerData *models.Peers) error {
	peerID := peerData.CommonModel.ID
	pid, err := PeerIDCheck(peerID)
	if err != nil {
		return fmt.Errorf("peer not found")
	}

	pubAddr := peerData.PublicAddr
	priAddr := peerData.PrivateAddr

	if pubAddr != "" {
		_, err := common.PublicValidator(pubAddr)
		if err != nil {
			return err
		}
	}

	_, err = common.PrivateValidator(priAddr)
	if err != nil {
		return err
	}

	port := peerData.Port
	allowedIPs := peerData.AllowedIPs
	mtu := peerData.MTU
	dns := peerData.DNS
	keepalive := peerData.PersistentKeepalive

	updatedAt := time.Now().Unix()
	updateSql := fmt.Sprintf(`UPDATE %s SET updated_at=?,public_addr=?,private_addr=?,port=?,allowed_ips=?,mtu=?,dns=?,keepalive=? WHERE id=?`, dbs.PeerTable)
	_, err = dbs.Sqlite.Exec(updateSql, updatedAt, pubAddr, priAddr, port, allowedIPs, mtu, dns, keepalive, pid)
	if err != nil {
		return err
	}
	return nil
}

func PeerDelete(peerID int) error {
	pid, err := PeerIDCheck(peerID)
	if err != nil {
		return fmt.Errorf("peer not found")
	}

	deleteSql := fmt.Sprintf(`UPDATE %s SET state=0 WHERE id=?`, dbs.PeerTable)
	_, err = dbs.Sqlite.Exec(deleteSql, pid)
	if err != nil {
		return err
	}
	return nil
}

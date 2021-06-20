package services

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"wgm/models"

	"github.com/liushuochen/gotable"
)

func Output(otype string, data []map[string]string) {
	var headers []string

	switch otype {
	case "userList":
		headers = []string{"ID", "Username", "ServerID", "Server", "IP", "Private", "Public"}
	case "serverList":
		headers = []string{"ID", "Server", "Endpoint", "Lan"}
	case "ruleList":
		headers = []string{"ID", "AllowedIP"}
	}

	tb, _ := gotable.CreateTable(headers)
	for _, d := range data {
		tb.AddValue(d)
	}
	tb.PrintTable()
}

func ShowUserList() {
	userQuery := `
		SELECT
			servers.id,
			servers.title,
			users.username,
			users.id,
			users.ip,
			users.prikey,
			users.pubkey
		FROM
			users
		LEFT JOIN servers on users.server_id = servers.id
		WHERE servers.status = 1 and users.status = 1
	`

	rows, err := models.DBQuery(userQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sID string
	var sTitle string
	var uUsername string
	var uID string
	var uIP string
	var uPrikey string
	var uPubKey string
	userData := make([]map[string]string, 0)

	for rows.Next() {
		rows.Scan(&sID, &sTitle, &uUsername, &uID, &uIP, &uPrikey, &uPubKey)
		uData := make(map[string]string)
		uData["ID"] = uID
		uData["Username"] = uUsername
		uData["ServerID"] = sID
		uData["Server"] = sTitle
		uData["IP"] = uIP
		uData["Private"] = uPrikey
		uData["Public"] = uPubKey
		userData = append(userData, uData)
	}
	Output("userList", userData)
}

func ShowServerList() {
	serverQuery := `
		SELECT
			id,
			title,
			address,
			port,
			lan_ip,
			lan_netmask
		FROM
			servers
		WHERE status = 1
	`
	rows, err := models.DBQuery(serverQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sid int
	var stitle string
	var saddress string
	var sport string
	var slanIP string
	var slanNetmask string
	serverData := make([]map[string]string, 0)

	for rows.Next() {
		rows.Scan(&sid, &stitle, &saddress, &sport, &slanIP, &slanNetmask)
		sData := make(map[string]string)
		sData["ID"] = strconv.Itoa(sid)
		sData["Server"] = stitle
		sData["Endpoint"] = fmt.Sprintf("%s:%s", saddress, sport)
		sData["Lan"] = fmt.Sprintf("%s/%s", slanIP, slanNetmask)
		serverData = append(serverData, sData)
	}
	Output("serverList", serverData)
}

func ShowRuleList() {
	ruleQuery := `
		SELECT
			id,
			allowedip
		FROM
			rules
		WHERE status = 1
	`
	rows, err := models.DBQuery(ruleQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rid int
	var rAllowed string
	ruleData := make([]map[string]string, 0)

	for rows.Next() {
		rows.Scan(&rid, &rAllowed)
		rData := make(map[string]string)
		rData["ID"] = strconv.Itoa(rid)
		rData["AllowedIP"] = rAllowed
		ruleData = append(ruleData, rData)
	}
	Output("ruleList", ruleData)
}

func ServerCfg(server int, username string, extraRule bool) (string, string) {
	interfaceQuery := `
		SELECT
			users.prikey,
			users.ip,
			servers.title,
			servers.port,
			servers.mtu,
			servers.dns
		FROM
			users
		LEFT JOIN servers on users.server_id = servers.id
		WHERE servers.status = 1 and users.status = 1 and servers.id = ? and users.username = ?
	`

	interfaceRow := models.DBQueryOne(interfaceQuery, server, username)

	var iPrikey string
	var iIP string
	var iTitle string
	var iPort string
	var iMTU int
	var iDNS string

	interfaceRow.Scan(&iPrikey, &iIP, &iTitle, &iPort, &iMTU, &iDNS)
	cfgInterface := "[Interface]\n"
	cfgPrikey := fmt.Sprintf("PrivateKey = %s\n", iPrikey)
	cfgAddress := fmt.Sprintf("Address = %s\n", iIP)
	cfgListenPort := fmt.Sprintf("ListenPort = %s\n", iPort)
	cfgMTU := ""
	cfgDNS := ""
	if iMTU != 0 {
		cfgMTU = fmt.Sprintf("MTU = %d\n", iMTU)
	}
	if cfgDNS != "" {
		cfgDNS = fmt.Sprintf("ListenPort = %s\n", iDNS)
	}

	peerQuery := `
	SELECT
		users.username,
		users.pubkey,
		users.default_rule
	FROM
		users
	LEFT JOIN servers on users.server_id = servers.id
	WHERE users.status = 1 and servers.status = 1 and servers.id = ? and users.username != ?
	`

	peerRows, _ := models.DBQuery(peerQuery, server, username)
	defer peerRows.Close()
	var pusername string
	var ppubkey string
	var pdrule string

	cfgPeers := ""
	for peerRows.Next() {
		peerRows.Scan(&pusername, &ppubkey, &pdrule)

		var rAllowedIPs []string
		rAllowedIPs = append(rAllowedIPs, pdrule)

		if extraRule {
			ruleQuery := `
			SELECT
				rules.allowedip as extra_rule
			FROM
				users
			LEFT JOIN rulemap on users.id = rulemap.user_id
			LEFT JOIN rules on rules.id = rulemap.rule_id
			WHERE users.status = 1 and rulemap.status = 1 and rules.status = 1 and users.username=? and users.is_extra=1 and users.server_id = ?
		`
			ruleRows, _ := models.DBQuery(ruleQuery, pusername, server)
			defer ruleRows.Close()

			var rAllowedIP string

			for ruleRows.Next() {
				ruleRows.Scan(&rAllowedIP)
				rAllowedIPs = append(rAllowedIPs, rAllowedIP)
			}
		}

		rule := strings.Join(rAllowedIPs, ",")

		cfgPeer := "[Peer]\n"
		cfgUser := fmt.Sprintf("#%s\n", pusername)
		cfgPubkey := fmt.Sprintf("PublicKey = %s\n", ppubkey)
		cfgAllowed := fmt.Sprintf("AllowedIPs = %s\n", rule)
		cfgPeers = cfgPeers + cfgPeer + cfgUser + cfgPubkey + cfgAllowed + "\n"
	}

	cfg := cfgInterface + cfgPrikey + cfgAddress + cfgListenPort + cfgMTU + cfgDNS + "\n" + cfgPeers
	return cfg, iTitle
}

func ClientCfg(server int, username, mainnode string, extraRule bool) (string, string) {
	interfaceQuery := `
		SELECT
			users.prikey,
			users.ip,
			servers.title,
			servers.mtu,
			servers.dns
		FROM
			users
		LEFT JOIN servers on users.server_id = servers.id
		WHERE servers.status = 1 and users.status = 1 and servers.id = ? and users.username = ?
	`

	interfaceRow := models.DBQueryOne(interfaceQuery, server, username)

	var iPrikey string
	var iIP string
	var iTitle string
	var iMTU int
	var iDNS string

	interfaceRow.Scan(&iPrikey, &iIP, &iTitle, &iMTU, &iDNS)
	cfgInterface := "[Interface]\n"
	cfgPrikey := fmt.Sprintf("PrivateKey = %s\n", iPrikey)
	cfgAddress := fmt.Sprintf("Address = %s\n", iIP)
	cfgMTU := ""
	cfgDNS := ""
	if iMTU != 0 {
		cfgMTU = fmt.Sprintf("MTU = %d\n", iMTU)
	}
	if cfgDNS != "" {
		cfgDNS = fmt.Sprintf("ListenPort = %s\n", iDNS)
	}

	peerQuery := `
	SELECT
		servers.address,
		servers.port,
		users.username,
		users.pubkey,
		users.default_rule,
		users.keepalive
	FROM
		users
	LEFT JOIN servers on users.server_id = servers.id
	WHERE users.status = 1 and servers.status = 1 and servers.id = ? and users.username = ?
	`

	peerRows, _ := models.DBQuery(peerQuery, server, mainnode)
	defer peerRows.Close()

	var paddress string
	var pport string
	var pusername string
	var ppubkey string
	var pdrule string
	var pkeepalive int

	cfgPeers := ""
	for peerRows.Next() {
		peerRows.Scan(&paddress, &pport, &pusername, &ppubkey, &pdrule, &pkeepalive)

		var rAllowedIPs []string
		rAllowedIPs = append(rAllowedIPs, pdrule)

		if extraRule {
			ruleQuery := `
			SELECT
				rules.allowedip
			FROM
				users
			LEFT JOIN rulemap on users.id = rulemap.user_id
			LEFT JOIN rules on rules.id = rulemap.rule_id
			WHERE users.status = 1 and rulemap.status = 1 and rules.status = 1 and users.username=? and users.server_id = ?
		`

			ruleRows, _ := models.DBQuery(ruleQuery, username, server)
			defer ruleRows.Close()

			var rAllowedIP string

			for ruleRows.Next() {
				ruleRows.Scan(&rAllowedIP)
				rAllowedIPs = append(rAllowedIPs, rAllowedIP)
			}
		}

		rule := strings.Join(rAllowedIPs, ",")

		cfgPeer := "[Peer]\n"
		cfgPubkey := fmt.Sprintf("PublicKey = %s\n", ppubkey)
		cfgAllowed := fmt.Sprintf("AllowedIPs = %s\n", rule)
		cfgEndpoint := fmt.Sprintf("Endpoint = %s:%s\n", paddress, pport)
		cfgKeepalive := fmt.Sprintf("PersistentKeepalive = %d\n", pkeepalive)
		cfgPeers = cfgPeers + cfgPeer + cfgPubkey + cfgAllowed + cfgEndpoint + cfgKeepalive + "\n"
	}

	cfg := cfgInterface + cfgPrikey + cfgAddress + cfgMTU + cfgDNS + "\n" + cfgPeers
	return cfg, iTitle
}

func ShowUserCfg(utype int, server int, username, mainnode string, extraRule bool, cfgname string) string {
	if utype == 1 {
		cfg, serverTitle := ServerCfg(server, username, extraRule)
		if cfgname != "" {
			MakeConfig(serverTitle, cfgname, cfg)
		}
		return cfg
	} else if utype == 2 {
		cfg, serverTitle := ClientCfg(server, username, mainnode, extraRule)
		if cfgname != "" {
			MakeConfig(serverTitle, username, cfg)
		}
		return cfg
	}
	return ""
}

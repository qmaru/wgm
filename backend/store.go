package backend

import (
	"github.com/qmaru/wgm/common/dbs/models"
	"github.com/qmaru/wgm/common/services/peers"
	"github.com/qmaru/wgm/common/services/routes"
	"github.com/qmaru/wgm/common/services/users"
)

// User, Route and Peer are data rows the UI displays.
type User struct {
	ID         int
	Username   string
	PrivateKey string
	PublicKey  string
}

type Route struct {
	ID   int
	CIDR string
}

type Peer struct {
	ID          int
	Username    string
	Prikey      string
	Pubkey      string
	PrivateAddr string
	PublicAddr  string
	Port        int
	AllowedIPs  string
	MTU         int
	DNS         string
	Keepalive   int
}

// PeerInput contains editable fields from the add/update peer form.
type PeerInput struct {
	UserID      int
	PrivateAddr string
	PublicAddr  string
	Port        int
	AllowedIPs  string
	MTU         int
	DNS         string
	Keepalive   int
}

// Store defines the persistence operations used by the UI.
type Store interface {
	Users() ([]User, error)
	AddUser(username string) error
	UpdateUser(id int, username string) error
	DeleteUser(id int) error

	Peers() ([]Peer, error)
	AddPeer(in PeerInput) error
	UpdatePeer(id int, in PeerInput) error
	DeletePeer(id int) error

	Routes() ([]Route, error)
	AddRoute(cidr string) error
	UpdateRoute(id int, cidr string) error
	DeleteRoute(id int) error
}

// DBStore adapts the existing service layer to the UI's store interface.
type DBStore struct{}

func (DBStore) Users() ([]User, error) {
	rows, err := users.UserList()
	if err != nil {
		return nil, err
	}
	items := make([]User, 0, len(rows))
	for _, row := range rows {
		items = append(items, User{
			ID:         row["id"].(int),
			Username:   row["username"].(string),
			PrivateKey: row["private_key"].(string),
			PublicKey:  row["public_key"].(string),
		})
	}
	return items, nil
}

func (DBStore) AddUser(username string) error {
	return users.UserAdd(&models.Users{Username: username})
}

func (DBStore) UpdateUser(id int, username string) error {
	return users.UserUpdate(&models.Users{
		CommonModel: models.CommonModel{ID: id},
		Username:    username,
	})
}

func (DBStore) DeleteUser(id int) error {
	return users.UserDelete(id)
}

func (DBStore) Peers() ([]Peer, error) {
	rows, err := peers.PeerList()
	if err != nil {
		return nil, err
	}
	items := make([]Peer, 0, len(rows))
	for _, row := range rows {
		items = append(items, Peer{
			ID:          row["id"].(int),
			Username:    row["username"].(string),
			Prikey:      row["prikey"].(string),
			Pubkey:      row["pubkey"].(string),
			PrivateAddr: row["private_addr"].(string),
			PublicAddr:  row["public_addr"].(string),
			Port:        row["port"].(int),
			AllowedIPs:  row["allowed_ips"].(string),
			MTU:         row["mtu"].(int),
			DNS:         row["dns"].(string),
			Keepalive:   row["keepalive"].(int),
		})
	}
	return items, nil
}

func (DBStore) AddPeer(in PeerInput) error {
	return peers.PeerAdd(&models.Peers{
		UserID:              in.UserID,
		PublicAddr:          in.PublicAddr,
		PrivateAddr:         in.PrivateAddr,
		Port:                in.Port,
		AllowedIPs:          in.AllowedIPs,
		MTU:                 in.MTU,
		DNS:                 in.DNS,
		PersistentKeepalive: in.Keepalive,
	})
}

func (DBStore) UpdatePeer(id int, in PeerInput) error {
	return peers.PeerUpdate(&models.Peers{
		CommonModel:         models.CommonModel{ID: id},
		PublicAddr:          in.PublicAddr,
		PrivateAddr:         in.PrivateAddr,
		Port:                in.Port,
		AllowedIPs:          in.AllowedIPs,
		MTU:                 in.MTU,
		DNS:                 in.DNS,
		PersistentKeepalive: in.Keepalive,
	})
}

func (DBStore) DeletePeer(id int) error {
	return peers.PeerDelete(id)
}

func (DBStore) Routes() ([]Route, error) {
	rows, err := routes.RouteList()
	if err != nil {
		return nil, err
	}
	items := make([]Route, 0, len(rows))
	for _, row := range rows {
		items = append(items, Route{
			ID:   row["id"].(int),
			CIDR: row["cidr"].(string),
		})
	}
	return items, nil
}

func (DBStore) AddRoute(cidr string) error {
	return routes.RouteAdd(&models.Routes{CIDR: cidr})
}

func (DBStore) UpdateRoute(id int, cidr string) error {
	return routes.RouteUpdate(&models.Routes{
		CommonModel: models.CommonModel{ID: id},
		CIDR:        cidr,
	})
}

func (DBStore) DeleteRoute(id int) error {
	return routes.RouteDelete(id)
}

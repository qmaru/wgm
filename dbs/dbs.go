package dbs

import (
	"log"

	"wgm/configs"
	"wgm/dbs/models"

	"github.com/qmaru/qdb/sqlitep"
)

const (
	UserTable  string = "users"
	RouteTable string = "routes"
	PeerTable  string = "peers"
)

var Sqlite *sqlitep.Sqlitep

func init() {
	db, err := configs.DatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}
	Sqlite = sqlitep.New(db)
	CreataTable()
}

// CreataTable 初始化数据表
func CreataTable() error {
	tables := []any{
		models.Users{},
		models.Routes{},
		models.Peers{},
	}
	return Sqlite.CreateTable(tables)
}

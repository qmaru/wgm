package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"wgm/config"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string

func init() {
	dbPath = config.SetDBPath("wgm.db")
}

func TableCheck() bool {
	_, err := os.Stat(dbPath)
	return err == nil
}

func connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("SQLite3 Connect Failed")
	}
	return db
}

// DBExec 原生执行
func DBExec(raw string, args ...interface{}) {
	db := connect()
	defer db.Close()
	stmt, err := db.Prepare(raw)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
}

// DBQuery 查询数据
func DBQuery(raw string, args ...interface{}) (rows *sql.Rows, err error) {
	db := connect()
	defer db.Close()
	stmt, err := db.Prepare(raw)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err = stmt.Query(args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, nil
}

// DBQueryOne 查询单条数据
func DBQueryOne(raw string, args ...interface{}) (row *sql.Row) {
	db := connect()
	defer db.Close()
	stmt, err := db.Prepare(raw)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()
	return stmt.QueryRow(args...)
}

// InitTable 初始化数据表
func InitTable() {
	tabls := []interface{}{
		Users{},
		Servers{},
		Rules{},
		RuleMap{},
	}
	for _, table := range tabls {
		var buffer bytes.Buffer
		rType := reflect.TypeOf(table)
		rName := strings.ToLower(rType.Name())
		DBFiled(rType, &buffer)
		rFiled := buffer.Bytes()[0 : len(buffer.Bytes())-1]

		sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", rName, rFiled)
		DBExec(sql)
	}
}

package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MapConnection struct {
	db *sql.DB
}

var m map[string]MapConnection

func Connect(dataSourceName string) (*sql.DB, error) {
	var err error

	if m == nil {
		m = make(map[string]MapConnection)
	}
	if m[dataSourceName].db != nil {
		fmt.Println("connect-reuse:" + dataSourceName)

		db := m[dataSourceName].db
		if err = db.Ping(); err != nil {
			log.Panic(err)
			return nil, err
		}
		return db, nil
	}
	fmt.Println("connect-new:" + dataSourceName)
	params := strings.Split(dataSourceName, "|")
	ds := params[0]
	dr := params[1]
	maxOpen, _ := strconv.Atoi(params[2])
	maxIdle, _ := strconv.Atoi(params[3])
	//maxLifetime, _ := strconv.ParseFloat(params[4], 64)
	//fmt.Println(dr, ds, maxOpen, maxIdle)

	dbx, err := sql.Open(dr, ds)
	dbx.SetConnMaxLifetime(time.Second * 60 * 12)
	dbx.SetMaxIdleConns(maxIdle)
	dbx.SetMaxOpenConns(maxOpen)

	if err != nil {
		log.Panic(err)
		return nil, err
	}
	if err = dbx.Ping(); err != nil {
		log.Panic(err)
		return nil, err
	}
	m[dataSourceName] = MapConnection{dbx}

	return dbx, nil
}

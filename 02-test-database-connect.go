package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func main() {
	var err error
	Db, err = sql.Open("mysql", "root:root@(127.0.0.1)/love_ple")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接数据库成功")
}

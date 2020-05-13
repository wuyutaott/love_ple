package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

type Member struct {
	Id 		int
	CardId	int
	Phone	string
}

func main() {
	var err error
	Db, err = sql.Open("mysql", "root:root@(127.0.0.1)/love_ple")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接数据库成功")

	// 查询数据


	defer Db.Close()
}

func Members() (members []Member, err error) {
	rows, err := Db.Query("SELECT id, card_id, phone FROM members ")
}

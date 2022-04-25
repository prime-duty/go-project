package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	name string
	pwd  string
}

func main() {
	dsn := "duty:123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}
	rows, _ := db.Query("select * from test")
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.name, &user.pwd)
		fmt.Println(user)
	}
}

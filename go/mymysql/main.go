package main

// Imports.
import (
	"database/sql"
	"fmt"
	// "github.com/ziutek/mymysql/mysql"
	// _ "github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

/*
func main() {
	db := mysql.New("tcp", "", "127.0.0.1:3306", "dev", "dev", "devel")
	err := db.Connect()
	if err != nil {
		fmt.Println("Open;", err)
	}

	sql := "DROP SCHEMA IF EXISTS foo; CREATE SCHEMA IF NOT EXISTS foo;"
	_, _, err = db.Query(sql)
	if err != nil {
		fmt.Println(sql, err)
	}

	sql = "USE DATABASE foo;"
	_, _, err = db.Query(sql)
	if err != nil {
		fmt.Println(sql, err)
	}

	db.Close()
}
*/
import _ "github.com/ziutek/mymysql/godrv"

func main() {
	db, err := sql.Open("mymysql", "devel/dev/dev")
	if err != nil {
		fmt.Println("Open;", err)
	}

	sql := "DROP SCHEMA IF EXISTS foo; CREATE SCHEMA IF NOT EXISTS foo;"
	db.Start(sql)
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(sql, err)
	}

	sql = "USE DATABASE foo;"
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(sql, err)
	}

	db.Close()
}

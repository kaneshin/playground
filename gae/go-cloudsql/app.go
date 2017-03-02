package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const dbName = "sample_db"

var db *sql.DB

func init() {
	user := os.Getenv("CLOUDSQL_USER")
	password := os.Getenv("CLOUDSQL_PASSWORD")
	connectionName := os.Getenv("CLOUDSQL_CONNECTION_NAME")

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err != nil {
		panic(err)
	}
}

func init() {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		panic(err)
	}
}

func init() {
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.user (%s, %s, %s)",
		dbName,
		`id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY`,
		`first_name VARCHAR(30) NOT NULL`,
		`last_name VARCHAR(30) NOT NULL`,
	))
	if err != nil {
		panic(err)
	}
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		switch r.Method {
		case http.MethodGet:
			rows, err := db.Query(fmt.Sprintf("SELECT first_name, last_name FROM %s.user", dbName))
			if err != nil {
				log.Errorf(ctx, "db.Query: %v", err)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var firstName, lastName string
				if err := rows.Scan(&firstName, &lastName); err != nil {
					log.Errorf(ctx, "rows.Scan: %v", err)
					continue
				}
				log.Infof(ctx, "First: %v - Last: %v", firstName, lastName)
			}
			if err := rows.Err(); err != nil {
				log.Errorf(ctx, "Row error: %v", err)
			}

		case http.MethodPost:
			r.ParseForm()
			firstName := r.PostForm.Get("first_name")
			lastName := r.PostForm.Get("last_name")

			_, err := db.Exec(fmt.Sprintf("INSERT INTO %s.user (first_name, last_name) VALUES (?, ?)", dbName), firstName, lastName)
			if err != nil {
				log.Errorf(ctx, "db.Query: %v", err)
				return
			}

		}
	})
}

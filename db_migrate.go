package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	stmt, err := db.Prepare("CREATE TABLE user (id int NOT NULL AUTO_INCREMENT, sso_id varchar(255) NOT NULL, password varchar(255), first_name varchar(255), last_name varchar(255),  email varchar(255) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, deleted_at TIMESTAMP ,PRIMARY KEY (id), UNIQUE (sso_id, email));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Person Table successfully migrated....")
	}
}

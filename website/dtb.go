package website

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func StartDb() {
	var err error
	db, err = sql.Open("mysql", "Alvarez:89173036695Druno@tcp(127.0.0.1:3306)/")

	err = db.Ping()
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS scoreboard")
	if err != nil {
		log.Print("Wait a second")
	}

	db, err = sql.Open("mysql", "Alvarez:89173036695Druno@tcp(127.0.0.1:3306)/scoreboard")
	err = db.Ping()
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("USE scoreboard")
	if err != nil {
		log.Print("Wait a second")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users " +
		"(id INT AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(50)," +
		"password VARCHAR(50))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories " +
		"(id INT AUTO_INCREMENT PRIMARY KEY," +
		"category_name VARCHAR(50)," +
		"count_of_tasks VARCHAR(50))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rating " +
		"(id INT AUTO_INCREMENT PRIMARY KEY," +
		"user_id INT," +
		"category_id INT," +
		"score INT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user_role " +
		"(id INT AUTO_INCREMENT PRIMARY KEY," +
		"userId INT," +
		"role INT)")
	if err != nil {
		log.Fatal(err)
	}
}

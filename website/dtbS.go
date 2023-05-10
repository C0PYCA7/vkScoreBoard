package website

import (
	"database/sql"
	"log"
)

var Count int

func FindUser(login string, password string) string {
	query := "SELECT username FROM users WHERE username = ? AND password = ?"
	row := db.QueryRow(query, login, password)
	var UserNameValue string
	err := row.Scan(&UserNameValue)
	if err != nil {
		log.Println("Error scanning")
		return ""
	}
	return UserNameValue
}

func GetRatingTotal(db *sql.DB, userName string) []TaskRating {
	rows, err := db.Query("SELECT category_name, score FROM rating WHERE user_name = ?", userName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ratings []TaskRating
	for rows.Next() {
		var category string
		var rating int
		err := rows.Scan(&category, &rating)
		if err != nil {
			log.Fatal(err)
		}

		r := TaskRating{
			Category: category,
			Rating:   rating,
		}
		ratings = append(ratings, r)
	}
	if err != nil {
		log.Fatal(err)
	}
	return ratings
}

func GetRatingByEmployee(db *sql.DB, category_name string, score int) []TaskRating {

	rows, err := db.Query("SELECT category_name, COUNT(*) FROM rating "+
		"WHERE category_name = ? AND score > ? GROUP BY category_name", category_name, score)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ratings []TaskRating
	for rows.Next() {
		var category string
		var count int
		err := rows.Scan(&category, &count)
		if err != nil {
			log.Fatal(err)
		}

		r := TaskRating{
			Category: category,
			Rating:   count + 1,
		}
		ratings = append(ratings, r)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ratings
}

func Check() {
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&Count)
	if err != nil {
		log.Fatal(err)
	}

	if Count == 0 {

		query := "INSERT INTO categories (category_name, count_of_tasks) " +
			"VALUES (?, ?)"
		cname1 := "Кибербезопасность"
		countOf := 20

		_, err := db.Exec(query, cname1, countOf)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO categories (category_name, count_of_tasks) " +
			"VALUES (?, ?)"
		cname2 := "Скрытые угрозы"
		countOf = 30

		_, err = db.Exec(query, cname2, countOf)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO categories (category_name, count_of_tasks) " +
			"VALUES (?, ?)"
		cname3 := "Информационная безопасность"
		countOf = 25

		_, err = db.Exec(query, cname3, countOf)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO users (username, password) " +
			"VALUES (?, ?)"
		uname := "Billy"
		upass := "Herrington"
		urole := "Admin"
		uscore1 := 20
		uscore2 := 30
		uscore3 := 25
		_, err = db.Exec(query, uname, upass)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO user_role (user_id, role) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), ?)"
		_, err = db.Exec(query, uname, upass, urole)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname3, uscore3)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO users (username, password) " +
			"VALUES (?, ?)"
		uname = "Tyler"
		upass = "Derden"
		urole = "User"
		uscore1 = 15
		uscore2 = 23
		uscore3 = 21
		_, err = db.Exec(query, uname, upass)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO user_role (user_id, role) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), ?)"
		_, err = db.Exec(query, uname, upass, urole)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname3, uscore3)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO users (username, password) " +
			"VALUES (?, ?)"
		uname = "Mickey"
		upass = "Pirson"
		urole = "User"
		uscore1 = 12
		uscore2 = 25
		uscore3 = 20
		_, err = db.Exec(query, uname, upass)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO user_role (user_id, role) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), ?)"
		_, err = db.Exec(query, uname, upass, urole)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_name, category_name, score) " +
			"VALUES (?,?,?)"
		_, err = db.Exec(query, uname, cname3, uscore3)
		if err != nil {
			log.Fatal(err)
		}
	}
}

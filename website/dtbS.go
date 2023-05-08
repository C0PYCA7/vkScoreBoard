package website

import "log"

func CheckUser() {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {

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

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname3, uscore3)
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

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname3, uscore3)
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

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname1, uscore1)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname2, uscore2)
		if err != nil {
			log.Fatal(err)
		}

		query = "INSERT INTO rating (user_id, category_id, score) " +
			"VALUES ((SELECT id FROM users WHERE username = ? AND password = ?), " +
			"(SELECT id FROM categories WHERE category_name = ?), ?)"
		_, err = db.Exec(query, uname, upass, cname3, uscore3)
		if err != nil {
			log.Fatal(err)
		}
	}
}

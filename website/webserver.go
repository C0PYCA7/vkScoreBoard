package website

import (
	"html/template"
	"log"
	"net/http"
)

type TaskRating struct {
	Category string
	Rating   int
}

type ViewData struct {
	User             string
	RatingByEmployee []TaskRating
	RatingByTotal    []TaskRating
}

type ViewAllData struct {
	ViewAll []ViewData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data ViewAllData) {
	tmplPath := "templates/index.html"
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		userName := FindUser("Tyler", "Derden")

		ratingByTotalTasks := GetRatingTotal(db, userName)
		ratingByEmployee1 := GetRatingByEmployee(db, "Кибербезопасность", 15)
		data1 := ViewData{
			User:             userName,
			RatingByTotal:    ratingByTotalTasks,
			RatingByEmployee: ratingByEmployee1,
		}
		ratingByEmployee2 := GetRatingByEmployee(db, "Скрытые угрозы", 23)
		data2 := ViewData{
			RatingByEmployee: ratingByEmployee2,
		}
		ratingByEmployee3 := GetRatingByEmployee(db, "Информационная безопасность", 21)
		data3 := ViewData{
			RatingByEmployee: ratingByEmployee3,
		}
		viewAllD := ViewAllData{
			ViewAll: []ViewData{data1, data2, data3},
		}
		RenderTemplate(w, "index.html", viewAllD)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

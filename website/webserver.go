package website

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Message string
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := "templates/registration.html"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "GET" {
			login := r.FormValue("login")
			password := r.FormValue("password")
			FindUser(login, password)
			if Count > 0 {
				http.Redirect(w, r, "/main", http.StatusFound)
			} else {
				data := ViewData{
					Message: "Неверный логин или пароль",
				}
				RenderTemplate(w, "registration.html", data)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Description string
	Done        bool
}

type MainPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	t := template.Must(template.ParseFiles("templates/index.tmpl"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, MainPageData{PageTitle: "Todo App",
			Todos: []Todo{
				{Description: "Go to store"},
				{Description: "Clean windows"},
				{Description: "Pay taxes", Done: true},
			}})

		if err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println("🚀 App running on port http://localhost:80")
	http.ListenAndServe(":80", nil)
}

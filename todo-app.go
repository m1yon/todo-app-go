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

	todos := []Todo{
		{Description: "Go to store"},
		{Description: "Clean windows"},
		{Description: "Pay taxes", Done: true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := t.Execute(w, MainPageData{PageTitle: "Todo App", Todos: todos})

			if err != nil {
				fmt.Println(err)
			}

		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				fmt.Println(err)
			}

			newTodo := Todo{r.FormValue("description"), false}
			todos = append(todos, newTodo)

			err := t.Execute(w, MainPageData{PageTitle: "Todo App", Todos: todos})

			if err != nil {
				fmt.Println(err)
			}
		}
	})

	fmt.Println("🚀 App running on port http://localhost:80")
	http.ListenAndServe(":80", nil)
}

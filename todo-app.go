package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there!!")
	})

	fmt.Println("🚀 App running on port 80")
	http.ListenAndServe(":80", nil)
}

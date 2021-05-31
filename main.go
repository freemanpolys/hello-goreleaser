package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting web server at port 80")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World!</h1>")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

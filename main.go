package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

const file = "access.log"

func main() {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("log data:", string(data))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("Hello, %q <3", html.EscapeString(r.URL.Path))
		_, err := f.WriteString(response + "\n")
		if err != nil {
			fmt.Fprintf(w, "could not write response to file: %v", err)
			return
		}
		fmt.Fprint(w, response)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

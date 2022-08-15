package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/greetings" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hi, Welcome!")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "Request successful")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name => %s\n", name)
	fmt.Fprintf(w, "Surname => %s\n", surname)
	fmt.Fprintf(w, "Email => %s\n", email)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/greetings", handleGreeting)
	http.HandleFunc("/form", handleForm)

	fmt.Printf("server running at http://localhost:5001\n")
	log.Fatal(http.ListenAndServe(":5001", nil))
}

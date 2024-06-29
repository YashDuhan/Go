package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, r.ParseForm().Error())
		return
	}
	fmt.Fprintf(w, "POST requiest successfuk")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s", name)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Address :  %s", address)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	// Root Route signifying "/" (using index.html)
	http.Handle("/", fileserver)
	// To handle the form function (using form.html)
	http.HandleFunc("/form", formhandler)
	// To handle the hellohandler function
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

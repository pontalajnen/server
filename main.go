package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	fmt.Fprintln(w,
		`<!DOCTYPE html>
	<html>
	<form method="post" action="/name">
    <label for="firstname">First name:</label>
    <input type="text" name="firstname" /><br />
    <label for="lastname">Last name:</label>
    <input type="text" name="lastname" /><br />
    <input type="submit" />
	</form></html>`)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := r.Form.Get("firstname")
	lastname := r.Form.Get("lastname")
	fmt.Println(name, lastname)

	fmt.Fprintln(w, "Form submitted fine")
}

func main() {
	fmt.Println("Hello from BLHub server!")

	http.HandleFunc("/index.html", helloHandler)
	http.HandleFunc("/name", postHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

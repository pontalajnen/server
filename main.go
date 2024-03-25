package main

//Test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	// fmt.Fprintln(c.Writer,
	// 	`<!DOCTYPE html>
	// <html>
	// <form method="post" action="/name">
	// <label for="firstname">First name:</label>
	// <input type="text" name="firstname" /><br />
	// <label for="lastname">Last name:</label>
	// <input type="text" name="lastname" /><br />
	// <input type="submit" />
	// </form></html>`)

	// Av någon anledning behöver man skicka med något sorts objekt, ingen aning varför.
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "BLhub"})
}

func postHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := c.Request.Form.Get("firstname")
	lastname := c.Request.Form.Get("lastname")
	fmt.Println(name, lastname)

	fmt.Fprintln(c.Writer, "Form submitted fine")
}

func loginHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := c.Request.Form.Get("name")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	fmt.Println(name, email, password)

	fmt.Fprintln(c.Writer, "Form submitted fine")
}

func main() {
	fmt.Println("Hello from BLHub server!")

	engine := gin.Default()
	engine.LoadHTMLFiles("index.html")

	engine.GET("/", helloHandler)
	engine.POST("/name", postHandler)
	engine.POST("/login", loginHandler)

	// Run the server
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

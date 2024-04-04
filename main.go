package main

//Test

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	// fmt.Fprintln(c.Writer,
	// 	`<!DOCTYPE html>
	// <html>
	// <form method="post" action="/name">
	// <label for="firstName">First name:</label>
	// <input type="text" name="firstName" /><br />
	// <label for="lastName">Last name:</label>
	// <input type="text" name="lastName" /><br />
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

	name := c.Request.Form.Get("firstName")
	lastName := c.Request.Form.Get("lastName")
	fmt.Println(name, lastName)

	fmt.Fprintln(c.Writer, "Form submitted fine")
}

func postTest(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)

	//err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	// var user struct {
	// 	Name  string `json:"name"`
	// 	Email string `json:"email"`
	// }
	// if err := c.BindJSON(&user); err != nil {
	// 	fmt.Fprintln(c.Writer, "hej")
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// // TODO: create user in database
	//c.JSON(201, gin.H{"message": "User created"})
	fmt.Println("Wallahi")
	fmt.Println(jsonData)
	//fmt.Println(user)
	fmt.Fprintln(c.Writer, jsonData)
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
	engine.POST("/test", postTest)

	// Run the server
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

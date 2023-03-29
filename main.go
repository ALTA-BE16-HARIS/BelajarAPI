package main

import (
	"belajarAPI/books"
	"belajarAPI/users"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/login", users.Login())
	e.POST("/register", users.Register())
	e.POST("/books", books.Addbook())
	e.GET("/books", books.Getallbooks())
	e.Start(":3758")
}

package main

import (
	"book-rentals/config"
	"book-rentals/routes"
)

func main() {
	//initialize database
	config.InitDB()

	//run project
	e := routes.New()
	e.Start(":8080")
}

package main

import (
	"gin_example/src"
)

func main() {
	// Create a new Gin router with default middleware
	var router = src.MainRouter()

	// Run the server
	router.Run(":9000")
}

package main

import (
	"gin_example/configure"
	"gin_example/src"
)

func main() {
	configure.InitLogger()
	// Create a new Gin router with default middleware
	var router = src.MainRouter()

	// Run the server
	router.Run(":9000")
}

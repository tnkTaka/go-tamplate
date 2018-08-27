package main

import "github.com/go-example/router"

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}

package main

import (
	"assignment3/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}

package main

import (
	"log"

	"github.com/ocfBNj/hello/server"
)

func main() {
	log.Fatal(server.FileServer(":80", "C:/"))
}

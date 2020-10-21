package main

import (
	"log"

	"github.com/ocfbnj/hello/server"
)

func main() {
	log.Fatal(server.FileServer(":80", "C:/"))
}

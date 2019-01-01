package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muratsplat/mehmet/command"
)

const (
	host = "localhost"
	port = 8090
)

var (
	addr = fmt.Sprintf("%s:%d", host, port)
)

func main() {

	router := command.NewPHP()
	log.Fatal(http.ListenAndServe(addr, router))
}

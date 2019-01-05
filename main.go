package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/muratsplat/mehmet/command"
)

const (
	host = ""
	port = 8090
)

func init() {
	log.Printf("Server is starting.\n")
	flag.StringVar(&scriptPath, "path", "", "the path of script PHP file")
	flag.IntVar(&givenPort, "port", port, "port number for serving the script")
	flag.Parse()
	if scriptPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if givenPort == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

var (
	addr       = fmt.Sprintf("%s:%d", host, port)
	givenPort  int
	scriptPath string
)

func main() {
	router := command.NewPHP(scriptPath)
	log.Printf("%s listening...", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

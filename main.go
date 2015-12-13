package main

import (
	"fmt"
	"github.com/phayes/hookserve/hookserve"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	hookSecret = "circus$Alert=shiver"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("OPENSHIFT_GO_PORT"))
	if err != nil {
		log.Fatalln("Invalid port number")
	}

	server := hookserve.NewServer()

	server.Port = port
	server.Secret = hookSecret

	server.GoListenAndServe()

	for {
		select {
		case event := <-server.Events:
			fmt.Println(event.Owner + " " + event.Repo + " " + event.Branch + " " + event.Commit)
		default:
			time.Sleep(100)
		}
	}
}

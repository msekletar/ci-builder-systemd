package main

import (
	"fmt"
	"github.com/phayes/hookserve/hookserve"
	"os"
	"time"
)

func main() {
	server := hookserve.NewServer()
	server.Port = os.Getenv("OPENSHIFT_GO_PORT")
	server.Secret = "circus$Alert=shiver"
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

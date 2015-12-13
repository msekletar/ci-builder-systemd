package main

import (
	"fmt"
	"github.com/msekletar/hookserve/hookserve"
	"os"
	"time"
)

func main() {
	server := hookserve.NewServer()

	server.Port = os.Getenv("OPENSHIFT_GO_PORT")
	server.Address = os.Getenv("OPENSHIFT_GO_IP")

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

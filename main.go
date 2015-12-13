package main

import (
	"fmt"
	"github.com/msekletar/hookserve/hookserve"
	"os"
)

func main() {
	server := hookserve.NewServer()

	server.Port = os.Getenv("OPENSHIFT_GO_PORT")
	server.Address = os.Getenv("OPENSHIFT_GO_IP")

	server.GoListenAndServe()

	for event := range server.Events {
		fmt.Println(event.Owner + " " + event.Repo + " " + event.Branch + " " + event.Commit)
	}
}

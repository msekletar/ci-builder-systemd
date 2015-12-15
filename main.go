package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/msekletar/hookserve/hookserve"
)

func main() {
	server := hookserve.NewServer()

	server.Port = os.Getenv("OPENSHIFT_GO_PORT")
	server.Address = os.Getenv("OPENSHIFT_GO_IP")

	server.GoListenAndServe()

	workDir, err := createWorkdir()
	if err != nil {
		log.Fatalf("Failed to create workspace directory: %s\n", err)
	}

	defer func() {
		os.RemoveAll(workDir)
	}()

	for event := range server.Events {
		fmt.Println(event.Owner + " " + event.Repo + " " + event.Branch + " " + event.Commit)
	}
}

func createWorkdir() (string, error) {
	homeDir := os.Getenv("HOME")

	tempDir, err := ioutil.TempDir(homeDir, "workdir-")
	if err != nil {
		return "", err
	}

	return tempDir, nil

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

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

	// remove workdir if process terminates normally
	defer os.RemoveAll(workDir)

	// also remove it upon receiving SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		os.RemoveAll(workDir)
		os.Exit(0)
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

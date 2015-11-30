package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type apiHandler struct {
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to read request body: %s", err)
	}

	fmt.Printf("%s", body)
}

func setupHandlers() {
	http.Handle("/buildsrpm", apiHandler{})
}

func main() {
	serverAddress := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}

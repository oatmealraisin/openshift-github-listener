package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

// TODO:
var secretKey []byte

func ReceiveWebhook(w http.ResponseWriter, req *http.Request) {
	//	client := github.NewClient(nil)

	payload, err := github.ValidatePayload(req, secretKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch event := event.(type) {
	case *github.PushEvent:
		fmt.Println("asdf")
	default:
		fmt.Println(event)
	}
}

func main() {
	host := "0.0.0.0"
	port := "8080"

	if len(os.Getenv("PORT_LISTENER")) != 0 {
		port = os.Getenv("PORT_LISTENER")
		fmt.Printf("Using port %s...\n", port)
	}

	if len(os.Getenv("HOST_LISTENER")) != 0 {
		host = os.Getenv("HOST_LISTENER")
		fmt.Printf("Using host %s...\n", host)
	}

	http.HandleFunc("/", ReceiveWebhook)
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}()
}

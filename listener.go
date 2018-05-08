package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/google/go-github/github"
	"github.com/oatmealraisin/openshift-github-listener/pkg/oc"
)

// TODO:
var secretKey []byte
var UNITTEST_TAG = "[test]"
var INTEGRATION_TAG = "[integrate]"
var MERGE_TAG = "[merge]"

// TODO: Capture initial PR creation, start Acceptance test, assign reviewers
// TODO: Capture Test, Merge comments, start relevant tests
// TODO:
func ReceiveWebhook(w http.ResponseWriter, req *http.Request) {
	// client := github.NewClient(nil)
	secretKey = []byte("do the thing")

	payload, err := github.ValidatePayload(req, secretKey)
	if err != nil {
		fmt.Printf("Error when validating webhook: %s\n", err.Error())
	}

	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		fmt.Printf("Error when parsing webhook: %s\n", err.Error())
	}

	switch event := event.(type) {
	case *github.IssueCommentEvent:
		fmt.Println("TODO: Do something w/ a Issue Comment event")
		comment := event.GetComment().GetBody()
		fmt.Printf("Clone URL: %s\n", event.GetRepo().GetCloneURL())
		fmt.Printf("Git URL: %s\n", event.GetRepo().GetURL())
		fmt.Printf("Allow Merge Commit: %s\n", event.GetRepo().GetAllowMergeCommit())
		if strings.Contains(comment, MERGE_TAG) {
			oc.LaunchMergeTests(event.GetRepo())
			return
		}
		if strings.Contains(comment, UNITTEST_TAG) {
			oc.LaunchUnitTests(event.GetRepo())
		}
		if strings.Contains(comment, INTEGRATION_TAG) {
			oc.LaunchIntegrationTests(event.GetRepo())
		}
	case *github.PullRequestEvent:
		fmt.Println("TODO: Do something w/ a PR event")
		// Includes assigning a person
		// Includes tagging a PR
		// Includes adding a reviewer
	case *github.PullRequestReviewEvent:
		fmt.Println("TODO: Do something w/ a PR Review event")
	case *github.PullRequestReviewCommentEvent:
		fmt.Println("TODO: Do something w/ a PR Review Comment event")
	case *github.PushEvent:
		fmt.Println("TODO: Do something w/ a push event")
	case *github.PingEvent:
		fmt.Println("Received a ping event! Ready to serve :)")
	case *github.MemberEvent:
		fmt.Println("TODO: Do something w/ a Member Event")
		// Happens when someone accepts an invitation
	case *github.CreateEvent:
		fmt.Println("TODO: Do something w/ a create event")
	default:
		fmt.Println(reflect.TypeOf(event).String())
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

	// TODO: Initialize OpenShift, get finiteautomatom creds

	http.HandleFunc("/", ReceiveWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

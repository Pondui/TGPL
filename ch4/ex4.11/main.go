// Exercise 4.11: Build a tool that lets users create, read, update, and close GitHub issues from the
// command line, invoking their preferred text editor when substantial text input is required.
// issuecli get issues
// issuecli get issue issuenumber
// issuecli create issue myissue
// issuecli update issue myissue
// issuecli close issue myissue

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	// Command line args
	args := os.Args[1:]
	ownerPtr := flag.String("OWNER", "", "Owner of Github repository")
	repoPtr := flag.String("REPO", "", "Github repository to which issues belong")
	flag.Parse()

	// Authentication
	token := os.Getenv("GHPAT")
	client := &http.Client{}

	if slices.Contains(args, "get") && slices.Contains(args, "issues") && len(args) <= 4 {
		getAllIssues(*ownerPtr, *repoPtr, client, token);
		return
	} else if slices.Contains(args, "get") && slices.Contains(args, "issue") && len(args) <= 5 {
		getIssue(*ownerPtr, *repoPtr, args[len(args)-1], client, token)
	} else if slices.Contains(args, "close") && slices.Contains(args, "issue") && len(args) <= 5 {
		closeIssue(*ownerPtr, *repoPtr, args[len(args)-1], client, token)
	}
}

// Get all issues for a public repo
func getAllIssues(owner string, repo string, client *http.Client, token string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		addHeaders(req, token)
		printResponse(client, req)
	}
	return
}

// Get an issue by number for a given owner and repo
func getIssue(owner string, repo string, issueNumber string, client *http.Client, token string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues" + "/" + issueNumber

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		addHeaders(req, token)
		printResponse(client, req)
	}
	return
}

// Close an issue by number for a given owner and repo
func closeIssue(owner string, repo string, issueNumber string, client *http.Client, token string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues" + "/" + issueNumber

	reqBody := strings.NewReader("{\"state\": \"closed\"}")
	req, err := http.NewRequest("PATCH", url, reqBody)
	if err != nil {
		log.Fatal(err)
	} else {
		addHeaders(req, token)
		printResponse(client, req)
	}
	return
}

// Adds HTTP Headers to request
func addHeaders(req *http.Request, token string) {
	authHeader := "Bearer " + token
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Accept", "application/vnd.github+json")
}

func printResponse(client *http.Client, req *http.Request) {
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		b, err := io.ReadAll(resp.Body)

		var jsonResp bytes.Buffer
		json.Indent(&jsonResp, b, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(jsonResp.String())
	}
}
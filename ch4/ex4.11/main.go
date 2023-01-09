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
	"os/exec"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	// Command line args
	args := os.Args[1:]
	ownerPtr := flag.String("OWNER", "", "Owner of Github repository")
	repoPtr := flag.String("REPO", "", "Github repository to which issues belong")
	editorPtr := flag.String("EDITOR", "", "Which text editor to use when inputting longer JSON input")
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
	} else if slices.Contains(args, "create") && slices.Contains(args, "issue") && len(args) <= 6 {
		createIssue(*ownerPtr, *repoPtr, *editorPtr, args[len(args)-1], client, token)
	} else if slices.Contains(args, "update") && slices.Contains(args, "issue") && len(args) <= 6 {
		updateIssue(*ownerPtr, *repoPtr, *editorPtr, args[len(args)-1], client, token)
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

// Create an issue, invoking a text editor to add
func createIssue(owner string, repo string, editor string, title string, client *http.Client, token string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues"
	var req *http.Request
	var err error

	// Optionally write JSON payload using editor, or create an issue just with it's title
	req = generatePayload(http.MethodPost, url, title, "", editor, req)
	
	if err != nil {
		log.Fatal(err)
	} else {
		addHeaders(req, token)
		printResponse(client, req)
	}
}

// Update an issue, invoking a text editor to add
func updateIssue(owner string, repo string, editor string, issueNumber string, client *http.Client, token string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues/" + issueNumber
	fmt.Println(url)
	var req *http.Request
	var err error

	req = generatePayload(http.MethodPatch, url, "", issueNumber, editor, req)
	if err != nil {
		log.Fatal(err)
	} else {
		addHeaders(req, token)
		printResponse(client, req)
	}
}

func generatePayload(httpMethod string, url string, title string, issueNumber string, editor string, req *http.Request) *http.Request {
	var err error
	if editor != "" {
		filename := "tmp.json"
		os.Create("tmp.json")
		cmd := exec.Command(editor, filename)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		
		var reqBodyFile *os.File
		// cmd.Run starts and waits fo the cmd to run
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		} else {
			reqBodyFile, err = os.Open(filename)
		}
		req, err = http.NewRequest(httpMethod, url, reqBodyFile)
		if err != nil {
			log.Fatal(err)
		}
		os.Remove(filename)
		return req
	} else {
		jsonBody , _ := json.Marshal(map[string]string{"title": title})
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Fatal(err)
		}
		return req
	}
} 

// Adds HTTP Headers to request
func addHeaders(req *http.Request, token string) {
	authHeader := "Bearer " + token
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Accept", "application/vnd.github+json")
}

func printResponse(client *http.Client, req *http.Request) {
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
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
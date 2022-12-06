// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"ex4.10/github"
)

func main() {
    orderedResult, totalCount, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", totalCount)
    for _, issuesByAge := range orderedResult {
		fmt.Printf("Issue Age: %s\n", issuesByAge.Category)
		for _, item := range issuesByAge.Issues {
			fmt.Printf("#%-5d %9.9s %.55s\n",
            item.Number, item.User.Login, item.Title)
		}
    }
}

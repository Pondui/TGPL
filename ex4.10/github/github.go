package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
    TotalCount int `json:"total_count"`
    Items      []*Issue
}

type IssuesByAge struct {
	Category string
	Issues []*Issue
}

type Issue struct {
    Number    int
    HTMLURL   string `json:"html_url"`
    Title     string
    State     string
    User      *User
    CreatedAt time.Time `json:"created_at"`
    Body      string    // in Markdown format
}

type User struct {
    Login   string
    HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) ([]*IssuesByAge, int, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesURL + "?q=" + q)
    if err != nil {
        return nil, 0, err
    }

    // We must close resp.Body on all execution paths.
    // (Chapter 5 presents 'defer', which makes this simpler.)
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, 0, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, 0, err
    }
    resp.Body.Close()

	lessThanAMonth := IssuesByAge{
		Category: "Less than a month",
		Issues: []*Issue{},
	}

	lessThanAYear := IssuesByAge{
		Category: "Less than a year",
		Issues: []*Issue{},
	}

	moreThanAYear := IssuesByAge{
		Category: "More than a year",
		Issues: []*Issue{},
	}

	for _, r := range result.Items {
		if r.CreatedAt.After(time.Now().Add(-30*24*time.Hour)) {
			lessThanAMonth.Issues = append(lessThanAMonth.Issues, r)
		} else if r.CreatedAt.Before(time.Now().Add(-30*24*time.Hour)) && r.CreatedAt.After(time.Now().Add(-365*24*time.Hour)) {
			lessThanAYear.Issues = append(lessThanAYear.Issues, r)
		} else {
			moreThanAYear.Issues = append(moreThanAYear.Issues, r)
		}
	}

	var orderedResult []*IssuesByAge
	orderedResult = append(orderedResult, &lessThanAMonth)
	orderedResult = append(orderedResult, &lessThanAYear)
	orderedResult = append(orderedResult, &moreThanAYear)
    return orderedResult, result.TotalCount, nil
}

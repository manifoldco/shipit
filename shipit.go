package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// IssueMessage : a simple struct to hold all the components of an issue
type IssueMessage struct {
	Release string //Message : of issue
}

func main() {

	var (
		owner      = flag.String("o", "manifoldco", "owner name")
		repository = flag.String("r", "engineering", "repositoryName")
		version    = flag.String("v", "", "version tag / provider name")
		fileName   = flag.String("f", "tmpl/marketplace.txt", "template file")
		label      = flag.String("l", "ops", "labels to be used")
	)
	flag.Parse()

	if *version == "" {
		panic("missing version")
	}

	token, err := getToken()
	if err != nil {
		panic(err)
	}

	url, err := makeIssue(*owner, *repository, *version, *fileName, token, *label)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)

}

func getToken() (string, error) {

	val, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return "", errors.New("Token not found")
	}
	return val, nil
}

func makeIssue(owner, repo, version, templateFile, token, label string) (string, error) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	ctx := context.Background()

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return "Template Parse File Issue: ", err
	}
	ghIssue := IssueMessage{version}

	var buf bytes.Buffer

	err = t.ExecuteTemplate(&buf, "Message", ghIssue)
	if err != nil {
		return "Could not make the Message: ", err
	}
	message := buf.String()
	buf.Reset()

	err = t.ExecuteTemplate(&buf, "Title", ghIssue)
	if err != nil {
		return "Could not make the title: ", err
	}
	title := buf.String()
	buf.Reset()

	body := &github.IssueRequest{
		Title:    &title,
		Body:     &message,
		Assignee: &owner,
		Labels:   &[]string{label},
	}

	issue, _, err := client.Issues.Create(ctx, owner, repo, body)
	if err != nil {
		return "Issues.Create returned error", err
	}

	return issue.GetHTMLURL(), nil

}

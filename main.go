package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {

	// the user
	owner := "donuts-are-good"

	// the repo
	repo := "bearclaw"

	// creating a background context as a requirement of oauth
	ctx := context.Background()

	// read the environment variable as a token
	token := os.Getenv("GITHUB_TOKEN")

	// more oauth setup
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// we need to pull results 100 at a time otherwise we get a truncated list
	opt := &github.ListOptions{PerPage: 100}

	// create a place for our stargazers to go
	var allStargazers []*github.Stargazer

	// loop through our stargazers and append the pages
	for {

		// fetch them stargazers
		stargazers, resp, err := client.Activity.ListStargazers(ctx, owner, repo, opt)
		if err != nil {
			fmt.Printf("Error fetching stargazers: %v\n", err)
			os.Exit(1)
		}

		// append to the greater list
		allStargazers = append(allStargazers, stargazers...)

		// stop when we are done
		if resp.NextPage == 0 {
			break
		}

		// but if we aren't done, go to the next page
		opt.Page = resp.NextPage
	}

	// before we print the list, we tell the user how many results there were
	fmt.Printf("Found %d stargazers for %s/%s\n", len(allStargazers), owner, repo)

	// format this into something that will make a grid in your readme
	for _, user := range allStargazers {
		fmt.Printf("<img width=24 height=24 src=\"%s\" alt=\"%s's avatar\" title=\"%s\"/> ", *user.User.AvatarURL, *user.User.Login, *user.User.Login)
	}
}

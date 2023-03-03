package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	owner := "donuts-are-good"
	repo := "bearclaw"

	opt := &github.ListOptions{PerPage: 100}
	var allStargazers []*github.Stargazer

	for {
		stargazers, resp, err := client.Activity.ListStargazers(ctx, owner, repo, opt)
		if err != nil {
			fmt.Printf("Error fetching stargazers: %v\n", err)
			os.Exit(1)
		}

		allStargazers = append(allStargazers, stargazers...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	fmt.Printf("Found %d stargazers for %s/%s\n", len(allStargazers), owner, repo)

	for _, user := range allStargazers {
		fmt.Printf("<img src=\"%s\" alt=\"%s's avatar\" title=\"%s\"/> ", *user.User.AvatarURL, *user.User.Login, *user.User.Login)
	}
}

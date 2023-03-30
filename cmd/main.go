package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

// function to connect to GitHub API with the client as input and return the list of repositories
func getRepos(client *github.Client) ([]*github.Repository, error) {
	ctx := context.Background()
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

// function to prompt user for GitHub token and return the client
func getClient() *github.Client {
	// prompt user for token
	fmt.Println("Enter your GitHub token:")
	var token string
	fmt.Scanln(&token)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

// prompt user for Github repository and owner and PR label and return the list of commits with the label
func getCommits(client *github.Client) ([]*github.RepositoryCommit, error) {
	// prompt user for repository and owner
	fmt.Println("Enter the repository name:")
	var repo string
	fmt.Scanln(&repo)
	fmt.Println("Enter the owner name:")
	var owner string
	fmt.Scanln(&owner)
	// prompt user for label
	fmt.Println("Enter the label name:")
	var label string
	fmt.Scanln(&label)
	ctx := context.Background()
	// get the list of commits with the label
	commits, _, err := client.PullRequests.ListCommits(ctx, owner, repo, 1, &github.ListOptions{})
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func main() {

	// get client
	client := getClient()
	repos, err := getRepos(client)
	if err != nil {
		log.Fatal(err)
	}
	// print the list of repositories
	for _, repo := range repos {
		fmt.Printf("Repo: %s	Stars: %d	Watchers: %d	Forks: %d	Open Issues: %d	\n", *repo.Name, *repo.StargazersCount, *repo.WatchersCount, *repo.ForksCount, *repo.OpenIssuesCount)
	}

	// get the list of commits with the label
	commits, err := getCommits(client)
	if err != nil {
		log.Fatal(err)
	}
	// print the list of commits
	for _, commit := range commits {
		fmt.Printf("Commit: %s	Author: %s	\n", *commit.SHA, *commit.Commit.Author.Name)
	}

}

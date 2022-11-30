package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
)

func main() {
	godotenv.Load()

	tokenAPI := os.Getenv("GITLAB_API_TOKEN")
	baseURL := os.Getenv("GITLAB_API_BASEURL")
	var options []gitlab.ClientOptionFunc
	if baseURL != "" {
		options = append(options, gitlab.WithBaseURL(baseURL))
	}
	client, err := NewClient(tokenAPI, options...)
	if err != nil {
		log.Fatalln(err)
	}
	release, _, err := client.Releases.LatestRelease("sphotonchat%2Fsphotonchat", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(release)
}

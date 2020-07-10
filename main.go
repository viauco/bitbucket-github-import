package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bitbucketUsername := os.Getenv("BITBUCKET_USERNAME")
	bitbucketPassword := os.Getenv("BITBUCKET_PASSWORD")

	gitHubUsername := os.Getenv("GITHUB_USERNAME")
	gitHubPassword := os.Getenv("GITHUB_PASSWORD")
	gitHubOrg := os.Getenv("GITHUB_ORG")

	bitbucketAPIURL := "https://api.bitbucket.org/2.0/repositories/" + bitbucketUsername + "?q=is_private=true"

	scm := map[string]string{
		"git": "git",
		"svn": "subversion",
	}

	for {
		response := &BitbucketRepositoriesResponse{}

		req := &Request{
			Method:   "GET",
			URL:      bitbucketAPIURL,
			Username: bitbucketUsername,
			Password: bitbucketPassword,
		}
		if err := request(req, response); err != nil {
			log.Println(err.Error())
			break
		}

		wg := sync.WaitGroup{}
		wg.Add(len(response.Repositories))

		for _, r := range response.Repositories {
			go func(r BitbucketRepository) {
				defer wg.Done()

				// Check
				if _, ok := scm[r.SCM]; !ok {
					log.Printf("%s: cannot import SCM type: %s\n", r.Name, r.SCM)
					return
				}

				// Create
				repo := GitHubRepository{
					Name:        r.Name,
					Description: r.Description,
					Private:     r.Private,
				}

				req := &Request{
					Method:   "POST",
					URL:      fmt.Sprintf("https://api.github.com/orgs/%s/repos", gitHubOrg),
					Body:     repo,
					Username: gitHubUsername,
					Password: gitHubPassword,
				}

				res := &GitHubError{}

				if err := request(req, res); err != nil {
					log.Printf("%s: error with (%v)", r.Name, res)
					return
				}

				// Import
				imp := GitHubImport{
					VCSURL:      r.Links.HTML.HREF,
					VCS:         scm[r.SCM],
					VCSUsername: bitbucketUsername,
					VCSPassword: bitbucketPassword,
				}

				req = &Request{
					Method:   "PUT",
					URL:      fmt.Sprintf("https://api.github.com/repos/%s/%s/import", gitHubOrg, r.Name),
					Body:     imp,
					Header:   make(map[string]string),
					Username: gitHubUsername,
					Password: gitHubPassword,
				}
				req.Header["Accept"] = "application/vnd.github.barred-rock-preview"

				if err := request(req, res); err != nil {
					log.Printf("%s: error (%v)", r.Name, res)
					return
				}

				log.Printf("Import was started: https://github.com/%s/%s\n\n", gitHubOrg, r.Name)
			}(r)
		}

		wg.Wait()

		if response.Next == "" {
			break
		}

		bitbucketAPIURL = response.Next
	}
}

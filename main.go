package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cli/go-gh/v2"
)

func main() {
	// get user
	userBuf, _, err := gh.Exec("api", "user", "--jq", ".login")
	user := userBuf.String()
	if err != nil {
		log.Fatalln(err)
	}

	repo := fmt.Sprintf("%s/nixpkgs-review-gha", strings.TrimSpace(user))

	// check if user has a fork of nixpkgs-review-gha
	_, _, err = gh.Exec("repo", "view", repo)
	if err != nil {
		// create fork if it doesn't exist
		_, _, err = gh.Exec("repo", "fork", "Defelo/nixpkgs-review-gha")
	}

	// enable action
	_, _, _ = gh.Exec("workflow", "enable", "-R", repo, "review")

	// run action

}

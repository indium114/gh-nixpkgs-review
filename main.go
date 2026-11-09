package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cli/go-gh/v2"
	"github.com/indium114/slag"
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
	slag.Log("checking that you have a fork of nixpkgs-review-gha\n")
	_, _, err = gh.Exec("repo", "view", repo)
	if err != nil {
		// create fork if it doesn't exist
		slag.Hint("forking nixpkgs-review-gha\n")
		_, _, err = gh.Exec("repo", "fork", "Defelo/nixpkgs-review-gha")
	}

	// enable action
	slag.Hint("enabling 'review' workflow\n")
	_, _, _ = gh.Exec("workflow", "enable", "-R", repo, "review")

	// run action
	slag.Hint("dispatching action\n")

}

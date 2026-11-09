package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cli/go-gh/v2"
	"github.com/indium114/slag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(slag.Err("please provide the pr # as an argument").Error())
		os.Exit(1)
	}
	pr := os.Args[1]

	// get user
	userBuf, _, err := gh.Exec("api", "user", "--jq", ".login")
	user := userBuf.String()
	if err != nil {
		fmt.Println(slag.Err(err.Error()).Error())
		os.Exit(1)
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
	out, _, _ := gh.Exec("workflow", "run", "-R", repo, "review",
		"-f", "x86_64-linux=true",
		"-f", "aarch64-linux=true",
		"-f", "aarch64-darwin=yes_sandbox_relaxed",
		"-f", "allow-unfree=true",
		"-f", "allow-insecure=true",
		"-f", "builders=gha",
		"-f", "push-to-cache=false",
		"-f", "upterm=false",
		"-f", "post-result=true",
		"-f", "on-success=nothing",
		"-f", "pr="+pr)

	url := strings.TrimSpace(out.String())
	slag.Ok("workflow dispatched; vist %s to see its status", url)
}

package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2"
)

func main() {
	userBuf, _, err := gh.Exec("api", "user", "--jq", ".login")
	user := userBuf.String()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}

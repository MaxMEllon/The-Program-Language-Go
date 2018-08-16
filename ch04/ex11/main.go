package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MaxMEllon/The-Program-Language-Go/ch04/ex11/github"
)

func main() {
	var user, repo string
	flag.StringVar(&user, "user", "MaxMEllon", "github user")
	flag.StringVar(&repo, "repo", "The-Program-Language-Go", "github repository")
	flag.Parse()

	cmds := flag.Args()
	if len(cmds) == 0 {
		log.Fatal("sub command [get create]")
	}

	switch cmds[0] {
	case "get":
		if len(cmds) == 1 {
			c, err := github.NewClient("")
			if err != nil {
				log.Fatal(err)
			}
			issues, err := c.GetIssues(repo, user)
			if err != nil {
				log.Fatal(err)
			}
			for _, item := range *issues {
				fmt.Printf("- #%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
			}

		} else {
			c, err := github.NewClient("")
			if err != nil {
				log.Fatal(err)
			}
			issue, err := c.GetIssue(repo, user, cmds[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("- #%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
		}
	}
}

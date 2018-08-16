package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MaxMEllon/The-Program-Language-Go/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var month, year, after []*github.Issue
	now := time.Now()
	for _, item := range result.Items {
		hours := now.Sub(item.CreatedAt).Hours()
		if hours < 24*31 {
			month = append(month, item)
			continue
		}
		if hours < 24*365 {
			year = append(year, item)
			continue
		}
		after = append(after, item)
	}

	fmt.Printf("# issues: %d\n", len(result.Items))

	fmt.Println("\n今月\n---")
	for _, item := range month {
		fmt.Printf("- [#%-5d](%-70s) %9.9s %.55s\n", item.Number, item.HtmlURL, item.User.Login, item.Title)
	}
	fmt.Println("\n今年\n---")
	for _, item := range year {
		fmt.Printf("- [#%-5d](%-70s) %9.9s %.55s\n", item.Number, item.HtmlURL, item.User.Login, item.Title)
	}
	fmt.Println("\n1年以上前\n---")
	for _, item := range after {
		fmt.Printf("- [#%-5d](%-70s) %9.9s %.55s\n", item.Number, item.HtmlURL, item.User.Login, item.Title)
	}
}

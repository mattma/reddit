package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mattma/reddit/geddit"
)

func main() {
	q := flag.String("q", "golang", "query for subreddit")
	flag.Parse()

	items, err := reddit.Get(*q)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}

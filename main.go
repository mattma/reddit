package main

import (
  "fmt"
  "log"

  "github.com/mattma/reddit/geddit"
)

func main() {
  items, err := reddit.Get("golang")

  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item)
  }
}

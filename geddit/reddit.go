// Package reddit implements a basic client for the Reddit API
// this is the document the package itself
package reddit

import (
  "encoding/json"
  "errors"
  "fmt"
  "net/http"
)

type Item struct {
  Title string
  Url   string
  // Comments field has a "struct tag", a string that annotates the field
  // Go use `reflect` package to inspect this info at runtime
  // Note: there is no space in between "json" and "key"
  Comments int `json:"num_comments"`
}

// decode JSON-encoded data into native Go data structure
type response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

// Item describes a Reddit item
// when pass value to `fmt.Print`, will check if it implements the `fmt.Stringer` interface
// Any type that implements a `String() string` method is a Stringer
// `fmt` package will use that method to format values of that type
// A method declaration is like a func but the receiver comes first
func (i Item) String() string {
  com := ""
  switch i.Comments {
  case 0:
    // Nothing
  case 1:
    com = " (1 comment)"
  default:
    com = fmt.Sprintf(" (%d comments)", i.Comments)
  }
  return fmt.Sprintf("%s %s\n%s\n", i.Title, com, i.Url)
}

// Get fetches the most recent Items posted to the specified subreddit
func Get(reddit string) ([]Item, error) {
  // construct the request URL and return a value
  url := fmt.Sprintf("https://www.reddit.com//r/%s.json", reddit)
  resp, err := http.Get(url)

  if err != nil {
    return nil, err
  }

  // Clean up HTTP request, when function exit, shut down underlying TCP connetion
  // this func will be made after the function returns
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return nil, errors.New(resp.Status)
  }

  // Copy HTTP reponse body to standard output
  // resp.Body type `io.Reader`, os.Stdout type is `io.Writer`
  // _, err = io.Copy(os.Stdout, resp.Body)

  // Initialize a new Response value, store a pointer to variable `r`
  r := new(response)
  // Create a new decoder object and decode the response body into `r`
  // pass in Response struct to `resp.Body`
  err = json.NewDecoder(resp.Body).Decode(r)

  if err != nil {
    return nil, err
  }

  items := make([]Item, len(r.Data.Children))

  // print the data
  for i, child := range r.Data.Children {
    items[i] = child.Data
  }

  return items, nil
}

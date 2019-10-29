package main

import "golang.org/x/net/html"

html := `
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
`

type Link struct {
	Href string
	Text string
}
func main() {
	html.NewTokenizer()
}

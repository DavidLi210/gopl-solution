package exe

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		findlinks3(url)
	}
}

func findlinks3(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	fmt.Println(findElementById(doc, "data") != nil)
}

func findElementById(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

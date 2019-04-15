package exe

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	Helper(os.Args[1])
}

func Helper(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	res := make([]*html.Node, 0)
	res = ElementsByTagName(doc, os.Args[2:]...)
	for _, c := range res {
		fmt.Printf("%s %v \n", c.Data, c.Type)
	}
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	res := make([]*html.Node, 0)
	for _, name := range names {
		if doc.Type == html.ElementNode && doc.Data == name {
			res = append(res, doc)
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, ElementsByTagName(c, names...)...)
	}
	return res
}

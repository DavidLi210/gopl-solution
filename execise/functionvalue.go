package exe

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		findlinks2(url)
	}
}

func findlinks2(url string) {
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
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	if n.Type == html.CommentNode {
		fmt.Println("Comment Tag:" + n.Data)
	}
	if n.Type == html.TextNode {
		fmt.Println("Text Tag:" + n.Data)
	}
	if n.Type == html.DoctypeNode {
		fmt.Println("DoctypeNode Tag:" + n.Data)
	}
	if n.Type == html.ElementNode {
		fmt.Println("ElementNode Tag:" + n.Data)
	}
	visitAttribute(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func endElement(n *html.Node) {
	depth--
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func visitAttribute(n *html.Node) {
	fmt.Println("Start:====>")
	for _, attr := range n.Attr {
		fmt.Printf("[%v]=%v\n", attr.Key, attr.Val)
	}
	fmt.Println("End:====>")
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

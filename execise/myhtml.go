package exe

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var sums = make(map[string]int)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		fmt.Println(images)
		fmt.Println(words)
		fmt.Println("words:" + string(words) + " images:" + string(images))
	} // https://www.zhihu.com/

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func CountWordsAndImages(url string) (words, images int, err error) {
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
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(node *html.Node) (int, int) {
	if node.Type == html.ElementNode {
		if node.Data == "img" {
			return 0, 1
		}
		if node.Data == "a" {
			return 1, 0
		}
	}
	res := []int{0, 0}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		t1, t2 := countWordsAndImages(c)
		res[0] = res[0] + t1
		res[1] = res[1] + t2
	}
	return res[0], res[1]
}

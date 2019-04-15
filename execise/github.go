package exe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int
	Items      []*Issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"create_at"`
	Body     string
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	recorder1 := make([]string, 0)
	recorder2 := make([]string, 0)
	recorder3 := make([]string, 0)
	for _, item := range result.Items {
		//fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", item.CreateAt.Year(), item.CreateAt.Month(), item.CreateAt.Day(), item.CreateAt.Hour(), item.CreateAt.Minute(), item.CreateAt.Second())
		fmt.Println(item.CreateAt)
		fmt.Printf("#%5d%9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		if time.Now().Sub(item.CreateAt).Hours() <= 30*24 {
			recorder1 = append(recorder1, item.Title)
		} else if time.Now().Sub(item.CreateAt).Hours() <= 30*24*12 {
			recorder2 = append(recorder2, item.Title)
		} else {
			recorder3 = append(recorder3, item.Title)
		}
	}
	fmt.Println("<======================================>")
	fmt.Println(recorder1)
	fmt.Println(recorder2)
	fmt.Println(recorder3)
}

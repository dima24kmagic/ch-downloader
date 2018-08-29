package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/dima24kmagic/course_hunters/page"

	"golang.org/x/net/html"
)

type Link struct {
	URL string `xml:"href"`
}

// GetBody - Get body tag of a page
func GetDownloadLinks(doc *html.Node) ([]*html.Node, error) {
	b := make([]*html.Node, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "link" {
			b = append(b, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n []*html.Node) []string {
	nodes := make([]string, 0)
	for _, node := range n {
		var buf bytes.Buffer
		w := io.Writer(&buf)
		html.Render(w, node)
		nodes = append(nodes, buf.String())
	}
	courses := filterLists(nodes)
	return courses
}

func main() {
	page := page.GetOfflinePage()
	doc, _ := html.Parse(strings.NewReader(string(page)))
	bn, err := GetDownloadLinks(doc)
	if err != nil {
		return
	}
	courseLists := renderNode(bn)

	courseLink := make(chan string)
	done := make(chan bool)
	go func() {
		for {
			select {
			case link := <-courseLink:
				fmt.Println(link)
			case <-done:
				return
			}
		}
	}()
	func() {
		for _, v := range courseLists {

			reader := strings.NewReader(v)
			z := html.NewTokenizer(reader)
			func() {
				for {
					tt := z.Next()
					switch {
					case tt == html.SelfClosingTagToken:
						t := z.Token()
						for _, a := range t.Attr {
							if a.Key == "href" {
								fmt.Println(a.Val)
								return
							}
						}
					}
				}
			}()
		}
		close(done)
	}()

}

// Filter li tags for those, that contain class .lessons-list__li

func filterLists(allLists []string) []string {
	courseLists := make([]string, 0)
	for _, v := range allLists {
		// check classname
		if strings.Contains(v, "contentUrl") {
			courseLists = append(courseLists, v)
		}
	}
	return courseLists
}

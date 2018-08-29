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

// GetBody - Get body tag of a page
func GetBody(doc *html.Node) ([]*html.Node, error) {
	b := make([]*html.Node, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
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
	bn, err := GetBody(doc)
	if err != nil {
		return
	}
	courseLists := renderNode(bn)

	fmt.Println(courseLists)

}

// Filter li tags for those, that contain class .lessons-list__li

func filterLists(allLists []string) []string {
	courseLists := make([]string, 0)
	for _, v := range allLists {
		// check classname
		if strings.Contains(v, "lessons-list__li") {
			fmt.Println(v)
			courseLists = append(courseLists, v)
		}
	}
	fmt.Println("LEN is:", len(courseLists))
	return courseLists
}

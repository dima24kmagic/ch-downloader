package page

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/dima24kmagic/course_hunters/courses"
	"golang.org/x/net/html"
)

// GetAllCourses - get courses from http links
func GetAllCourses() []courses.Course {
	page := GetOfflinePage()
	doc, _ := html.Parse(strings.NewReader(string(page)))
	bn, err := getDownloadLinks(doc)
	if err != nil {
		log.Println("Error", err)
	}
	courseLists := renderNode(bn)

	courses := getCoursesFromList(courseLists)
	return courses
}

// getDownloadLinks - Get all the links from page tag of a page
func getDownloadLinks(doc *html.Node) ([]*html.Node, error) {
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

// renderNode - get slice of all links with course mo4 link
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

// getCoursesFromList - create course objects from existing links
// TODO: get the coursename from list
func getCoursesFromList(courseLists []string) []courses.Course {
	allCourses := make([]courses.Course, 0)
	for i, v := range courseLists {
		newCourse := courses.Course{}
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
							newCourse.CourseURL = a.Val
							num := strconv.Itoa(i + 1)
							newCourse.CourseName = "lesson " + num
							newCourse.CourseNum = i + 1
							allCourses = append(allCourses, newCourse)
							return
						}
					}
				}
			}
		}()
	}
	return allCourses
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

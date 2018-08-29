package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dima24kmagic/course_hunters/page"
)

// DESCRIPTION:
// This is source code for downloading courses from course_hunters and gzip them into archive

// 1) Get the page of a course
// 2) Traversal through html and get coursename + courseurl
// 3) Download all courses with proper name
// 3) For all courses create gzip archive

// Course - struct for saving course
type Course struct {
	CourseName string
	CourseURL  string
}

type Html struct {
	Body Link `xml:"body"`
}
type Link struct {
	Content string `xml:",innerxml"`
}

func main() {
	// get the html page
	// filename := "Lesson1"
	// filepath := "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson1.mp4"
	// download(filename, filepath)
	// courseHTML := getCourse("https://coursehunters.net/course/terminal-i-komandnaya-stroka-dlya-ne-tehnarey")
	env := os.Getenv("TEST_ENV")
	fmt.Println(env)
	courseHTML := page.GetOfflinePage()
	scrapeThroughCourse(courseHTML)
}

func scrapeThroughCourse(b []byte) {
	fmt.Println(string(b))
	h := Html{}
	err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
	if err != nil {
		log.Println("Error scraping page", err)
	}
	fmt.Println("HTML Response:", h)
}

func getCourse(courseLink string) []byte {
	resp, err := http.Get(courseLink)
	if err != nil {
		log.Println("Can't get the course", err)
	}
	defer resp.Body.Close()
	rb, _ := ioutil.ReadAll(resp.Body)
	return rb
}

func download(filename, filepath string) {

	// Create the file
	out, err := os.Create("./" + filename + ".mp4")
	if err != nil {
		log.Println("Cannot create file", err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(filepath)
	if err != nil {
		log.Println("Cannot get file", err)
	}
	defer resp.Body.Close()

	// Write the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Cannot copy body of a response to a file", err)
	}
}

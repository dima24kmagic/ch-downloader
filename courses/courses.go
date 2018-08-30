package courses

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Course - main struct for course info (№, name, url)
type Course struct {
	CourseNum  int
	CourseName string
	CourseURL  string
}

// Download - download video from given courseVide type
func (course *Course) Download() {
	// Create the file
	filename := "./" + course.CourseName + ".mp4" // ex 1) Learning Terminal.mp4
	out, err := os.Create(filename)
	if err != nil {
		log.Println("Cannot create file", err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(course.CourseURL)
	if err != nil {
		log.Println("Cannot get file", err)
		return
	}
	defer resp.Body.Close()

	// Write the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Cannot copy body of a response to a file", err)
	}
}

func AskForDownload() []int {
	var cNumsForDownload []int

	fmt.Println("Download all courses (Y/N)")
	var allCourses bool
	var answer string
	fmt.Scan(&answer)
	allCourses = checkAnswer(answer)
	fmt.Println(allCourses)

	return cNumsForDownload
}

func checkAnswer(a string) bool {
	if strings.ToLower(a) == "y" || strings.ToLower(a) == "yes" {
		return true
	}
	return false
}

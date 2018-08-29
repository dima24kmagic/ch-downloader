package courses

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Course - main struct for course info (№, name, url)
type Course struct {
	CourseNum  int64
	CourseName string
	CourseURL  string
}

// Download - download video from given courseVide type
func Download(course Course) {
	// Create the file	
	filename := "./" + strconv.FormatInt(course.CourseNum, 64) + ") " + course.CourseName + ".mp4" // ex 1) Learning Terminal.mp4
	out, err := os.Create(filename)
	if err != nil {
		log.Println("Cannot create file", err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(course.CourseURL)
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

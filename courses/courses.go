package courses

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Course - main struct for course info (№, name, url)
type Course struct {
	CourseNum  int
	CourseName string
	CourseURL  string
}

func CreateFolderForCourses() {
	os.Mkdir("./lessons", 0777)
}

// Download - download video from given courseVide type
func (course *Course) Download() {
	// Create the file
	path := "./lessons/" + course.CourseName + ".mp4"
	filename := path // ex 1) Learning Terminal.mp4
	out, err := os.Create(filename)
	if err != nil {
		log.Println("Cannot create file", err)
	}
	defer out.Close()

	// Get size of file
	headResp, _ := http.Head(course.CourseURL)
	defer headResp.Body.Close()

	size, _ := strconv.Atoi(headResp.Header.Get("Content-length"))
	fmt.Println("Size of", course.CourseName, "is", float64(size)/1000000, "mb")
	fmt.Println("Download", course.CourseName)

	start := time.Now()
	done := make(chan int64)

	go PrintDownloadPercent(done, path, int64(size))

	// Get the data
	resp, err := http.Get(course.CourseURL)
	if err != nil {
		log.Println("Cannot get file", err)
		return
	}
	defer resp.Body.Close()

	// Write the body to the file
	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Cannot copy body of a response to a file", err)
	}

	done <- n

	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)
}

func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent = float64(size) / float64(total) * 100

			fmt.Printf("%.0f", percent)
			fmt.Print("=>")
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

func AskForDownload(courses []Course) []int {
	var cNumsForDownload []int

	answer := getInput("Download all courses (Y/N)")
	allCourses := checkAnswer(answer)
	if allCourses {
		for _, v := range courses {
			cNumsForDownload = append(cNumsForDownload, v.CourseNum)
		}
	} else {
		a := getInput("What courses u wannna to download", "ex. --> 2 + 4-15 + 17-21")
		parseInputForNumbers(a)
	}

	return cNumsForDownload
}

func parseInputForNumbers(input string) []string {
	coursePositions := make([]string, 0)
	for _, letter := range input {
		// TODO: While not space - add this to a slice of parsed inputs, from there i can see, is that range or not
		fmt.Println(string(letter))
	}
	return coursePositions
}

func getInput(prompt ...string) string {
	for _, v := range prompt {
		fmt.Println(v)
	}
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] scanning", err)
	}
	return line
}

func checkAnswer(a string) bool {
	if strings.ToLower(a) == "y" || strings.ToLower(a) == "yes" {
		return true
	}
	return false
}

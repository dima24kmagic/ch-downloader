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
		fmt.Println("all courses gonna be downloaded")
		return cNumsForDownload
	} else {
		a := getInput("What courses u wannna to download", "ex. --> 2 + 4-15 + 17-21")
		args := parseInputForNumbers(a)
		positions := getCoursesPosition(args)
		fmt.Println("Courses with th-is positions gonna be downloaded: ", positions)
		return positions
	}
}

func getCoursesPosition(args []string) []int {
	positions := make([]int, 0)
	for _, v := range args {
		if strings.Contains(v, "-") {
			ns := parseRange(string(v))
			for _, n := range ns {
				positions = append(positions, n)
			}
		} else {
			n, _ := strconv.Atoi(v)
			positions = append(positions, n)
		}
	}
	return positions
}

func parseRange(argRange string) []int {
	ns := make([]int, 0)
	var arg string
	args := make([]string, 0)
	for i, v := range argRange {
		if i+1 == len(argRange) {
			arg += string(v)
			args = append(args, arg)
		} else if v == '-' {
			args = append(args, arg)
			arg = ""
		} else {
			arg += string(v)
		}
	}

	start, err := strconv.Atoi(args[0])
	end, err := strconv.Atoi(args[1])

	if err != nil {
		log.Println("[ERROR] Conversion", err)
	}
	for i := start; i <= end; i++ {
		ns = append(ns, i)
	}
	return ns
}

func removeSpaceFromArgs(args []string) []string {
	resp := make([]string, 0)
	for _, arg := range args {
		s := strings.Trim(arg, " ")
		resp = append(resp, s)
	}
	return resp
}

// parseInputForNumbers - get array of arguments, (one or range *with dash*)
func parseInputForNumbers(input string) []string {
	allInputs := make([]string, 0)
	var arg string
	for i, letter := range input {
		if i+1 == len(input) {
			allInputs = append(allInputs, arg)
		} else if letter != '+' { // for input delimetr is + sign
			arg += string(letter)
		} else {
			allInputs = append(allInputs, arg)
			arg = ""
		}
		// fmt.Println("[LETTER]", string(letter))
	}
	resp := removeSpaceFromArgs(allInputs)
	return resp
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
	if trimAndLow(a) == "y" || trimAndLow(a) == "yes" {
		return true
	}
	return false
}

func trimAndLow(a string) string {
	return strings.ToLower(strings.TrimRight(a, "\n"))
}

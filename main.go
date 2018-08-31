package main

import (
	"fmt"

	"github.com/dima24kmagic/course_hunters/courses"
	"github.com/dima24kmagic/course_hunters/page"
)

func main() {
	courses.CreateFolderForCourses()
	cs := page.GetAllCourses()
	downloadNo := courses.AskForDownload(cs)

	done := make(chan bool)
	checkDownloadClose(downloadNo, cs, done)
	select {
	case <-done:
		fmt.Println("All Courses Downloaded!")
		return
	}
}

func checkDownloadClose(downloadNo []int, cs []courses.Course, done chan bool) {
	for _, course := range cs {
		for _, no := range downloadNo {
			if course.CourseNum == no {
				course.Download()
			}
		}
	}
	close(done)
}

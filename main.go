package main

import (
	"github.com/dima24kmagic/course_hunters/courses"
	"github.com/dima24kmagic/course_hunters/page"
)

func main() {
	courses.CreateFolderForCourses()
	cs := page.GetAllCourses()
	// courses.AskForDownload(cs)
	for _, course := range cs {
		course.Download()
		break
	}
}

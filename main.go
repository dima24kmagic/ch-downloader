package main

import (
	"github.com/dima24kmagic/course_hunters/page"
)

func main() {
	cs := page.GetAllCourses()
	for _, course := range cs {
		course.Download()
	}
}

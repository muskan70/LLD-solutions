package main

import (
	"fmt"
	"log"
)

var customers map[string]*Customer
var courses map[string]Course

func main() {
	customers = make(map[string]*Customer)
	admin := NewAdmin("muskan", "9878353465", "mini45@gmail.com")
	courses = make(map[string]Course)

	if err := admin.AddCourse("C++", []string{}); err != nil {
		log.Println(err.Error())
	}

	if err := admin.AddCourse("Python", []string{"C++"}); err != nil {
		log.Println(err.Error())
	}

	if err := admin.AddCourse("Go", []string{"C++", "Python"}); err != nil {
		log.Println(err.Error())
	}
	log.Println(courses)

	NewCustomer("vipul", "5678364432", "vips67@gmail.com")
	NewCustomer("manju", "5679994432", "man97@gmail.com")
	NewCustomer("yash", "9998364432", "yash72@gmail.com")

	if precourses, err := customers["vipul"].OptForCourse("C++"); err != nil {
		log.Println(err, precourses)
	} else {
		log.Println("course opted")
	}

	customers["vipul"].UpdateCourseStatus("C++")
	if precourses, err := customers["vipul"].OptForCourse("Go"); err != nil {
		log.Println(err, precourses)
	} else {
		log.Println("course opted")
	}

	if err := admin.UpdateCourse("Python", []string{}); err != nil {
		log.Println(err.Error())
	}
	fmt.Println(courses)
}

/*
Curriculum manager:
There are  number of course
Admin have access to add/delete/update course.
Each course may have some pre-requisites, ds/algo -> python.
Pre-requisites course should be completed before going for the parent course.

Actors : Admin, users
user opt for course , list of courses of users
topological sort

Entities:

User
-name
-phone
-email
-password

UserCourse
 	-CourseId int
-Status -> fresh, completed

Customer
-User (inherited)
-Courses UserCourse
+GetAllUserCourses(userid)
+OptForCourse(userID,courseId) [] prerequisiteCourses –

Admin
-User(inherited)

Course
-CourseId
-name
-listofPrerequisiteCourseIds
+Add(course)
+Delete(courseId)
+Update(courseId, prerequisites) –
*/

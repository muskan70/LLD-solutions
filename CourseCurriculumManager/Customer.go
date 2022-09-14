package main

import "errors"

type Customer struct {
	User
	Courses map[string]string
}

func NewCustomer(name, phone, email string) {
	usr := NewUser(name, phone, email, UserTypeCustomer)
	customers[name] = &Customer{User: usr, Courses: make(map[string]string)}
}

func (c *Customer) OptForCourse(courseName string) ([]string, error) {
	csr, ok := courses[courseName]
	if !ok {
		return nil, errors.New("invalid course")
	}
	var prerequisites []string
	for _, preCourse := range csr.PrerequisiteCourses {
		if status, ok := c.Courses[preCourse]; ok && status == UserCourseStatusCompleted {
			continue
		}
		prerequisites = append(prerequisites, preCourse)
	}
	if len(prerequisites) > 0 {
		return prerequisites, errors.New("you have to complete some courses before opting for this course")
	}
	c.Courses[courseName] = UserCourseStatusNew
	return nil, nil
}

func (c *Customer) UpdateCourseStatus(courseName string) error {
	if _, ok := c.Courses[courseName]; !ok {
		return errors.New("invalid course for user")
	}
	c.Courses[courseName] = UserCourseStatusCompleted
	return nil
}

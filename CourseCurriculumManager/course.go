package main

import "errors"

var c = 0

type Course struct {
	CourseId            int
	Name                string
	PrerequisiteCourses []string
}

func NewCourse(name string, preCourses []string) (*Course, error) {
	c++
	for _, i := range preCourses {
		if _, ok := courses[i]; !ok {
			return nil, errors.New("invalid Prerequisite Courses")
		}
	}
	crs := Course{CourseId: c, Name: name, PrerequisiteCourses: preCourses}
	return &crs, nil
}

func (crs *Course) UpdateCoursePrerequisites(preCourses []string) error {
	for _, i := range preCourses {
		if _, ok := courses[i]; !ok {
			return errors.New("invalid Prerequisite Courses")
		}
	}
	crs.PrerequisiteCourses = preCourses
	return nil
}

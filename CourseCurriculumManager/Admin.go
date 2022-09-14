package main

import "errors"

type Admin struct {
	User
}

func NewAdmin(name, phone, email string) Admin {
	usr := NewUser(name, phone, email, UserTypeAdmin)
	return Admin{User: usr}
}

func (a *Admin) AddCourse(name string, precourses []string) error {
	crs, err := NewCourse(name, precourses)
	if err != nil {
		return err
	}
	courses[name] = *crs
	return nil
}

func (a *Admin) DeleteCourse(name string) {
	delete(courses, name)
}

func (a *Admin) UpdateCourse(name string, precourses []string) error {
	crs, ok := courses[name]
	if !ok {
		return errors.New("inavlid course")
	}
	err := crs.UpdateCoursePrerequisites(precourses)
	if err != nil {
		return err
	}
	courses[name] = crs
	return nil
}

package main

import "errors"

type User struct {
	Name                string
	BookedActivitySlots []WorkoutSession
}

func registerUser(name string) error {
	if len(name) == 0 {
		return errors.New("invalid name")
	}
	user := User{Name: name}
	users[name] = &user
	return nil
}

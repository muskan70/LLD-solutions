package main

import (
	"dotpe/projects/usecase"
	"fmt"
)

func main() {
	fmt.Println("Lets create a N-Elevator syatem. Please enter the mentioned details:")

	var startFloor, endFloor, elevators int
	fmt.Println("Enter StartFloor:")
	fmt.Scanln(&startFloor)
	fmt.Println("Enter EndFloor:")
	fmt.Scanln(&endFloor)
	fmt.Println("Enter No.of elevators:")
	fmt.Scanln(&elevators)

	dispacher := usecase.NewDispacherUcase(startFloor, endFloor, elevators)
	fmt.Println(dispacher.Elevators[2].Direction)
}

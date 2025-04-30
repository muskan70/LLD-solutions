package main

import "flip/usecase"

func main() {
	usecase.NewCentreUsecase()
	usecase.NewUserUsecase()
	usecase.NewSlotUsecase()
	test()
}

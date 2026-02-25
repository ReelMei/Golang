package main

import "fmt"

func main() {
	// ARRAYS

	var name = [3]string{"Enzo", "John", "Kito"}
	var cars = [4]string{"Benz", "Toyota", "Lexus", "Audi"}
	cars[3] = "RTOS"

	fmt.Println(name)
	fmt.Println(cars, len(cars))

	// SLICES

}

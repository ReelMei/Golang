package main

import "fmt"

func main() {
	// ARRAYS

	var name = [3]string{"Enzo", "John", "Kito"}
	var cars = []string{"Benz", "Toyota", "Lexus", "Audi"}
	cars[3] = "RTOS"

	cars = append(cars, "Lamborghini")

	fmt.Println(name[0])
	fmt.Println(cars, len(cars))

	// SLICES
	var scores = []int{23, 25, 18, 21, 16, 21, 19}
	scores[5] = 20

	scores = append(scores, 30)

	fmt.Println(scores)

	// SLICES RANGE
	rangeOne := cars[2:]

	fmt.Println(rangeOne)
}

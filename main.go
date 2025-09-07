package main

import "fmt"

var globalname string = "Global" // this is a gloabl var and can be used in all the main package

func main() {
	var name string = "Harry"

	// var name3 = "Jass" // can be use without saying it is string because we make a string var once inside the func

	var number int = 23 // can use int8-16-32-64 for memory

	//var number2 uint = 22 // can not use as negetive number

	var number3 float64 = 74691.3757 // can only have float64 and float32

	name2 := "AI" // this can only be used inside the function

	// Arrays
	var numbers[6]int = [6]int{23, 25, 27, 28, 32, 56}
	numbers[5] = 52
	var scores = []int{34, 46, 756, 12}
	scores = append(scores, 14)
	rangenumbers := numbers[0:3]

	fmt.Printf("Hello my name is %v, my age is somrthing betwwen %v and %v and you know what %v %v \n",
	name, number, number3, name2, name2)
	fmt.Print("Lets Rock")
	fmt.Println("\nOK?")
	fmt.Println(numbers, len(numbers))
	fmt.Println(scores)
	fmt.Println(rangenumbers)

	// fmt.Print("Hello") // without ln, dosent add a new line
	// fmt.Print(name)

}


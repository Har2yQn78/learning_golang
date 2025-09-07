package main

import (
	"fmt"
	"sort"
	"strings"
)

var globalname string = "Global" // this is a gloabl var and can be used in all the main package

func main() {
	var name string = "Harry"
	 var greeting = "Hello World"
	// var name3 = "Jass" // can be use without saying it is string because we make a string var once inside the func
	var number int = 23 // can use int8-16-32-64 for memory
	//var number2 uint = 22 // can not use as negetive number
	var number3 float64 = 74691.3757 // can only have float64 and float32
	name2 := "AI" // this can only be used inside the function
	// Arrays
	var numbers [6]int = [6]int{23, 25, 27, 28, 32, 56}
	numbers[5] = 52
	var scores = []int{34, 46, 756, 12}
	scores = append(scores, 14)
	rangenumbers := numbers[0:3]
	ages := []int{25, 23, 24, 22, 21, 45, 56, 75, 78}

	sort.Ints(ages) // sort numbers in order  -  can also worked for string  - typical basic methods
	index := sort.SearchInts(ages, 24)

	fmt.Printf("Hello my name is %v, my age is somrthing betwwen %v and %v and you know what %v %v \n",
		name, number, number3, name2, name2)
	fmt.Print("Lets Rock")
	fmt.Println("\nOK?")
	fmt.Println(numbers, len(numbers))
	fmt.Println(scores)
	fmt.Println(rangenumbers)
	fmt.Println(strings.Contains(greeting, "Hello")) // just some basic import and package
	fmt.Println(strings.ReplaceAll(greeting, "World", "Harry"))
	fmt.Println(strings.ToUpper(greeting)) //typical upper method in all lang
	fmt.Println(strings.Index(greeting, "W")) 
	fmt.Println(ages)
	fmt.Println(index)

	// fmt.Print("Hello") // without ln, dosent add a new line
	// fmt.Print(name)

}

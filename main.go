package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Greeting(n string) {
	fmt.Printf("Hello There\nGeneral %v\n", n)
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

var globalname string = "Global" // this is a gloabl var and can be used in all the main package

func main() {
	var name string = "Harry"
	var greeting = "Hello World"
	// var name3 = "Jass" // can be use without saying it is string because we make a string var once inside the func
	var number int = 23 // can use int8-16-32-64 for memory
	//var number2 uint = 22 // can not use as negetive number
	var number3 float64 = 74691.3757 // can only have float64 and float32
	name2 := "AI"                    // this can only be used inside the function
	// Arrays
	var numbers [6]int = [6]int{23, 25, 27, 28, 32, 56}
	a1 := CircleArea(4.3)
	a2 := CircleArea(10.5)
	numbers[5] = 52
	var scores = []int{34, 46, 756, 12}
	scores = append(scores, 14)
	rangenumbers := numbers[0:3]
	ages := []int{25, 23, 24, 22, 21, 45, 56, 75, 78}
	i := 0
	menu := map[string]float64{
		"coffee": 2.5,
		"iced coffee": 3.0,
		"latte": 2.8,
		"iced latte": 3.5,
	}

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

	for i <= 3 {
		fmt.Println("i =", i)
		i++
	}

	for f := 0; f < 4; f++ {
		fmt.Println("f =", f)
		f++
	}

	for i := 1; i <= 3; i++ {
		// Inner loop: controls columns
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d × %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	size := 4

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := 1; i <= 3; i++ {
		for j := 3; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 2; i >= 1; i-- {
		for j := 3; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println(number <= 24)
	fmt.Println(number != 23)

	if number < 25 {
		fmt.Println("So What?")
	} else if number > 25 {
		fmt.Println("Also So What?")
	} else {
		fmt.Println("for real, So What?")
	}

	Greeting("Loop Breaker")

	fmt.Println(a1, "\n", a2)
	fmt.Println(menu)
	fmt.Println(menu["coffee"])

	for k, v := range menu{
		fmt.Println(k,"-",v)
	}

	// some typical for loop to draw shapes
	// fmt.Print("Hello") // without ln, dosent add a new line
	// fmt.Print(name)
}

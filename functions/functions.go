package main

import (
	"fmt"
)

func main() {

	//call variadic function
	test("kunal", "taitkar", 1, 2, 3, 4, 5)

	// function as a parameter
	var oneFunc func(n int) int // default value is nil.
	oneFunc = myFunction1

	twoFunc := myFunction2

	fmt.Println("ONE Function:", oneFunc(1))
	fmt.Println("ONE Function:", twoFunc(5, myFunction1))

	//anonymous  function
	anonymousFunc := func(n int) func() bool {
		return func() bool {
			return n%2 == 0
		}
	}

	is5Even := anonymousFunc(5)
	is6Even := anonymousFunc(6)

	fmt.Println("anonymousFunc5:", is5Even())
	fmt.Println("anonymousFunc6:", is6Even())

	//go to statement
	goto myLabel

myLabel:
	fmt.Println("A called!")

}

//varidiac function

func test(firstName, lastName string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println("NUMBER:", number)
	}
}

func myFunction1(n int) int {
	return n + n
}

func myFunction2(n int, f func(int) int) int {
	return f(n)
}

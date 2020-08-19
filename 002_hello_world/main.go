package main

import "fmt"

func main() {
	fmt.Println("Hi my name is Alejandro, and if you like follow me in github, my user is AlejandroCarballo")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}

	bar()
}

func foo() {
	fmt.Println("Hi, I am in foo")
}
func bar() {
	fmt.Println("and then we exited")
}

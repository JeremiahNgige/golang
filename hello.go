package main

import "fmt"

var y = 2

//declare there's a variable y and is of type int

func main() {
	fmt.Println(" This is my first ever go program")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i * y)
		}
	}

	fmt.Println("THis is awesome")
}

func foo() {
	fmt.Println("THis is in Foo")
}

package main

import "fmt"

func main() {
	fmt.Println(" This is my first ever go program")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("THis is awesome")
}

func foo() {
	fmt.Println("THis is in Foo")
}

package main

import "fmt"

var a int

type hotdog int

var b hotdog

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

	// creating a type
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}

func foo() {
	fmt.Println("THis is in Foo")
}

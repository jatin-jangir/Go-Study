package main

import "fmt"

func main() {
	var name string = "jatin"
	var p = &name
	fmt.Println(name)
	fmt.Println(*p)

	var b = [5]int{1, 20, 30, 40, 50}
	var i = &b
	fmt.Println(*(i))
}

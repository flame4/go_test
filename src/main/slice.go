package main

import "fmt"

func main() {
	var b []int
	fmt.Printf("%+v", b)
	fmt.Printf("%#v", b)
	fmt.Println(b == nil)
	fmt.Println(len(b))
	a := []int{}
	fmt.Printf("%+v", a)
	fmt.Printf("%#v", a)
	fmt.Println(a == nil)
	fmt.Println(len(a), cap(a))

	var c string
	fmt.Println(c == "")
}

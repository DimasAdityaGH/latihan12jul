package main

import "fmt"

func main() {
	var o = 10
	o++
	fmt.Println(o)

	var x = 10
	x += 10
	fmt.Println(x)

	var i = 10
	i++
	i++
	i += 8
	fmt.Println(i)
}
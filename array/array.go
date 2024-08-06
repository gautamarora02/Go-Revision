package main

import "fmt"

func main() {
	var name = [4]string{"Gautam", "Arora", "Gaurav"}

	var price = [3]int{3, 2, 3}

	fmt.Println(price)
	fmt.Printf("%q", name) // q stands for qutation

	var check [4]int
	check[0] = 1 // access and update value

	fmt.Println(len(check)) // length function

	fmt.Println(check) // initial values check
}

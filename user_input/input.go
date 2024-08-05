package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	// fmt.Scan(&name)
	// it will only read one word
	/*
		For complete line to read, we would need buffer input output
		to access that, we have to initialize reader
		reader = bufio.Newreader(os.Stdio) ...
	*/
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n') // it will read until new line

	fmt.Println(name)
}

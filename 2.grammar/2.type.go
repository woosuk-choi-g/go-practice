package main

import "fmt"

func main() {
	// string escape
	str1 := "AA"
	str2 := "B\nB"

	fmt.Println(str1)
	fmt.Println(str2)

	// raw string literal
	str3 := `AA`
	str4 := `B\nB`

	fmt.Println(str3)
	fmt.Println(str4)
}
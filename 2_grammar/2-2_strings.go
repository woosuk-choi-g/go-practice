package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, ":"))  // a:b:c

	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1)
	fmt.Println(r)  // a_b_c
}
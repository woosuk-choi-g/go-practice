package main

import "fmt"

func main() {
	for index := 1; index <= 10; index++ {
		fmt.Printf("index: %d\n", index)
	}

	for index := 1; index <= 10; {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index := 1
	for ; index <= 10; {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index = 1
	for true {
		fmt.Printf("index: %d\n", index)
		index++

		if index > 10 {
			break
		}
	}

	strArray := []string{"A", "B", "C"}
	for index, str := range strArray {
		fmt.Printf("index: %d, str: %s\n", index, str)
	}

	dictionary := map[string]string{
		"key_A": "value_A",
		"key_B": "value_B",
	}

	for key, value := range dictionary {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
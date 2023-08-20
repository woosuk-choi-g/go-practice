package main

import "fmt"

func main() {
	fmt.Println("---- map ----")
	intMap := map[int]string{}
	intMap[100] = "A"
	intMap[101] = "B"
	fmt.Printf("%+v\n", intMap) // map[100:A 101:B]
	fmt.Println(intMap[100]) // A

	fmt.Println("---- map by make ----")
	var stringMap = make(map[string]string)
	stringMap["A"] = "Hello"
	stringMap["B"] = "World"
	fmt.Printf("%+v\n", intMap) // map[100:A 101:B]
	fmt.Println(stringMap["B"]) // World

	fmt.Println("---- map by literal ----")
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB": "FaceBook",
	}

	fmt.Printf("%+v\n", tickers)
	fmt.Println(tickers["FB"])

	fmt.Println("---- map exists ----")
	val, exists := tickers["FB"]
	fmt.Println(val, exists) // FaceBook true

	val, exists = tickers["NOT_EXISTS"]
	fmt.Println(val, exists) // false

	fmt.Println("---- map iterate ----")
	for key, value := range tickers {
		fmt.Println(key, value)
	}
}
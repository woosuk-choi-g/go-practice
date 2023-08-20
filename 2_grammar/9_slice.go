package main

import "fmt"

func main() {
	// 슬라이스 생성
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters) // [a b c d]

	// make 를 이용한 slice 생성
	s1 := make([]int, 0)
	fmt.Println(s1, len(s1), cap(s1)) // [] 0 0

	s2 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s2)) // [] 0 5

	s3 := make([]int, 3, 5)
	fmt.Println(s2, len(s2), cap(s3)) // [0 0 0 0 0] 5 5

	fmt.Println("---- append")
	intArray := []int{100, 101, 102}
	s1 = append(s1, 100)
	s2 = append(s2, intArray...)
	s3 = append(s3, 100, 101, 102)
	fmt.Println(s1, len(s1), cap(s1)) // [100] 1 1
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 100 101 102] 8 10
	fmt.Println(s3, len(s3), cap(s3)) // [0 0 0 100 101 102] 6 10

	fmt.Println("---- copy")
	lettersCopy1 := make([]string, 0)
	copy(lettersCopy1, letters)
	fmt.Println(lettersCopy1) // []

	lettersCopy2 := make([]string, len(letters), len(letters))
	copy(lettersCopy2, letters)
	fmt.Println(lettersCopy2) // [a b c d]

	lettersCopy2[3] = "="
	fmt.Println(letters) // [a b c d]
	fmt.Println(lettersCopy2) // [a b c =]

	fmt.Println("---- slice")
	integers := []int{1, 2, 3, 4, 5}
	sub1 := integers[1:4]
	sub2 := integers[2:4]
	fmt.Println(integers, len(integers), cap(integers)) // [1 2 3 4 5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1)) // [2 3 4] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2)) // [3 4] 2 3

	sub1[2] = 100
	fmt.Println(integers, len(integers), cap(integers)) // [1 2 100 4 5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1)) // [2 3 100] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2)) // [3 100] 2 3

	fmt.Println("---- slice2")
	testSlice()
}

func testSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:0]
	s3 := s1[0:1]
	s4 := s1[0:2]
	s5 := s1[0:3]

	fmt.Println(s1, len(s1), cap(s1)) // [1 2 3 4 5] 5 5
	fmt.Println(s2, len(s2), cap(s2)) // [] 0 5
	fmt.Println(s3, len(s3), cap(s3)) // [1] 1 5
	fmt.Println(s4, len(s4), cap(s4)) // [1 2] 2 5
	fmt.Println(s5, len(s5), cap(s5)) // [1 2 3] 3 5

	ss2 := s1[1:1]
	ss3 := s1[1:2]
	ss4 := s1[1:3]
	ss5 := s1[1:4]

	fmt.Println(ss2, len(ss2), cap(ss2)) // [] 0 4
	fmt.Println(ss3, len(ss3), cap(ss3)) // [2] 1 4
	fmt.Println(ss4, len(ss4), cap(ss4)) // [2 3] 2 4
	fmt.Println(ss5, len(ss5), cap(ss5)) // [2 3 4] 3 4

	ss2 = append(ss2, 100)
	fmt.Println(s1, len(s1), cap(s1)) // [1 100 3 4 5] 5 5

	fmt.Println(ss2, len(ss2), cap(ss2)) // [100] 1 4
	fmt.Println(ss3, len(ss3), cap(ss3)) // [100] 1 4
	fmt.Println(ss4, len(ss4), cap(ss4)) // [100 3] 2 4
	fmt.Println(ss5, len(ss5), cap(ss5)) // [100 3 4] 3 4
}
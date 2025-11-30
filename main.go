package main

import "fmt"

func MathChangePlace(first *int, second *int) {
	*second = *first + *second
	*first = *second - *first
	*second = *second - *first
}

func XoRChangePlace(first *int, second *int) {
	*first = *first ^ *second
	*second = *first ^ *second
	*first = *first ^ *second
}
func main() {

	first := 2
	second := 3
	fmt.Printf("first:%d second:%d\n", first, second)
	MathChangePlace(&first, &second)
	fmt.Printf("Math change place first:%d second:%d\n", first, second)
	XoRChangePlace(&first, &second)
	fmt.Printf("XoR change place first:%d second:%d\n", first, second)
}

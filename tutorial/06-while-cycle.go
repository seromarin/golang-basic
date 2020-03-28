package main

import "fmt"

// While we live here...
func While() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

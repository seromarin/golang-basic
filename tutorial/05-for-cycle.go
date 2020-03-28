package main

import "fmt"

// For1 will iterate 10 times
func For1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// For2 will iterate 10 times
func For2() {
	sum := 0
	for sum < 100 {
		if sum == 0 {
			sum++
		}
		sum += sum
	}
	fmt.Println(sum)
}

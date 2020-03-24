package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome home")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()
	fmt.Println(operation)

	values := strings.Split(operation, "+")
	fmt.Println(values)

	fmt.Println(values[0] + values[1])
	operator1, err1 := strconv.Atoi(values[0])
	operator2, err2 := strconv.Atoi(values[1])

	fmt.Println(operator1 + operator2)

}

package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operator string) int {

	clearInput := strings.Split(input, operator)

	operator1, _ := parser(clearInput[0])
	operator2, _ := parser(clearInput[1])

	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println("No pereation allowed")
		return 0
	}
}

func parser(input string) (int, error) {
	operator, err := strconv.Atoi(input)
	return operator, err
}

func readUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()
	fmt.Println(operation)
	return operation
}

func main() {

	userInput := readUserInput()
	userOperator := readUserInput()

	fmt.Println(userInput, userOperator)

	c := calc{}
	fmt.Println(c.operate(userInput, userOperator))
}

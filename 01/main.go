package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b float64
	var operand byte

	fmt.Println("Enter the left operand: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	a, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}

	fmt.Println("Enter the operator: ")
	scanner.Scan()
	text = scanner.Text()
	if len(text) != 1 {
		fmt.Println("Invalid input", text)
		return
	}
	if text[0] != '+' && text[0] != '-' && text[0] != '*' && text[0] != '/' {
		fmt.Println("Invalid input", text)
		return
	}
	operand = text[0]

	fmt.Println("Enter the right operand: ")
	scanner.Scan()
	text = scanner.Text()
	b, err = strconv.ParseFloat(text, 64)
	// fmt.Println(text, a)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}
	fmt.Println("Result:")
	switch operand {
	case '+':
		fmt.Println(a + b)
	case '-':
		fmt.Println(a - b)
	case '*':
		fmt.Println(a * b)
	case '/':
		fmt.Printf("%.3f\n", float64(a)/float64(b))
	default:
		fmt.Println("Invalid input", operand)
	}
}

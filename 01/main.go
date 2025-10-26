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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the left operand: ")
	a = readOperand(scanner)

	fmt.Println("Enter the operator: ")
	operand = readOperator(scanner)

	fmt.Println("Enter the right operand: ")
	b = readOperand(scanner)

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

func readOperand(scanner *bufio.Scanner) float64 {
	for {
		scanner.Scan()
		text := scanner.Text()
		a, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return a
	}
}
func readOperator(scanner *bufio.Scanner) byte {
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 1 {
			fmt.Println("Invalid input", text)
			continue
		}
		switch text[0] {
		case '+', '-', '*', '/':
			return text[0]

		default:
			fmt.Println("Invalid input", text)
			continue
		}
	}
}

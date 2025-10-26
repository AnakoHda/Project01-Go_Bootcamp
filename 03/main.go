package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var firstMass []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	firstInput := scanner.Text()
	arrFirst := strings.Fields(firstInput)
	for _, v := range arrFirst {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		firstMass = append(firstMass, num)
	}

	scanner.Scan()
	secondInput := scanner.Text()
	arrSecond := strings.Fields(secondInput)
	secomdMap := make(map[int]int)
	for _, v := range arrSecond {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		secomdMap[num] = 1
	}
	var result string
	for i := 0; i < len(firstMass); i++ {
		if secomdMap[firstMass[i]] == 1 {
			secomdMap[firstMass[i]] = 0
			result += strconv.Itoa(firstMass[i]) + " "
		}
	}
	if result == "" {
		fmt.Println("Empty intersection")
	} else {
		fmt.Println(result)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var abb []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	first := scanner.Text()
	arrfirst := strings.Fields(first)
	for _, v := range arrfirst {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		abb = append(abb, num)
	}

	scanner.Scan()
	second := scanner.Text()
	arrsecond := strings.Fields(second)
	daa := make(map[int]int)
	for _, v := range arrsecond {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		daa[num] = 1
	}
	var result string
	for i := 0; i < len(abb); i++ {
		if daa[abb[i]] == 1 {
			daa[abb[i]] = 0
			result += strconv.Itoa(abb[i]) + " "
		}
	}
	if result == "" {
		fmt.Println("Empty intersection")
	} else {
		fmt.Println(result)
	}
}

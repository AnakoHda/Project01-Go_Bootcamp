package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WordFrequency struct {
	Word      string
	Frequency int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Fields(line)

	scanner.Scan()
	kStr := scanner.Text()
	k, err := strconv.Atoi(strings.TrimSpace(kStr))
	if err != nil {
		return
	}
	fmt.Println(PrintWordFrequencies(words, k))
}
func PrintWordFrequencies(words []string, k int) string {
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}
	var w []WordFrequency
	for word, freq := range freqMap {
		w = append(w, WordFrequency{Word: word, Frequency: freq})
	}
	
	sort.Slice(w, func(i, j int) bool {
		if w[i].Frequency == w[j].Frequency {
			return w[i].Word < w[j].Word
		}
		return w[i].Frequency > w[j].Frequency
	})

	if len(w) < k {
		k = len(w)
	}
	var str string
	for i := 0; i < k; i++ {
		str += w[i].Word + " "
	}
	if len(str) == 0 {
		return ""
	}
	return str[:len(str)-1]
}

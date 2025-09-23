package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	text_file, err := os.Open("../algorithm_inputs/text_exerpt.txt")
	if err != nil {
		panic(err)
	}

	word_occurrances := make(map[string]int)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		for _, word := range row {
			word_occurrances[word]++
		}
	}
	for word, count := range word_occurrances {
		fmt.Println("\n\n\n\nword: ", word, "   count: ", count)
	}
}

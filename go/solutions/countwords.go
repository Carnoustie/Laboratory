//count words in a text file

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Open file
	text_file, err := os.Open("../algorithm_inputs/text_exerpt.txt")

	//bail out if fileopen fails
	if err != nil {
		panic(err)
	}

	var count int //Initializes to zero

	text_scanner := bufio.NewScanner(text_file) //scanner variable

	var temp_text string      // my c++ brain keeps declaring variables before using them....
	for text_scanner.Scan() { //for-loop with go-typical boolean pattern
		temp_text = text_scanner.Text()
		text_slice := strings.Split(temp_text, " ")

		count += len(text_slice)
	}

	fmt.Println("\n\n\n\nWordcount: ", count, "\n\n")
}

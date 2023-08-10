package main

import (
	"bad-algorithms/bad"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := []int{1, 5, 3}
	algo := os.Args[1]
	switch strings.ToLower(algo) {
	case "bad":
		fmt.Printf("O(n!) algorithm where n has a length of: %d\ntotal iterations: ", len(input))
		fmt.Println(bad.NFactorial(input))

	case "worse":
		fmt.Printf("O(n^n) algorithm where n has a length of: %d\ntotal iterations: ", len(input))
		fmt.Println(bad.NToTheN(input, len(input)))

	case "worst":
		fmt.Printf("O((n^n)!) algorithm where n has a length of: %d\ntotal iterations: ", len(input))
		nToTheN := make([]int, bad.NToTheN(input, len(input)))
		fmt.Print(bad.NFactorial(nToTheN))
	}

}

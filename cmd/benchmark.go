package cmd

import (
	"bad-algorithms/pkg/bad"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const expectedInputFormat = "expected args: `[bad||worse||worst] [setSize integer] [other golang flags]...`"

func parseArgs(args []string) (string, []int) {
	if len(args) < 2 {
		panic(expectedInputFormat)
	}
	algo := args[1]
	setSize := args[2]
	size, e := strconv.Atoi(setSize)
	if e != nil {
		panic(expectedInputFormat)
	}
	return algo, make([]int, size)
}

func Benchmark(args []string) {
	algo, input := parseArgs(args)

	switch strings.ToLower(algo) {
	case "bad":
		benchBad(input)
	case "worse":
		benchWorse(input)
	case "worst":
		benchWorst(input)
	default:
		fmt.Println(expectedInputFormat)
	}
}

const totalIterationsF = "total iterations: %d"

func benchBad(n []int) {
	defer funcTimer()()
	log.Printf("O(n!) algorithm where n has a size of: %d", len(n))
	log.Printf(totalIterationsF, bad.NFactorial(n))
	return
}

func benchWorse(n []int) {
	l := len(n)
	log.Printf("O(n^n) algorithm where n has a size of: %d", l)
	defer funcTimer()()
	log.Printf(totalIterationsF, bad.NToTheN(n, l))
	return
}

func benchWorst(n []int) {
	l := len(n)
	log.Printf("O((n^n)!) algorithm where n has a size of: %d", l)
	defer funcTimer()()
	log.Printf(totalIterationsF, bad.NtoTheNFactorial(n, l))
}

func funcTimer() func() {
	t := time.Now()

	return func() {
		log.Printf("time elapsed: %v", time.Since(t))
	}
}

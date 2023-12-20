package main

import (
	"math"
	"strings"

	"github.com/maxmac99/adventofcode/shared"
)

var words = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func main() {
	shared.RunTask(solutionMulti)
}

func solutionMulti(lines chan string) (int, error) {
	return ResultFromWords(lines, words), nil
}

func solutionSingle(lines chan string) (int, error) {
	return ResultFromWordsSingle(lines, words), nil
}

func ResultFromWords(lines chan string, words map[string]int) int {
	const numWorkers = 8
	workPool := shared.StartNewWorkPool(numWorkers)

	// Schedule tasks
	go func() {
		for line := range lines {
			copied := line
			workPool.AddTask(func() int {
				return getNumFromLine(copied, words)
			})
		}
		workPool.CompleteTasks()
	}()

	return workPool.Sum()
}

func ResultFromWordsSingle(lines chan string, words map[string]int) int {
	total := 0
	for line := range lines {
		total += getNumFromLine(line, words)
	}
	return total
}

func getNumFromLine(line string, words map[string]int) int {
	values := map[int]int{}
	for word, value := range words {
		result := strings.Index(line, word)
		if result != -1 {
			values[result] = value
		}
	}

	firstChar := values[min(values)]

	values = map[int]int{}
	for word, value := range words {
		result := strings.LastIndex(line, word)
		if result != -1 {
			values[result] = value
		}
	}
	lastChar := values[max(values)]

	return firstChar*10 + lastChar
}

func min(values map[int]int) int {
	min := math.MaxInt
	for k := range values {
		if k < min {
			min = k
		}
	}
	return min
}

func max(values map[int]int) int {
	max := math.MinInt
	for k := range values {
		if k > max {
			max = k
		}
	}
	return max
}

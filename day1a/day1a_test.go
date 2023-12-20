package main

import (
	"github.com/maxmac99/adventofcode/shared"
	"testing"
)

func Test_Day1A(t *testing.T) {
	shared.TestSolution(solutionMulti, 56465, t)
}

func BenchmarkDay1A(b *testing.B) {
	tests := []shared.BenchRun{
		{
			Name:        "Multi",
			Filename:    "input.txt",
			Solution:    solutionMulti,
			BufferSize:  []int{1, 10, 100, 100},
			MaxNumLines: []int{1, 10, 100, 1000},
		},
		{
			Name:        "Single",
			Filename:    "input.txt",
			Solution:    solutionSingle,
			BufferSize:  []int{1, 10, 100, 100},
			MaxNumLines: []int{1, 10, 100, 1000},
		},
	}

	shared.BenchSolutions(tests, b)
}

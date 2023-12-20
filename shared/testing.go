package shared

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type Solution func(chan string) (int, error)

func TestSolution(solution Solution, expected int, t *testing.T) {
	lines := ReadLines("input.txt")

	got, err := solution(lines)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

type BenchRun struct {
	Name        string
	Filename    string
	Solution    Solution
	BufferSize  []int
	MaxNumLines []int
}

func BenchSolutions(tests []BenchRun, b *testing.B) {
	for _, bb := range tests {
		for i, maxNumLines := range bb.MaxNumLines {
			b.Run(bb.Name+"_"+strconv.Itoa(maxNumLines), func(b *testing.B) {
				lines := ReadFixedLines(bb.Filename, maxNumLines)
				for n := 0; n < b.N; n++ {
					outputChan := make(chan string, bb.BufferSize[i])
					go func() {
						for _, line := range lines {
							outputChan <- line
						}
						close(outputChan)
					}()
					bb.Solution(outputChan)
				}
			})
		}
	}
}

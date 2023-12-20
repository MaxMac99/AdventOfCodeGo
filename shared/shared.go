package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunTask(solution func(chan string) (int, error)) {
	lines := ReadLines("input.txt")
	result, err := solution(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", result)
}

func ReadLines(filename string) chan string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := make(chan string, 10)

	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			lines <- line
		}
		close(lines)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return lines
}

func ReadFixedLines(filename string, maxNumLines int) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for i := 0; i < maxNumLines && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	return lines
}

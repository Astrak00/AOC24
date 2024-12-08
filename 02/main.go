package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Result part 1:", part1())
	fmt.Println("Result part 2:", part2())

}

func part2() int {
	acc := 0
	numbers := readFile()
	for _, line := range numbers {

		// Create a line without one of the elements and test it for every element in the slice:
		for i := range len(line) {
			tempLine := make([]int, 0)
			for j := range len(line) {
				if i != j {
					tempLine = append(tempLine, line[j])
				}
			}
			if isSafe2(tempLine) {
				acc++
				break
			}
		}
	}
	return acc
}

func isSafe2(numbers []int) bool {
	decreasing := false
	if numbers[0] > numbers[1] {
		decreasing = true
	}
	if decreasing {
		for i := range len(numbers) - 1 {
			temp := numbers[i] - numbers[i+1]
			if temp > 3 || temp < 1 {
				return false
			}
		}
	} else {
		for i := range len(numbers) - 1 {
			temp := numbers[i+1] - numbers[i]
			if temp > 3 || temp < 1 {
				return false
			}
		}
	}
	return true
}

func part1() int {
	acc := 0
	numbers := readFile()
	for _, line := range numbers {
		if isSafe(line) {
			acc++
		}
	}
	return acc
}

func isSafe(numbers []int) bool {
	decreasing := false
	if numbers[0] > numbers[1] {
		decreasing = true
	}
	if decreasing {
		for i := range len(numbers) - 1 {
			temp := numbers[i] - numbers[i+1]
			if temp > 3 || temp < 1 {
				return false
			}
		}
	} else {
		for i := range len(numbers) - 1 {
			temp := numbers[i+1] - numbers[i]
			if temp > 3 || temp < 1 {
				return false
			}
		}
	}
	return true
}

func readFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		line_number := make([]int, len(line))
		for i, elem := range line {
			elem_int, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatalf("Parsing to int error: %v", err)
			}
			line_number[i] = elem_int
		}
		numbers = append(numbers, line_number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

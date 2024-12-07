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
	lines := readFile("input.txt")
	acc := 0
	for _, line := range lines {
		result := line[0]
		restNumbers := line[1:]
		ops := make([]bool, len(restNumbers)-1)
		operations := permutations(ops)
		for _, operation := range operations {
			if checkOperation(result, restNumbers, operation) {
				acc += result
				break
			}
		}
	}
	fmt.Println(acc)
}

func checkOperation(result int, restNumbers []int, operations []bool) bool {
	temp_acc := restNumbers[0]
	for i, operation := range operations {
		if operation {
			temp_acc += restNumbers[i+1]
		} else {
			temp_acc *= restNumbers[i+1]
		}
	}
	return result == temp_acc

}

func permutations(operations []bool) [][]bool {
	var result [][]bool
	permutationsHelper(operations, 0, &result)
	return result
}

func permutationsHelper(operations []bool, start int, result *[][]bool) {
	if start == len(operations) {
		perm := make([]bool, len(operations))
		copy(perm, operations)
		*result = append(*result, perm)
		return
	}

	operations[start] = false
	permutationsHelper(operations, start+1, result)

	operations[start] = true
	permutationsHelper(operations, start+1, result)
}

func readFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		line[0] = line[0][:len(line[0])-1]
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

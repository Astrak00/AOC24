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
	acc_part1 := 0
	acc_part2 := 0
	for _, line := range lines {
		result := line[0]
		restNumbers := line[1:]
		ops := make([]int8, len(restNumbers)-1)
		for i := range 2 {
			operations := permutations(ops, i+2)
			for _, operation := range operations {
				if checkOperation(result, restNumbers, operation) {
					if i == 0 {
						acc_part1 += result
					} else {
						acc_part2 += result
					}
					break
				}
			}
		}

	}
	fmt.Println("Result for part1: ", acc_part1)
	fmt.Println("Result for part2: ", acc_part2)
}

func checkOperation(result int, restNumbers []int, operations []int8) bool {
	temp_acc := restNumbers[0]
	for i, operation := range operations {
		switch operation {
		case 0:
			temp_acc += restNumbers[i+1]
		case 1:
			temp_acc *= restNumbers[i+1]
		case 2:
			left := strconv.Itoa(temp_acc)
			right := strconv.Itoa(restNumbers[i+1])
			temp_cnversion, err := strconv.Atoi(left + right)
			if err != nil {
				log.Fatal("Error with conversion ", err)
			}
			temp_acc = temp_cnversion
		}
	}
	return result == temp_acc
}

func permutations(ops []int8, variants int) [][]int8 {
	var helper func([]int8, int)
	res := [][]int8{}

	helper = func(arr []int8, n int) {
		if n == len(arr) {
			tmp := make([]int8, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := int8(0); i < int8(variants); i++ {
			arr[n] = i
			helper(arr, n+1)
		}
	}

	helper(ops, 0)
	return res
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

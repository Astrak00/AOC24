package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Result part 1:")
	part1()
	fmt.Println("Result part 2:")
	part2()
}

func part2() {
	left, right := readFile()

	// Creating a map to store the appearance of the number of the second array(right)
	map_right := make(map[int]int)

	// We creare a frequency register for each ocurrence in the right column
	for _, num := range right {
		map_right[num]++
	}

	simmilarity := 0

	for i := 0; i < len(left); i++ {
		simmilarity += left[i] * map_right[left[i]]
	}
	fmt.Println(simmilarity)

}

func part1() {
	left, right := readFile()

	// Sorting the arrays
	left = quickSort(left)
	right = quickSort(right)

	difference := 0

	for i := 0; i < len(left); i++ {
		difference += abs(left[i] - right[i])
	}

	fmt.Println(difference)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr)
	pivot := arr[n/2]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < n; i++ {
		if i == n/2 {
			continue
		}

		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func readFile() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%d   %d", &left[i], &right[i])
	}

	return left, right
}

package main

import (
	"fmt"
	"os"
)

var directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var word = "XMAS"

func main() {
	grid, err := ReadGrid("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	result1, result2 := countWord(grid, word)
	fmt.Println("The result for the part 1 is: ", result1)
	fmt.Println("The result for the part 2 is: ", result2)

}

func countWord(grid [][]rune, word string) (int, int) {
	wordLen := len(word)
	wordChars := []rune(word)
	rows := len(grid)
	columns := len(grid[0])

	count1 := 0
	count2 := 0

	// loop though each character looking for the word as this prevents us from just "looking" for one word and need to substitute by another option
	for i := range rows {
		for j := range columns {
			for _, dir := range directions {
				matchFound := true
				for k := range wordLen {
					ni := i + k*dir[0]
					nj := j + k*dir[1]

					if ni < 0 || nj < 0 || ni >= rows || nj >= columns {
						matchFound = false
						break
					}

					if grid[ni][nj] != wordChars[k] {
						matchFound = false
						break
					}
				}
				if matchFound {
					count1++
				}
			}
			if crossMas(grid, i, j) {
				count2++
			}
		}
	}
	return count1, count2
}

func crossMas(grid [][]rune, i, j int) bool {
	if grid[i][j] != 'A' {
		return false
	}

	rows := len(grid)
	columns := len(grid[0])

	topLeft := [2]int{i - 1, j - 1}
	topRight := [2]int{i - 1, j + 1}
	bottomLeft := [2]int{i + 1, j - 1}
	bottomRight := [2]int{i + 1, j + 1}
	positions := [][2]int{topLeft, topRight, bottomLeft, bottomRight}
	for _, posCheck := range positions {
		if posCheck[0] < 0 || posCheck[0] >= rows || posCheck[1] < 0 || posCheck[1] >= columns {
			return false
		}
	}
	if (grid[topLeft[0]][topLeft[1]] == 'M' && grid[bottomRight[0]][bottomRight[1]] == 'S') ||
		(grid[topLeft[0]][topLeft[1]] == 'S' && grid[bottomRight[0]][bottomRight[1]] == 'M') {
		if (grid[topRight[0]][topRight[1]] == 'M' && grid[bottomLeft[0]][bottomLeft[1]] == 'S') ||
			(grid[topRight[0]][topRight[1]] == 'S' && grid[bottomLeft[0]][bottomLeft[1]] == 'M') {
			return true
		}

	}

	return false
}

// Function ReadGrid from my own word-searcher (https://github.com/Astrak00/GoAlphaSoup/blob/main/main.go#L145)
func ReadGrid(filename string) ([][]rune, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the grid
	var grid [][]rune
	for {
		var row []rune
		for {
			var cell rune
			_, err := fmt.Fscanf(file, "%c", &cell)
			if err != nil {
				break
			}
			if cell == '\n' {
				break
			} else {
				row = append(row, cell)
			}
		}
		if len(row) == 0 {
			break
		}
		grid = append(grid, row)
	}
	return grid, nil
}

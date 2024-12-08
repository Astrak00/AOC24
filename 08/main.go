package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Mapa [][]rune
type Position [2]int
type Estaciones map[rune][]Position

func main() {
	content := readFile("input.txt")
	//solution := readFile("solution.txt")
	rows := len(content)
	cols := len(content[0])

	estaciones := content.findStations()
	//sestaciones.printEstaciones()

	candidates := map[[2]int]bool{}
	// _, ok := s[6] // check for existence
	// s[8] = true   // add element
	// delete(s, 2)  // remove element

	for _, emisoras := range estaciones {
		for i, pos1 := range emisoras {
			for _, pos2 := range emisoras[i+1:] {
				distanceTop := calculateDistanceVector(pos1, pos2)
				// Top: 0,3 - Bottom: 1,0 = 0-1, 3-0 = Top += -1,3
				// Inviertes -1,3 -> 1,-3
				// Bottom += 1, -3
				pos1Tmp := pos1.addDistance(distanceTop)
				if pos1Tmp.checkBounds(rows, cols) {
					candidates[pos1Tmp] = true
				}
				distanceTop = distanceTop.invertDistance()
				pos2Tmp := pos2.addDistance(distanceTop)
				if pos2Tmp.checkBounds(rows, cols) {
					candidates[pos2Tmp] = true
				}

			}
		}
	}
	fmt.Println("Solution part 1:", len(candidates))
}

func (p Position) checkBounds(row, columns int) bool {
	if p[0] >= 0 && p[0] < row && p[1] >= 0 && p[1] < columns {
		return true
	}
	return false
}

func (p Position) invertDistance() Position {
	return Position{-p[0], -p[1]}
}

func (p Position) addDistance(addPos Position) Position {
	return Position{p[0] + addPos[0], p[1] + addPos[1]}
}

func calculateDistanceVector(pos1, pos2 Position) Position {
	x := pos1[0] - pos2[0]
	y := pos1[1] - pos2[1]
	return Position{x, y}
}

func (m Mapa) findStations() Estaciones {
	estaciones := make(Estaciones, 0)
	for i, row := range m {
		for j, elem := range row {
			if elem != '.' {
				estaciones[elem] = append(estaciones[elem], [2]int{i, j})
			}
		}
	}
	return estaciones
}

func (e Estaciones) printEstaciones() {
	for k, v := range e {
		fmt.Println(string(k), "Values:", v)
	}
}

func readFile(path string) Mapa {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fileContent := make(Mapa, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line_elems := make([]rune, len(line))
		for i, elem := range line {
			line_elems[i] = elem
		}
		fileContent = append(fileContent, line_elems)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileContent
}

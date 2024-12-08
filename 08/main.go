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

	rows := len(content)
	cols := len(content[0])

	estaciones := content.findStations()
	//sestaciones.printEstaciones()

	candidatesPart1 := map[Position]bool{}
	candidatesPart2 := map[Position]bool{}

	// We add the positions of the stations as we know they will be in-line and will form an antinode.
	for _, val := range estaciones {
		for _, station := range val {
			candidatesPart2[station] = true
		}
	}

	for _, emisoras := range estaciones {
		for i, pos1 := range emisoras {
			for _, pos2 := range emisoras[i+1:] {
				distanceTop := calculateDistanceVector(pos1, pos2)

				pos1Tmp := pos1.addDistance(distanceTop)
				if pos1Tmp.checkBounds(rows, cols) {
					candidatesPart1[pos1Tmp] = true
				}
				// While we are in-bounds, we continue the "line"
				for pos1Tmp.checkBounds(rows, cols) {
					candidatesPart2[pos1Tmp] = true
					pos1Tmp = pos1Tmp.addDistance(distanceTop)
				}

				// Inverting the direction of the displacement vector to form the oposite side of the line
				distanceTop = distanceTop.invertDistance()
				pos2Tmp := pos2.addDistance(distanceTop)
				if pos2Tmp.checkBounds(rows, cols) {
					candidatesPart1[pos2Tmp] = true
				}
				for pos2Tmp.checkBounds(rows, cols) {
					candidatesPart2[pos2Tmp] = true
					pos2Tmp = pos2Tmp.addDistance(distanceTop)
				}

			}
		}
	}
	fmt.Println("Solution part 1:", len(candidatesPart1))
	fmt.Println("Solution part 2:", len(candidatesPart2))

	//for k := range candidatesPart2 {
	//	content[k[0]][k[1]] = '#'
	//}
	//content.printMap()
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
				estaciones[elem] = append(estaciones[elem], Position{i, j})
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

func (m Mapa) printMap() {
	for _, row := range m {
		for _, elem := range row {
			fmt.Print(string(elem))
		}
		fmt.Print("\n")
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

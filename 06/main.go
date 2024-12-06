package main

import (
	"fmt"
	"log"
	"os"
)

type Map [][]rune
type Direction map[Facing]Position
type Position [2]int
type Facing string

var errorPos = Position{-1, -1}
var directions = Direction{"north": Position{-1, 0}, "south": Position{1, 0}, "east": Position{0, 1}, "west": Position{0, -1}}

func main() {
	mapa, err := readMap("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	position := mapa.searchStartingPos()
	if position == errorPos {
		log.Fatal("Position can not be found in initial Map")
	}
	var direction Facing = "north"

outerLoop:
	for {
		nextPosElem := mapa.getNextPositionElem(position, direction)
		switch nextPosElem {
		case '0':
			mapa.setVisited(position)
			break outerLoop
		case '#':
			direction = direction.turnRight()
		default:
			position = mapa.move(position, direction)
			if position == errorPos {
				log.Fatal("Invalid position")
			}
		}
	}
	mapa.printMap()
	fmt.Println(position)

	numberVisited := mapa.visited()
	fmt.Println("The number of disitinct positions the guard will go to is: ", numberVisited)

}

func (f Facing) turnRight() Facing {
	switch f {
	case "north":
		return "east"
	case "east":
		return "south"
	case "south":
		return "west"
	case "west":
		return "north"
	default:
		return "nada"
	}
}

func (m Map) visited() int {
	acc := 0
	for _, row := range m {
		for _, elem := range row {
			if elem == 'X' {
				acc += 1
			}
		}
	}
	return acc
}

func (m Map) move(position Position, direction Facing) Position {
	m.setVisited(position)
	nextX := position[0] + directions[direction][0]
	nextY := position[1] + directions[direction][1]
	if nextX < len(m) && nextY < len(m[0]) {
		return Position{nextX, nextY}
	}
	return errorPos
}

func (m Map) setVisited(position Position) {
	m[position[0]][position[1]] = 'X'
}

func (m Map) getNextPositionElem(position Position, direction Facing) rune {
	nextX := position[0] + directions[direction][0]
	nextY := position[1] + directions[direction][1]
	if nextX < len(m) && nextY < len(m[0]) {
		return m[nextX][nextY]
	}
	return '0'
}

func (m Map) searchStartingPos() Position {
	for i, row := range m {
		for j, elem := range row {
			if elem == '^' {
				return Position{i, j}
			}
		}
	}
	return errorPos
}

func (m Map) printMap() {
	for _, row := range m {
		for _, elem := range row {
			fmt.Print(string(elem))
		}
		fmt.Println()
	}
}

func readMap(filename string) (Map, error) {
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

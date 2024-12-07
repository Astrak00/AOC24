package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

type Map [][]rune
type Direction map[Facing]Position
type Position [2]int
type Facing string

var errorPos = Position{-1, -1}
var directions = Direction{"north": Position{-1, 0}, "south": Position{1, 0}, "east": Position{0, 1}, "west": Position{0, -1}}

func main() {

	mapa, startingPosition := mapInitialization()
	var direction Facing = "north"
	position := startingPosition

	position = navigateMap(mapa, position, direction)
	//mapa.printMap()
	fmt.Println(position)

	visited := mapa.visited()
	fmt.Println("The number of disitinct positions the guard will go to is:", len(visited))

	// newTempMap, pos2 := mapInitialization()
	// newTempMap = newTempMap.copyNewObstacle(Position{0, 0})
	// newTempMap.printMap()
	// var direction2 Facing = "north"
	// pos2 = navigateMap(newTempMap, pos2, direction2)
	// fmt.Println(pos2)

	// Part2
	// The idea for part two is creating a thread for each of the posible obstacles.
	// If after some seconds the prorgam has not finished, that means it has gone into a loop
	// If it finishes, it does not create a loop.

	/* numberPosibleObstacles := len(visited)
	fmt.Println(numberPosibleObstacles - 1)

	for _, pos := range visited {
		fmt.Println("Analizando posición:", pos)
		baseMap, startingPosition := mapInitialization()
		if startingPosition != pos {
			baseMap[pos[0]][pos[1]] = '#'
			baseMap.printMap()
			var direction2 Facing = "north"
			positionexit := navigateMap(baseMap, startingPosition, direction2)
			fmt.Println(positionexit, position)
		}
	}
	*/

	otherBaseMap, startingPosition := mapInitialization()

	var wg sync.WaitGroup

	for _, pos := range visited {
		if startingPosition != pos {
			wg.Add(1)
			go func(pos Position, mapaOriginal Map) {
				defer wg.Done()
				mapaOriginal[pos[0]][pos[1]] = '#'
				var direction2 Facing = "north"
				_ = navigateMap(mapaOriginal, startingPosition, direction2)
				//fmt.Println(positionexit, position)
			}(pos, otherBaseMap.copy())
		}

	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Number of object to create a guard-loop: ", runtime.NumGoroutine()-1)
	}
}

func navigateMap(mapaLocal Map, position Position, direction Facing) Position {
	for {
		nextPosElem := mapaLocal.getNextPositionElem(position, direction)
		switch nextPosElem {
		case '0':
			mapaLocal.setVisited(position)
			return position
		case '#':
			direction = direction.turnRight()
		default:
			position = mapaLocal.move(position, direction)
			if position == errorPos {
				log.Fatal("Invalid position")
			}
		}
	}
}

func mapInitialization() (Map, Position) {
	mapa, err := readMap("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	startingPosition := mapa.searchStartingPos()
	if startingPosition == errorPos {
		log.Fatal("Position can not be found in initial Map")
	}
	return mapa, startingPosition
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

func (m Map) copy() Map {
	var newMap Map
	for _, row := range m {
		rowTemp := make([]rune, len(row))
		copy(rowTemp, row)
		newMap = append(newMap, rowTemp)
	}
	return newMap
}

func (m Map) copyNewObstacle(pos Position) Map {
	var newMap Map
	for i, row := range m {
		rowTemp := make([]rune, len(row))
		copy(rowTemp, row)
		if i == pos[0] {
			rowTemp[pos[1]] = '#'
		}
		newMap = append(newMap, rowTemp)
	}
	return newMap
}

func (m Map) visited() []Position {
	acc := make([]Position, 0)
	for i, row := range m {
		for j, elem := range row {
			if elem != '#' && elem != '.' {
				acc = append(acc, Position{i, j})
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
	if nextX < len(m) && nextY < len(m[0]) && nextX >= 0 && nextY >= 0 {
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

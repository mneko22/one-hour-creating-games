package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const fieldWidth int = 120
const fieldHeight int = 120
const fps = 1
const interval = 1000 / fps

var field [fieldHeight][fieldWidth]bool

func patternTransfer(
	_destX, _destY int,
	_srcWidth, _srcHeight int,
	_pPattern [][]bool,
) {
	for y, row := range _pPattern {
		for x, c := range row {
			field[_destY+y][_destX+x] = c
		}
	}
}

func drawField() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	bf := ""
	for _, y := range field {
		for _, x := range y {
			if x == true {
				bf = bf + "■"
			} else {
				bf = bf + " "
			}
		}
		bf = bf + "\n"
	}
	fmt.Print(bf)
}

func getLivingCellsCount(_x int, _y int) int {
	count := 0
	for y := _y - 1; y <= _y+1; y++ {
		repodY := (fieldHeight + y) % fieldHeight
		for x := _x - 1; x <= _x+1; x++ {
			repodX := (fieldWidth + x) % fieldWidth
			if (repodY == _y) && (repodX == _x) {
				continue
			}
			if field[repodY][repodX] == true {
				count++
			}
		}
	}
	return count
}

func stepSimulation() {
	nextField := [fieldHeight][fieldWidth]bool{}
	for y := 0; y < fieldHeight; y++ {
		for x := 0; x < fieldWidth; x++ {
			switch livingCellCount := getLivingCellsCount(x, y); livingCellCount {
			case 2:
				nextField[y][x] = field[y][x]
			case 3:
				nextField[y][x] = true
			default:
				nextField[y][x] = false
			}
		}
	}
	field = nextField
}

func main() {
	const patternWidth = 10
	const patternHeight = 8

	pattern := [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, true, false, true, true, false},
		{false, false, false, false, false, true, false, true, false, false},
		{false, false, false, false, false, true, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, true, false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	patternTransfer(fieldWidth/2, fieldHeight/2, patternWidth, patternHeight, pattern)

	for {
		drawField()
		stepSimulation()
		time.Sleep(interval * time.Millisecond)
	}
}

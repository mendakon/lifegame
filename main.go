package main

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const CELLS_ROW = 30
const CELLS_COLUMN = 30
const PARSENT = 75

func main() {
	cells := initCells()
	for {
		printCells(cells)
		cells = calcCells(cells)
		time.Sleep(time.Millisecond * 10)
	}
}

func calcCells(cells [][]bool) [][]bool {
	newCells := make([][]bool, CELLS_ROW)
	for row := 0; row < CELLS_ROW; row++ {
		newCells[row] = make([]bool, CELLS_COLUMN)
		copy(newCells[row], cells[row])
	}
	for row := 0; row < CELLS_ROW; row++ {
		for col := 0; col < CELLS_COLUMN; col++ {
			rowMinusOver := row-1 < 0
			rowPlusOver := row+1 >= CELLS_ROW
			colMinusOver := col-1 < 0
			colPlusOver := col+1 >= CELLS_COLUMN

			cellCount := 0

			if !rowMinusOver {
				if !colMinusOver {
					if cells[row-1][col-1] {
						cellCount++
					}
				}
				if cells[row-1][col] {
					cellCount++
				}
				if !colPlusOver {
					if cells[row-1][col+1] {
						cellCount++
					}
				}
			}
			if !rowPlusOver {
				if !colMinusOver {
					if cells[row+1][col-1] {
						cellCount++
					}
				}
				if cells[row+1][col] {
					cellCount++
				}
				if !colPlusOver {
					if cells[row+1][col+1] {
						cellCount++
					}
				}
			}
			if !colMinusOver {
				if cells[row][col-1] {
					cellCount++
				}
			}
			if !colPlusOver {
				if cells[row][col+1] {
					cellCount++
				}
			}

			if cellCount <= 1 {
				newCells[row][col] = false
			} else if cellCount == 3 {
				newCells[row][col] = true
			} else if cellCount > 3 {
				newCells[row][col] = false
			}

		}
	}
	return newCells
}

func initCells() [][]bool {
	cells := make([][]bool, CELLS_ROW)

	for row := 0; row < CELLS_ROW; row++ {
		cells[row] = make([]bool, CELLS_COLUMN)
	}

	for row := 0; row < CELLS_ROW; row++ {
		for col := 0; col < CELLS_COLUMN; col++ {
			randInt := rand.Intn(2)
			if randInt == 1 {
				cells[row][col] = true
			}
		}
	}

	return cells
}

func printCells(cells [][]bool) {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	content := ""
	for row := 0; row < CELLS_ROW; row++ {
		for col := 0; col < CELLS_COLUMN; col++ {
			if cells[row][col] {
				content += "■"
			} else {
				content += "　"
			}
		}
		content += "\n"
	}
	print(content)
}

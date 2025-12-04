package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid struct {
	Rows int
	Cols int
	Data [][]int
}

func count(grid [][]int, row, col int) int {
	sum := 0

	offsets := []struct{ dr, dc int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	rows := len(grid)
	if rows == 0 {
		return sum
	}
	cols := len(grid[0])

	for _, offset := range offsets {
		r, c := row+offset.dr, col+offset.dc

		if r >= 0 && r < rows && c >= 0 && c < cols {
			sum += grid[r][c]
		}
	}

	return sum
}
func main() {
	var myGrid Grid
	rolls := 0
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	rowCount := 0
	var g [][]int
	for fileScanner.Scan() {
		row := strings.Split(fileScanner.Text(), "")
		newRow := make([]int, len(row))
		myGrid.Cols = len(row)
		for i, c := range row {
			if c == "@" {
				newRow[i] = 1
			}
		}
		g = append(g, newRow)
		rowCount++

	}
	myGrid.Rows = rowCount
	myGrid.Data = g
	for r := 0; r < myGrid.Rows; r++ {

		for c := 0; c < myGrid.Cols; c++ {
			if myGrid.Data[r][c] == 1 {
				if (count(myGrid.Data, r, c)) < 4 {
					rolls++

				}
			}
		}
	}
	removed := 0
	newGrid := myGrid
	for r := 0; r < myGrid.Rows; r++ {

		for c := 0; c < myGrid.Cols; c++ {
			if myGrid.Data[r][c] == 1 {
				if (count(myGrid.Data, r, c)) < 4 {
					newGrid.Data[r][c] = 0
					removed++
				}
			}
		}
	}
	prev := 0
	fmt.Println(rolls)
	for removed > prev {
		myGrid = newGrid

		prev = removed
		for r := 0; r < myGrid.Rows; r++ {

			for c := 0; c < myGrid.Cols; c++ {
				if myGrid.Data[r][c] == 1 {
					if (count(myGrid.Data, r, c)) < 4 {
						newGrid.Data[r][c] = 0
						removed++
					}
				}
			}
		}
	}
	fmt.Println(removed)
}

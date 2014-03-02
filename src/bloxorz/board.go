package bloxorz

/*
This file creates the board from file and sets the
attributes of each cell.
*/

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Cell struct {
	Type  string // human readable description
	PosX  int
	PosY  int
	Solid bool
}

/*
Reads in file and parses it.
  0 is empty space
  1 is solid space
  8 is the block starting position, assumed vertical
  9 is the goal space
The board is surrounded by 0s for edge detection
*/
func ReadBoardFromFile(filename string) ([]Cell, Bloxor) {
	board := make([]Cell, 0)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic("unable to find board")
	}
	// coordinates
	x, y := 0, 0
	theBlock := Bloxor{}
	for _, char := range contents {
		// if we hit a newline, we move down the file but up coordinate grid
		if string(char) == "\n" {
			// adjust coordinates
			y++
			x = 0
			continue
		}
		// parse this character on the board
		cell := makeCell(char, x, y)
		if cell.Type == "start" {
			a := Block{PosX: x, PosY: y}
			b := Block{PosX: x, PosY: y}
			theBlock.A = a
			theBlock.B = b
		}
		board = append(board, cell)
		x++ // adjust coordinates for next read
	}
	return board, theBlock
}

// depending on the character read, set attributes accordingly
func makeCell(char byte, x, y int) Cell {
	// all cell definitions are integers. This limits us to 10 cell types.
	cellType, err := strconv.Atoi(string(char))

	//if err != nil && string(char) != "x" { // originally, x was block start
	if err != nil {
		fmt.Println(err)
		panic("can't determine cell type; Must be an integer or x")
	}

	// defaults
	solid := false
	cellDesc := "empty"

	if cellType >= 1 {
		solid = true
		cellDesc = "solid"
	}
	if cellType == 9 {
		cellDesc = "goal"
	}
	if cellType == 8 {
		solid = true
		cellDesc = "start"
	}
	return Cell{cellDesc, x, y, solid}
}

// loops through []Cell and prints accordingly
func PrintBoard(Board []Cell, theBlock Bloxor) {
	char := "x"
	X, Y := 0, 0
	for _, c := range Board {
		if c.PosY > Y {
			// adjust coordinates
			fmt.Printf("\n")
			Y++
			X = 0
		}
		if c.Solid {
			char = "▩ "
		} else {
			//char = " ▢ "
			char = "  "
		}
		if c.Type == "goal" {
			char = "\033[1;32m◉\033[0m "
		}
		if X == theBlock.A.PosX && Y == theBlock.A.PosY {
			char = "\033[1;31m▣\033[0m "
		} else if X == theBlock.B.PosX && Y == theBlock.B.PosY {
			char = "\033[1;31m▣\033[0m "
		}
		fmt.Printf(char)
		X++ // adjust coordinates for next iteration
	}
	fmt.Printf("\n\n")
}

// checked to see if the game should end
func Status(Board []Cell, theBlock Bloxor) (win, dead bool) {
	// defaults
	dead = false
	win = false
	X, Y := 0, 0
	for _, c := range Board {
		if c.PosY > Y {
			// adjust the coordinates
			Y++
			X = 0
		}
		if !c.Solid {
			if X == theBlock.A.PosX && Y == theBlock.A.PosY {
				dead = true
			}
			if X == theBlock.B.PosX && Y == theBlock.B.PosY {
				dead = true
			}
		}
		if c.Type == "goal" && theBlock.isStacked() {
			if X == theBlock.A.PosX && Y == theBlock.A.PosY {
				win = true
			}
		}

		// give some feedback
		if dead {
			fmt.Println("You have died")
			return
		}
		if win {
			fmt.Println("You have won")
			return
		}
		X++
	}
	return
}

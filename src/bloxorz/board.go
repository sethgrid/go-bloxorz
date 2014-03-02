package bloxorz

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Cell struct {
	Type  string
	PosX  int
	PosY  int
	Solid bool
}

func ReadMap(filename string) ([]Cell, Bloxor) {
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
		if string(char) == "\n" {
			y++
			x = 0
			continue
		}
		cell := makeCell(char, x, y)
		if cell.Type == "start" {
			a := Block{PosX: x, PosY: y}
			b := Block{PosX: x, PosY: y}
			theBlock.A = a
			theBlock.B = b
		}
		board = append(board, cell)
		x++
	}
	return board, theBlock
}

func makeCell(char byte, x, y int) Cell {
	cellType, err := strconv.Atoi(string(char))
	if err != nil && string(char) != "x" {
		fmt.Println(err)
		panic("can't determine cell type; Must be an integer or x")
	}
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

func PrintBoard(Board []Cell, theBlock Bloxor) {
	char := "x"
	X, Y := 0, 0
	for _, c := range Board {
		if c.PosY > Y {
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
		X++
	}
	fmt.Printf("\n\n")
}

func Status(Board []Cell, theBlock Bloxor) (win, dead bool) {
	dead = false
	win = false
	X, Y := 0, 0
	for _, c := range Board {
		if c.PosY > Y {
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

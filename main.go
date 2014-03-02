package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Cell struct {
	Type  string
	PosX  int
	PosY  int
	Solid bool
}

type Block struct {
	PosX, PosY int
}

type Bloxor struct {
	// a bloxor is two block cubes attached
	A, B Block
}

func (b *Bloxor) isStacked() bool {
	if b.A.PosX == b.B.PosX && b.A.PosY == b.B.PosY {
		return true
	}
	return false
}

func (b *Bloxor) isVertical() bool {
	if b.A.PosX == b.B.PosX && b.A.PosY != b.B.PosY {
		return true
	}
	return false
}

func (b *Bloxor) isHorizontal() bool {
	if b.A.PosX != b.B.PosX && b.A.PosY == b.B.PosY {
		return true
	}
	return false
}

func (b *Bloxor) MoveUp() {
	if b.isStacked() {
		b.A.PosY -= 1
		b.B.PosY -= 2
	} else if b.isVertical() {
		if b.A.PosY < b.B.PosY {
			b.A.PosY -= 1
			b.B.PosY -= 2
		} else {
			b.A.PosY -= 2
			b.B.PosY -= 1
		}
	} else if b.isHorizontal() {
		b.A.PosY--
		b.B.PosY--
	} else {
		fmt.Println("Error, blocks not stacked, vertical, nor horizontal.")
	}
}

func (b *Bloxor) MoveDown() {
	if b.isStacked() {
		b.A.PosY += 1
		b.B.PosY += 2
	} else if b.isVertical() {
		if b.A.PosY < b.B.PosY {
			b.A.PosY += 2
			b.B.PosY += 1
		} else {
			b.A.PosY += 1
			b.B.PosY += 2
		}
	} else if b.isHorizontal() {
		b.A.PosY++
		b.B.PosY++
	} else {
		fmt.Println("Error, blocks not stacked, vertical, nor horizontal.")
	}
}

func (b *Bloxor) MoveLeft() {
	if b.isStacked() {
		b.A.PosX -= 1
		b.B.PosX -= 2
	} else if b.isHorizontal() {
		if b.A.PosX < b.B.PosX {
			b.A.PosX -= 1
			b.B.PosX -= 2
		} else {
			b.A.PosX -= 2
			b.B.PosX -= 1
		}
	} else if b.isVertical() {
		b.A.PosX--
		b.B.PosX--
	} else {
		fmt.Println("Error, blocks not stacked, vertical, nor horizontal.")
	}
}

func (b *Bloxor) MoveRight() {
	if b.isStacked() {
		b.A.PosX += 1
		b.B.PosX += 2
	} else if b.isHorizontal() {
		if b.A.PosX < b.B.PosX {
			b.A.PosX += 2
			b.B.PosX += 1
		} else {
			b.A.PosX += 1
			b.B.PosX += 2
		}
	} else if b.isVertical() {
		b.A.PosX++
		b.B.PosX++
	} else {
		fmt.Println("Error, blocks not stacked, vertical, nor horizontal.")
	}
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

func printBoard(Board []Cell, theBlock Bloxor) {
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

func status(Board []Cell, theBlock Bloxor) (win, dead bool) {
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

func main() {
	// TODO: upon death, go to beginning
	// TODO: upon win, go to next level
	// TODO: figure out how to have a solver
START:
	MapFile := "map1.txt"
	Board, theBlock := ReadMap(MapFile)
	numMoves := 0
	fmt.Println("Press (w,a,s,d + ENTER) to move")
	printBoard(Board, theBlock)
	input := bufio.NewReader(os.Stdin)
	for {
		command, err := input.ReadByte()
		if err != nil {
			fmt.Println(err)
		}
		numMoves++
		if err != nil {
			fmt.Println("err")
		}
		if string(command) == "w" {
			theBlock.MoveUp()
		} else if string(command) == "s" {
			theBlock.MoveDown()
		} else if string(command) == "a" {
			theBlock.MoveLeft()
		} else if string(command) == "d" {
			theBlock.MoveRight()
		} else {
			numMoves--
		}
		printBoard(Board, theBlock)
		fmt.Println("Moves: ", numMoves)

		isWinner, isDead := status(Board, theBlock)
		if isWinner || isDead {
			time.Sleep(2 * time.Second)
			goto START
		}
	}
}

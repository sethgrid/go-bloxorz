package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"bloxorz"
)

func main() {
	// TODO: upon win, go to next level
	// TODO: figure out how to have a solver
START:
	MapFile := "map1.txt"
	Board, theBlock := bloxorz.ReadBoardFromFile(MapFile)
	numMoves := 0
	fmt.Println("Press (w,a,s,d + ENTER) to move")
	bloxorz.PrintBoard(Board, theBlock)
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
		bloxorz.PrintBoard(Board, theBlock)
		fmt.Println("Moves: ", numMoves)

		isWinner, isDead := bloxorz.Status(Board, theBlock)
		if isWinner || isDead {
			time.Sleep(2 * time.Second)
			goto START
		}
	}
}

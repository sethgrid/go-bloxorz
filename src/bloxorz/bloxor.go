package bloxorz

/*
This file composes the bloxorz struct and defines how it moves
*/

import (
	"fmt"
)

// used to compose the Bloxor
type Block struct {
	PosX, PosY int
}

// a bloxor is two block cubes attached
type Bloxor struct {
	A, B Block
}

// think z-axis; both blocks at the same x,y coordinate
func (b *Bloxor) isStacked() bool {
	if b.A.PosX == b.B.PosX && b.A.PosY == b.B.PosY {
		return true
	}
	return false
}

// parallel to y-axis
func (b *Bloxor) isVertical() bool {
	if b.A.PosX == b.B.PosX && b.A.PosY != b.B.PosY {
		return true
	}
	return false
}

// parallel to x-axis
func (b *Bloxor) isHorizontal() bool {
	if b.A.PosX != b.B.PosX && b.A.PosY == b.B.PosY {
		return true
	}
	return false
}

/*
Move functions behave differently depending on block state.
*/

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
		// sanity check
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
		// sanity check
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
		// sanity check
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
		// sanity check
		fmt.Println("Error, blocks not stacked, vertical, nor horizontal.")
	}
}

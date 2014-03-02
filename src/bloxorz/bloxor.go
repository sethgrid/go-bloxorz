package bloxorz

import (
	"fmt"
)

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

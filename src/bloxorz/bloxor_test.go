package bloxorz

import (
	"testing"
)

func TestPositioning(t *testing.T) {
	// case 1: stacked
	blockA := Block{PosX: 3, PosY: 3}
	blockB := Block{PosX: 3, PosY: 3}
	bloxorStacked := Bloxor{A: blockA, B: blockB}
	isStacked := true
	isHorizontal := false
	isVertical := false
	checkPosition(t, bloxorStacked, isStacked, isHorizontal, isVertical)
	// case 2: horizontal
	blockA = Block{PosX: 2, PosY: 3}
	blockB = Block{PosX: 3, PosY: 3}
	bloxorHorizontal := Bloxor{A: blockA, B: blockB}
	isStacked = false
	isHorizontal = true
	isVertical = false
	checkPosition(t, bloxorHorizontal, isStacked, isHorizontal, isVertical)
	// case 3: vertical
	blockA = Block{PosX: 3, PosY: 2}
	blockB = Block{PosX: 3, PosY: 3}
	bloxorVertical := Bloxor{A: blockA, B: blockB}
	isStacked = false
	isHorizontal = false
	isVertical = true
	checkPosition(t, bloxorVertical, isStacked, isHorizontal, isVertical)
}

func checkPosition(t *testing.T, bloxor Bloxor, isStacked, isHorizontal, isVertical bool) {
	if bloxor.isStacked() != isStacked {
		t.Errorf("failed isStacked() - actual: %t, expected: %t", bloxor.isStacked(), isStacked)
	}
	if bloxor.isHorizontal() != isHorizontal {
		t.Errorf("failed isHorizontal() - actual: %t, expected: %t", bloxor.isHorizontal(), isHorizontal)
	}
	if bloxor.isVertical() != isVertical {
		t.Errorf("failed isVertical() - actual: %t, expected: %t", bloxor.isVertical(), isVertical)
	}
}

func TestMovesStacked(t *testing.T) {
	blockA := Block{PosX: 3, PosY: 3}
	blockB := Block{PosX: 3, PosY: 3}
	bloxor := Bloxor{A: blockA, B: blockB}

	bloxor.MoveDown()
	bloxor.MoveUp()
	bloxor.MoveRight()
	bloxor.MoveLeft()

	// should be back to where we started.
	// blockA and blockB were not mutated during the moves,
	// bloxor.A and bloxor.B were.
	if bloxor.A.PosX != blockA.PosX || bloxor.A.PosY != blockA.PosY {
		t.Errorf("block A move failure - actual:%v, expected:%v", bloxor.A, blockA)
	}
	if bloxor.B.PosX != blockB.PosX || bloxor.B.PosY != blockB.PosY {
		t.Errorf("block B move failure - actual:%v, expected:%v", bloxor.B, blockB)
	}
}

func TestMovesHorizontal(t *testing.T) {
	blockA := Block{PosX: 3, PosY: 2}
	blockB := Block{PosX: 3, PosY: 3}
	bloxor := Bloxor{A: blockA, B: blockB}

	bloxor.MoveDown()
	bloxor.MoveUp()
	bloxor.MoveRight()
	bloxor.MoveLeft()

	// should be back to where we started. However, A and B could have switched places.
	// TODO: fix that blocks can "switch" places.
	// blockA and blockB were not mutated during the moves,
	// bloxor.A and bloxor.B were.
	if bloxor.A.PosX != blockA.PosX || bloxor.A.PosY != blockA.PosY {
		if bloxor.A.PosX != blockB.PosX || bloxor.A.PosY != blockB.PosY {
			t.Errorf("block A move failure - actual:%v, expected:%v", bloxor.A, blockA)
		}
	}
	if bloxor.B.PosX != blockB.PosX || bloxor.B.PosY != blockB.PosY {
		if bloxor.B.PosX != blockA.PosX || bloxor.B.PosY != blockA.PosY {
			t.Errorf("block B move failure - actual:%v, expected:%v", bloxor.B, blockB)
		}
	}
}

func TestMovesVertical(t *testing.T) {
	blockA := Block{PosX: 2, PosY: 3}
	blockB := Block{PosX: 3, PosY: 3}
	bloxor := Bloxor{A: blockA, B: blockB}

	bloxor.MoveDown()
	bloxor.MoveUp()
	bloxor.MoveRight()
	bloxor.MoveLeft()

	// should be back to where we started. However, A and B could have switched places.
	// TODO: fix that blocks can "switch" places.
	// blockA and blockB were not mutated during the moves,
	// bloxor.A and bloxor.B were.
	if bloxor.A.PosX != blockA.PosX || bloxor.A.PosY != blockA.PosY {
		if bloxor.A.PosX != blockB.PosX || bloxor.A.PosY != blockB.PosY {
			t.Errorf("block A move failure - actual:%v, expected:%v", bloxor.A, blockA)
		}
	}
	if bloxor.B.PosX != blockB.PosX || bloxor.B.PosY != blockB.PosY {
		if bloxor.B.PosX != blockA.PosX || bloxor.B.PosY != blockA.PosY {
			t.Errorf("block B move failure - actual:%v, expected:%v", bloxor.B, blockB)
		}
	}
}

package main

import "math/rand"

type Board struct {
	cells []Cell
}

func InitBoard() *Board {
	// we are "escaping to the heap", meaning at compile time, the GO compiler will notice that we return the address of a local
	// variable that would be on the stack and instead moves it to the heap.
	board := Board{
		cells: make([]Cell, GRID_WIDTH*GRID_HEIGHT, GRID_WIDTH*GRID_HEIGHT),
	}

	for i := 0; i < int(GRID_HEIGHT*GRID_WIDTH); i++ {
		if (rand.Int() % 3) == 2 {
			board.cells[i] = CreateCell(true)
		} else {
			board.cells[i] = CreateCell(false)
		}
	}
	return &board
}

func (board Board) CheckStatus(x int32, y int32) bool {
	if x < 0 || x >= GRID_WIDTH || y < 0 || y >= GRID_HEIGHT {
		return false
	} else {
		return board.cells[(y*GRID_WIDTH)+x].alive
	}
}

func (board Board) SetNextStatus(x int32, y int32, status bool) {
	board.cells[(y*GRID_WIDTH)+x].next_round = status
}

func (board Board) aliveNeighbours(x int32, y int32) int {
	count := 0

	if board.CheckStatus(x-1, y-1) {
		count++
	}

	if board.CheckStatus(x+0, y-1) {
		count++
	}

	if board.CheckStatus(x+1, y-1) {
		count++
	}

	if board.CheckStatus(x-1, y+0) {
		count++
	}

	if board.CheckStatus(x+1, y+0) {
		count++
	}

	if board.CheckStatus(x-1, y+1) {
		count++
	}

	if board.CheckStatus(x+0, y+1) {
		count++
	}

	if board.CheckStatus(x+1, y+1) {
		count++
	}

	return count
}

func (board Board) PlayRound() {
	var y int32 = 0
	var x int32 = 0
	for y = 0; y < GRID_HEIGHT; y++ {
		for x = 0; x < GRID_WIDTH; x++ {
			// rules

			neighbours := board.aliveNeighbours(x, y)
			status := board.CheckStatus(x, y)

			if status {
				// the cell is alive
				if neighbours < 2 || neighbours > 3 {
					board.SetNextStatus(x, y, false)
				} else if neighbours == 2 || neighbours == 3 {
					board.SetNextStatus(x, y, true)
				}
			} else {
				if neighbours == 3 {
					board.SetNextStatus(x, y, true)
				}
			}
		}
	}

	// apply the outcome
	for i := range board.cells {
		board.cells[i].alive = board.cells[i].next_round
	}
}

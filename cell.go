package main

type Cell struct {
	alive      bool
	next_round bool
}

func CreateCell(isAlive bool) Cell {
	return Cell{
		alive:      isAlive,
		next_round: false,
	}
}

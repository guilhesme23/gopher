package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	cells [][]Cell
	w     int32
}

func (b Board) Draw() {
	for _, row := range b.cells {
		for _, c := range row {
			c.Draw(b.w)
		}
	}
}

func NewBoard(width int, height int, w int32) Board {
	rSize := width / int(w)
	cSize := height / int(w)
	b := Board{}
	b.w = w
	cells := make([][]Cell, rSize)
	for i := 0; i < rSize; i++ {
		row := make([]Cell, cSize)
		for j := 0; j < cSize; j++ {
			row[j] = Cell{
				pos:   rl.Vector2{X: float32(i * int(w)), Y: float32(j * int(w))},
				alive: true,
			}

			if rand.Float32() < 0.5 {
				row[j].alive = false
			}
		}
		cells[i] = row
	}
	b.cells = cells
	return b
}

func (b Board) NextGen() Board {
	rSize := len(b.cells)
	cSize := len(b.cells[0])
	cells := make([][]Cell, rSize)
	for i := range b.cells {
		row := make([]Cell, cSize)
		for j := range b.cells[i] {
			n := b.countNeighbors(i, j)
			row[j] = Cell{
				pos:   b.cells[i][j].pos,
				alive: b.cells[i][j].alive,
			}

			if n == 3 {
				row[j].alive = true
			}

			if n < 2 || n > 3 {
				row[j].alive = false
			}
		}
		cells[i] = row
	}

	return Board{
		w:     b.w,
		cells: cells,
	}
}

func (b Board) countNeighbors(i int, j int) int {
	row := []int{i - 1, i, i + 1}
	col := []int{j - 1, j, j + 1}
	maxRow := len(b.cells)
	maxCol := len(b.cells[0])

	total := 0
	for _, k := range row {
		for _, v := range col {
			if k == i && v == j {
				continue
			}

			k = int(math.Abs(float64(k+maxRow))) % maxRow
			v = int(math.Abs(float64(v+maxCol))) % maxCol

			if b.cells[k][v].alive {
				total++
			}
		}
	}
	return total
}

func (b *Board) SetCell(i int, j int, state bool) {
	b.cells[i][j].alive = state
}

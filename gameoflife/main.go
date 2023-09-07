package main

import (
	"flag"
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width := flag.Int("width", 800, "Window width")
	height := flag.Int("height", 450, "Window height")
	size := flag.Int("size", 5, "Cell size")
	flag.Parse()
	rl.InitWindow(int32(*width), int32(*height), "Conways Game of Life")
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	board := NewBoard(*width, *height, int32(*size))
	fmt.Println(*width, *height)
	fmt.Println(len(board.cells), len(board.cells[0]))
	run := true
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyN) {
			board = NewBoard(*width, *height, int32(*size))
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			run = false
			// Set cells alive
			pos := rl.GetMousePosition()
			wSize := len(board.cells)
			hSize := len(board.cells[0])
			mx := int(math.Round(float64((float32(wSize) / float32(*width)) * pos.X)))
			my := int(math.Round(float64((float32(hSize) / float32(*height)) * pos.Y)))
			if (mx+my) > my && my < hSize && mx < wSize {
				board.SetCell(mx, my, true)
			}
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			run = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Beige)
		board.Draw()
		if run {
			board = board.NextGen()
		}
		rl.EndDrawing()
	}
}

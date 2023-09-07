package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	pos   rl.Vector2
	alive bool
}

func (c Cell) Draw(w int32) {
	if c.alive {
		rl.DrawRectangle(int32(c.pos.X), int32(c.pos.Y), w, w, rl.Blue)
	}
}

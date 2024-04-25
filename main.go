package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
)

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "gotari")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGreen)

		rl.DrawCircle(WINDOW_WIDTH/2, WINDOW_HEIGHT/2, WINDOW_HEIGHT/2, rl.DarkBlue)

		rl.EndDrawing()
	}
}

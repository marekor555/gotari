package main

import (
	line "gotari/entities"

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

	player := line.NewLine()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.Draw()

		rl.EndDrawing()

		playerDirection := 0
		if rl.IsKeyDown(rl.KeyA) {
			playerDirection--
		}
		if rl.IsKeyDown(rl.KeyD) {
			playerDirection++
		}
		player.Move(playerDirection)
	}
}

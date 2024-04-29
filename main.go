package main

import (
	"gotari/constants"
	ball_ "gotari/entities/ball"
	"gotari/entities/block"
	"gotari/entities/line"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	blocks [8][3]block.Block
)

func main() {
	rl.InitWindow(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT, "gotari")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for x := range 8 {
		for y := range 3 {
			blocks[x][y] = block.NewBlock([2]float32{float32(x) * constants.BLOCK_WIDTH, float32(y)*constants.BLOCK_HEIGHT + 50})
		}
	}
	player := line.NewLine()
	ball := ball_.NewBall()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.Draw()
		ball.Draw()

		for x := range 8 {
			for y := range 3 {
				blocks[x][y].Draw()
			}
		}

		rl.EndDrawing()

		playerDirection := 0
		if rl.IsKeyDown(rl.KeyA) {
			playerDirection--
		}
		if rl.IsKeyDown(rl.KeyD) {
			playerDirection++
		}
		player.Move(playerDirection)
		ball.Move(&blocks, player)
	}
}

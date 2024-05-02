package main

import (
	"fmt"
	"gotari/constants"
	ball_ "gotari/entities/ball"
	"gotari/entities/block"
	"gotari/entities/line"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	blocks [constants.BLOCK_COLUMNS][constants.BLOCK_ROWS]block.Block
	score  int = 0
)

func main() {
	rl.InitWindow(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT, "gotari")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for x := range constants.BLOCK_COLUMNS {
		for y := range constants.BLOCK_ROWS {
			blocks[x][y] = block.NewBlock([2]float32{float32(x) * constants.BLOCK_WIDTH, float32(y)*constants.BLOCK_HEIGHT + 50})
		}
	}
	player := line.NewLine()
	ball := ball_.NewBall()
	lives := constants.LIVES

	for !rl.WindowShouldClose() {
		if lives == 0 {
			break
		}
		if ball.CheckFloorTouch() {
			score -= 5
			lives--
			ball.ResetPos()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for x := range constants.BLOCK_COLUMNS {
			for y := range constants.BLOCK_ROWS {
				blocks[x][y].Draw()
			}
		}

		player.Draw()
		ball.Draw()

		rl.DrawText(fmt.Sprintf("Lives: %d\nScore: %d", lives, score), 0, 0, constants.FONT_SIZE, rl.Green)

		rl.EndDrawing()

		playerDirection := 0
		if rl.IsKeyDown(rl.KeyA) {
			playerDirection--
		}
		if rl.IsKeyDown(rl.KeyD) {
			playerDirection++
		}
		player.Move(playerDirection)
		score += ball.Move(&blocks, player)
	}
}

package main

import (
	"fmt"
	"gotari/constants"
	"gotari/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	bestScore = 0
)

func main() {
	rl.InitWindow(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT, "gotari")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		menu()
		countDown()
		newScore := game()
		fmt.Println(newScore)
		if newScore > bestScore {
			bestScore = newScore
		}
	}
}

func countDown() {
	timer := constants.WAIT_TIME * 60

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		timer--
		rl.DrawRectangle((constants.WINDOW_WIDTH/2)-(constants.FONT_SIZE*2), (constants.WINDOW_HEIGHT/2)-(constants.FONT_SIZE*2), constants.FONT_SIZE, constants.FONT_SIZE*2, rl.White)
		rl.DrawText(fmt.Sprintf("%d", timer/60), (constants.WINDOW_WIDTH/2)-(constants.FONT_SIZE*2), (constants.WINDOW_HEIGHT/2)-(constants.FONT_SIZE*2), constants.FONT_SIZE*2, rl.Red)
		rl.EndDrawing()
		if timer <= 0 {
			return
		}
	}
}

func menu() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Press *space* to start", 0, 0, constants.FONT_SIZE*2, rl.Lime)
		if bestScore != 0 {
			rl.DrawText(fmt.Sprintf("Best score: %d", bestScore), 0, constants.FONT_SIZE*2, constants.FONT_SIZE*2, rl.Lime)
		}
		rl.EndDrawing()
		if rl.IsKeyPressed(rl.KeySpace) {
			break
		}
	}
}

func game() int {
	var blocks [constants.BLOCK_COLUMNS][constants.BLOCK_ROWS]entities.Block
	score := 0

	for x := range constants.BLOCK_COLUMNS {
		for y := range constants.BLOCK_ROWS {
			blocks[x][y] = entities.NewBlock([2]float32{float32(x) * constants.BLOCK_WIDTH, float32(y)*constants.BLOCK_HEIGHT + 50})
		}
	}
	player := entities.NewLine()
	ball := entities.NewBall()
	lives := constants.LIVES

	for !rl.WindowShouldClose() {
		if lives == 0 {
			break
		}
		if ball.CheckFloorTouch() {
			score -= 5
			lives--
			ball.ResetPos()
			countDown()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		blocksLeft := 0
		for x := range constants.BLOCK_COLUMNS {
			for y := range constants.BLOCK_ROWS {
				blocks[x][y].Draw()
				if !blocks[x][y].Destroyed {
					blocksLeft++
				}
			}
		}
		if blocksLeft == 0 {
			return score
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
		destroyedBlocks := ball.Move(&blocks, player)
		score += destroyedBlocks
	}
	return score
}

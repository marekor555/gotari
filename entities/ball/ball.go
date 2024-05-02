package ball

import (
	"gotari/constants"
	"gotari/entities/block"
	"gotari/entities/line"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Point     rl.Vector2
	Direction [2]float32
	Speed     float32
}

func (ball Ball) Draw() {
	rl.DrawCircle(int32(ball.Point.X), int32(ball.Point.Y), constants.BALL_RADIUS, rl.White)
}

func (ball Ball) CheckFloorTouch() bool {
	return rl.CheckCollisionPointLine(ball.Point, rl.NewVector2(0, float32(constants.WINDOW_HEIGHT)), rl.NewVector2(float32(constants.WINDOW_WIDTH), float32(constants.WINDOW_HEIGHT)), int32(constants.BALL_RADIUS))
}

func (ball *Ball) Move(blocks *[constants.BLOCK_COLUMNS][constants.BLOCK_ROWS]block.Block, player line.Line) int {
	newPoints := 0
	for x, column := range blocks {
		for y := range column {
			block := &blocks[x][y]
			if block.Destroyed {
				continue
			}
			if rl.CheckCollisionCircleRec(ball.Point, constants.BALL_RADIUS, block.Collider) {
				if ball.Point.Y > block.Collider.Y+(block.Collider.Height/2) || ball.Point.Y < block.Collider.Y+(block.Collider.Height/2) {
					ball.Direction[1] = -ball.Direction[1]
				} else if ball.Point.X > block.Collider.X+(block.Collider.Width/2) || ball.Point.X < block.Collider.X+(block.Collider.Width/2) {
					ball.Direction[0] = -ball.Direction[0]
				}
				if ball.Speed <= constants.BALL_SPEED_MAX {
					ball.Speed += constants.BALL_SPEED_STEP
				}
				newPoints++
				block.Destroy()
			}
		}
	}

	if rl.CheckCollisionCircleRec(ball.Point, constants.BALL_RADIUS, player.Collider) {
		ball.Speed = constants.BALL_SPEED_START
		ball.Direction[1] = -1
		ball.Direction[0] = ((ball.Point.X - (player.Collider.X + player.Collider.Width/2)) / player.Collider.Width) * 2
	}

	if ball.Point.X <= 0 || ball.Point.X+constants.BALL_RADIUS >= float32(constants.WINDOW_WIDTH) {
		ball.Direction[0] = -ball.Direction[0]
	}

	if ball.Point.Y <= 0 || ball.Point.Y+constants.BALL_RADIUS >= float32(constants.WINDOW_HEIGHT) {
		ball.Direction[1] = -ball.Direction[1]
	}

	ball.Point.X += float32(ball.Direction[0]) * ball.Speed
	ball.Point.Y += float32(ball.Direction[1]) * ball.Speed

	return newPoints
}

func (ball *Ball) ResetPos() {
	ball.Direction = [2]float32{-1, -1}
	ball.Point = rl.NewVector2(float32(constants.WINDOW_WIDTH/2), float32(constants.WINDOW_HEIGHT)-constants.LINE_HEIGHT_OFFSET*2)
}

func NewBall() Ball {
	return Ball{
		Direction: [2]float32{-1, -1},
		Point:     rl.NewVector2(float32(constants.WINDOW_WIDTH/2), float32(constants.WINDOW_HEIGHT)-constants.LINE_HEIGHT_OFFSET*2),
		Speed:     constants.BALL_SPEED_START,
	}
}

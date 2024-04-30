package ball

import (
	"gotari/constants"
	"gotari/entities/block"
	"gotari/entities/line"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Point     rl.Vector2
	Direction [2]int
	Speed     float32
}

func (ball Ball) Draw() {
	rl.DrawCircle(int32(ball.Point.X), int32(ball.Point.Y), constants.BALL_RADIUS, rl.White)
}

func (ball *Ball) Move(blocks *[constants.BLOCK_COLUMNS][constants.BLOCK_ROWS]block.Block, player line.Line) {
	for x, column := range blocks {
		for y := range column {
			block := &blocks[x][y]
			if block.Destroyed {
				continue
			}
			if rl.CheckCollisionCircleRec(ball.Point, constants.BALL_RADIUS, block.Collider) {
				if ball.Point.Y > block.Collider.Y+(block.Collider.Height/2) {
					ball.Direction[1] = 1
				} else if ball.Point.Y < block.Collider.Y+(block.Collider.Height/2) {
					ball.Direction[1] = -1
				}

				if ball.Point.X > block.Collider.X+(block.Collider.Width/2) {
					ball.Direction[0] = 1
				} else if ball.Point.X < block.Collider.X+(block.Collider.Width/2) {
					ball.Direction[0] = -1
				}
				if ball.Speed <= constants.BALL_SPEED_MAX {
					ball.Speed += 0.5
				}
				block.Destroy()
			}
		}
	}

	if rl.CheckCollisionCircleRec(ball.Point, constants.BALL_RADIUS, player.Collider) {
		ball.Speed = constants.BALL_SPEED_START
		ball.Direction[1] = -1
		if ball.Point.X > player.Collider.X+(player.Collider.Width/2) {
			ball.Direction[0] = 1
		} else {
			ball.Direction[0] = -1
		}
	}

	if ball.Point.X <= 0 {
		ball.Direction[0] = 1
	} else if ball.Point.X+constants.BALL_RADIUS >= float32(constants.WINDOW_WIDTH) {
		ball.Direction[0] = -1
	}

	if ball.Point.Y <= 0 {
		ball.Direction[1] = 1
	} else if ball.Point.Y+constants.BALL_RADIUS >= float32(constants.WINDOW_HEIGHT) {
		ball.Direction[1] = -1
	}

	ball.Point.X += float32(ball.Direction[0]) * ball.Speed
	ball.Point.Y += float32(ball.Direction[1]) * ball.Speed
}

func NewBall() Ball {
	return Ball{
		Direction: [2]int{-1, -1},
		Point:     rl.NewVector2(float32(constants.WINDOW_WIDTH/2), float32(constants.WINDOW_HEIGHT)-constants.LINE_HEIGHT_OFFSET*2),
		Speed:     constants.BALL_SPEED_START,
	}
}

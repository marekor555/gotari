package ball

import (
	"gotari/constants"
	"gotari/entities/block"
	"gotari/entities/line"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Collider  rl.Rectangle
	Direction [2]int
}

func (ball Ball) Draw() {
	rl.DrawRectangle(ball.Collider.ToInt32().X, ball.Collider.ToInt32().Y, ball.Collider.ToInt32().Width, ball.Collider.ToInt32().Height, rl.White)
}

func (ball *Ball) Move(blocks *[constants.BLOCK_COLUMNS][constants.BLOCK_ROWS]block.Block, player line.Line) {
	for x, column := range blocks {
		for y := range column {
			block := &blocks[x][y]
			if block.Destroyed {
				continue
			}
			if rl.CheckCollisionRecs(ball.Collider, block.Collider) {
				if ball.Collider.Y > block.Collider.Y {
					ball.Direction[1] = 1
				} else if ball.Collider.Y < block.Collider.Y {
					ball.Direction[1] = -1
				}

				if ball.Collider.X > block.Collider.X {
					ball.Direction[0] = 1
				} else if ball.Collider.X < block.Collider.X {
					ball.Direction[0] = 1
				}
				block.Destroy()
			}
		}
	}

	if rl.CheckCollisionRecs(ball.Collider, player.Collider) {
		ball.Direction[1] = -1
	}

	if ball.Collider.X <= 0 {
		ball.Direction[0] = 1
	} else if ball.Collider.X+constants.BALL_WIDTH >= float32(constants.WINDOW_WIDTH) {
		ball.Direction[0] = -1
	}

	if ball.Collider.Y <= 0 {
		ball.Direction[1] = 1
	} else if ball.Collider.Y+constants.BALL_HEIGHT >= float32(constants.WINDOW_HEIGHT) {
		ball.Direction[1] = -1
	}

	ball.Collider.X += float32(ball.Direction[0]) * constants.BALL_SPEED
	ball.Collider.Y += float32(ball.Direction[1]) * constants.BALL_SPEED
}

func NewBall() Ball {
	return Ball{
		Direction: [2]int{-1, -1},
		Collider:  rl.NewRectangle(float32(constants.WINDOW_HEIGHT/2), float32(constants.WINDOW_HEIGHT)-75, constants.BALL_WIDTH, constants.BALL_HEIGHT),
	}
}

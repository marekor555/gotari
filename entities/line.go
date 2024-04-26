package line

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
	SPEED         = 10
)

type Line struct {
	Collider rl.Rectangle
}

func (line *Line) Move(direction int) {
	if direction == -1 && line.Collider.X <= 0 {
		return
	} else if direction == 1 && line.Collider.X+line.Collider.Width >= WINDOW_WIDTH {
		return
	}
	line.Collider.X += float32(direction) * 10
}

func (line Line) Draw() {
	rl.DrawRectangle(line.Collider.ToInt32().X, line.Collider.ToInt32().Y, line.Collider.ToInt32().Width, line.Collider.ToInt32().Height, rl.White)
}

func NewLine() Line {
	return Line{
		Collider: rl.NewRectangle(0, WINDOW_HEIGHT-25, 200, 25),
	}
}

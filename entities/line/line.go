package line

import (
	"gotari/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Line struct {
	Collider rl.Rectangle
}

func (line *Line) Move(direction int) {
	if direction == -1 && line.Collider.X+(line.Collider.Width/2) <= 0 {
		return
	} else if direction == 1 && line.Collider.X+(line.Collider.Width/2) >= float32(constants.WINDOW_WIDTH) {
		return
	}
	line.Collider.X += float32(direction) * constants.LINE_SPEED
}

func (line Line) Draw() {
	rl.DrawRectangle(line.Collider.ToInt32().X, line.Collider.ToInt32().Y, line.Collider.ToInt32().Width, line.Collider.ToInt32().Height, rl.White)
}

func NewLine() Line {
	return Line{
		Collider: rl.NewRectangle(0, float32(constants.WINDOW_HEIGHT)-constants.LINE_HEIGHT_OFFSET, constants.LINE_WIDTH, constants.LINE_HEIGHT),
	}
}

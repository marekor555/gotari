package block

import (
	"gotari/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Collider  rl.Rectangle
	Destroyed bool
}

func (block Block) Draw() {
	rl.DrawRectangle(block.Collider.ToInt32().X, block.Collider.ToInt32().Y, block.Collider.ToInt32().Width, block.Collider.ToInt32().Height, rl.DarkBlue)
	rl.DrawRectangleLines(block.Collider.ToInt32().X, block.Collider.ToInt32().Y, block.Collider.ToInt32().Width, block.Collider.ToInt32().Height, rl.Blue)
}

func (block Block) CheckCollision(rect rl.Rectangle) bool {
	return rl.CheckCollisionRecs(block.Collider, rect)
}

func (block *Block) Destroy() {
	block.Destroyed = true
}

func NewBlock(pos [2]float32) Block {
	return Block{
		Collider:  rl.NewRectangle(pos[0], pos[1], constants.BLOCK_WIDTH, constants.BLOCK_HEIGHT),
		Destroyed: false,
	}
}

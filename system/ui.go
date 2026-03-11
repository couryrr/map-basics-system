package system

import (
	"github.com/couryrr/map-basics-system/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	// ...
}

func (mm *MainMenu) Draw() {
	rl.DrawRectangleLinesEx(rl.NewRectangle(0, 0, 480, 270), 1, rl.Red)
}

type InGameOverlay struct {
	IsCollision bool
}

func (igo *InGameOverlay) Update(destRec rl.Rectangle, scale float32) {
	point := rl.GetMousePosition()
	virtualX := (point.X - destRec.X) / scale
	virtualY := (point.Y - destRec.Y) / scale
	igo.IsCollision = rl.CheckCollisionPointRec(rl.NewVector2(virtualX, virtualY), rl.NewRectangle(64, float32(config.VirtualHeight-64), 32, 32))
}
func (igo *InGameOverlay) Draw() {
	posY := config.VirtualHeight - 64
	sizeX := config.VirtualWidth - 128
	rl.DrawRectangleLinesEx(rl.NewRectangle(64, float32(posY), float32(sizeX), 64), 1, rl.DarkGray)
	width := 1
	color := rl.DarkBlue
	if igo.IsCollision {
		width = 5
		color = rl.Red
	}
	rl.DrawRectangleLinesEx(rl.NewRectangle(64, float32(posY), 32, 32), float32(width), color)
}

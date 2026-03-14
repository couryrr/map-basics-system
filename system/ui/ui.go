package ui

import (
	"fmt"

	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/system/renderer"
	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	// ...
}

func (mm *MainMenu) Draw() {
	rl.DrawRectangleLinesEx(rl.NewRectangle(0, 0, 480, 270), 1, rl.Red)
}

type InGameOverlay struct {
}

// TODO: intentionally bad atm just getting some item functioning
func (igo *InGameOverlay) Draw(world world.World, rCtx *renderer.RenderContext) {
	posY := config.VirtualHeight - 64
	sizeX := config.VirtualWidth - 128
	rl.DrawRectangleLinesEx(rl.NewRectangle(64, float32(posY), float32(sizeX), 64), 1, rl.DarkGray)

	for i := range world.Items {
		point := rCtx.ScreenToVirtual(rl.GetMousePosition())
		pos := rl.NewVector2(float32(64+i*32), float32(posY))
		rec := rl.NewRectangle(pos.X, pos.Y, 32, 32)
		width := 1
		color := rl.DarkBlue
		if rl.CheckCollisionPointRec(point, rec) {
			width = 5
			color = rl.Red
		}
		rl.DrawText(fmt.Sprintf("%.0fx%.0f", world.Items[i].Size.X, world.Items[i].Size.Y), int32(pos.X+2), int32(pos.Y+2), 12, rl.DarkBlue)
		rl.DrawRectangleLinesEx(rec, float32(width), color)
	}
}

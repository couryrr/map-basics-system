package system

import rl "github.com/gen2brain/raylib-go/raylib"

type MainMenu struct {
	// ...
}

func (mm *MainMenu) Draw(renderTexture rl.RenderTexture2D) {
	rl.BeginTextureMode(renderTexture)
	rl.ClearBackground(rl.White)
	rl.DrawRectangleLinesEx(rl.NewRectangle(0, 0, 480, 270), 1, rl.Red)
	rl.EndTextureMode()

}

func CreateMainMenu() MainMenu {
	return MainMenu{}
}

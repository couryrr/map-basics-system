package main

import (
	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	VirtualWidth  float32 = 320
	VirtualHeight float32 = 180
)

var ()

func main() {
	game := CreateGamePlease()
	screenSetting := system.CreateScreenSetting(rl.NewVector2(game.settings.ScreenWidth, game.settings.ScreenHeight), rl.NewVector2(VirtualWidth, VirtualHeight), rl.NewVector2(game.settings.WindowedScreenWidth, game.settings.WindowededScreenHeight))

	rl.InitWindow(int32(screenSetting.ScreenSize.X), int32(screenSetting.ScreenSize.Y), "Map Basics")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	world := system.CreateWorld(system.CreateGameCamera(rl.Vector2Scale(screenSetting.ScreenSize, 0.5), rl.Vector2Scale(screenSetting.VirtualScreenSize, 0.5), 0.0, 1.0),
		rl.LoadRenderTexture(int32(screenSetting.VirtualScreenSize.X), int32(screenSetting.VirtualScreenSize.Y)))
	defer world.UnloadWorld()
	for !rl.WindowShouldClose() {
		system.HandleInput(&screenSetting, &world.Camera)
		world.Draw()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		source := rl.NewRectangle(0, 0, screenSetting.VirtualScreenSize.X, -screenSetting.VirtualScreenSize.Y)
		dest := rl.NewRectangle(screenSetting.DestinationPosition.X, screenSetting.DestinationPosition.Y, screenSetting.DestinationSize.X, screenSetting.DestinationSize.Y)
		rl.DrawTexturePro(world.RenderTexture.Texture, source, dest, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

package main

import (
	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := CreateGame(rl.NewVector2(1920, 1080), rl.NewVector2(1920, 1080))

	rl.InitWindow(int32(game.GameSettings.ScreenSetting.ScreenSize.X), int32(game.GameSettings.ScreenSetting.ScreenSize.Y), "Map Basics")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	world := system.CreateWorld(rl.NewVector2(320, 180))
	defer world.UnloadWorld()

	source := rl.NewRectangle(0, 0, world.WorldScreenSize.X, -world.WorldScreenSize.Y)
	viewport := game.GameSettings.ScreenSetting.CalculateViewport(world.WorldScreenSize)

	for !rl.WindowShouldClose() {
		system.HandleInput(&game.GameSettings.ScreenSetting, world)
		world.Draw()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(world.RenderTexture.Texture, source, viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

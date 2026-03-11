package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := CreateGame(rl.NewVector2(1920, 1080), rl.NewVector2(1920, 1080))

	rl.InitWindow(int32(game.SystemSettings.ScreenSetting.ScreenSize.X), int32(game.SystemSettings.ScreenSetting.ScreenSize.Y), "Map Basics")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	game.LoadGame()
	world := system.CreateWorld()
	defer world.UnloadWorld()

	source := rl.NewRectangle(0, 0, config.VirtualWidth, -config.VirtualHeight)
	viewport, scale := game.SystemSettings.ScreenSetting.CalculateViewport(rl.NewVector2(config.VirtualWidth, config.VirtualHeight))

	igo := system.InGameOverlay{
		IsCollision: false,
	}

	for !rl.WindowShouldClose() {
		system.HandleInput(&game.SystemSettings.ScreenSetting, game.GameCamera)
		igo.Update(viewport, scale)
		rl.BeginTextureMode(*game.RenderTexture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*game.GameCamera.Camera)
		world.Draw(game.GameCamera.Camera.Target)
		rl.EndMode2D()
		igo.Draw()
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(game.RenderTexture.Texture, source, viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

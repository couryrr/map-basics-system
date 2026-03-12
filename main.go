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

	game.Init()
	defer game.Unload()

	world := system.CreateWorld()
	defer world.UnloadWorld()

	source := rl.NewRectangle(0, 0, config.VirtualWidth, -config.VirtualHeight)
	game.RenderContext.Update(game.SystemSettings.ScreenSetting.ScreenSize)

	igo := system.InGameOverlay{
		IsCollision: false,
	}

	for !rl.WindowShouldClose() {
		system.HandleInput(&game.SystemSettings.ScreenSetting, game.GameCamera, game.RenderContext)
		igo.Update(game.RenderContext)
		rl.BeginTextureMode(*game.RenderContext.RenderTexture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*game.GameCamera.Camera)
		world.Draw(game.GameCamera.Camera.Target)
		rl.EndMode2D()
		igo.Draw()
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(game.RenderContext.RenderTexture.Texture, source, game.RenderContext.Viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

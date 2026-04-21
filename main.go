package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/framework/keyboard"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := NewGame(rl.NewVector2(1920, 1080), rl.NewVector2(1920, 1080))

	rl.InitWindow(int32(game.SystemSettings.ScreenSetting.ScreenSize.X), int32(game.SystemSettings.ScreenSetting.ScreenSize.Y), "Map Basics")
	rl.SetTargetFPS(60)

	game.LoadWorld()
	game.LoadResources()

	defer game.Unload()
	defer rl.CloseWindow()

	source := rl.NewRectangle(0, 0, config.VirtualWidth, -config.VirtualHeight)
	game.RenderContext.Update(game.SystemSettings.ScreenSetting.ScreenSize)

	game.EventQueue.Subscribe(100, game.ToggleScreenSize)

	for !rl.WindowShouldClose() {
		game.EventQueue.Drain()
		event := keyboard.HandleInput(game.RenderContext)
		game.UiManager.Update(event)
		if game.UiManager.Hovered != nil {
			rl.TraceLog(rl.LogInfo, "%v", game.UiManager.Hovered)
		}
		game.Update()
		rl.BeginTextureMode(*game.RenderContext.RenderTexture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*game.GameCamera.Camera)
		game.Draw()
		rl.EndMode2D()
		if game.UiManager != nil {
			game.UiManager.Draw()
		}
		rl.EndTextureMode()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(game.RenderContext.RenderTexture.Texture, source, game.RenderContext.Viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

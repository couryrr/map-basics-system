package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/framework/keyboard"
	"github.com/couryrr/map-basics-system/system/ui"
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

	game.Broker.Register(keyboard.TopicScreenToggle, game.ToggleScreenSize)
	game.Broker.Register(keyboard.TopicInputRotate, game.Player.Rotate)
	game.Broker.Register(keyboard.TopicInputRotateReset, game.Player.RotateReset)
	game.Broker.Register(keyboard.TopicInputMove, game.Player.Move)
	game.Broker.Register(keyboard.TopicInputZoom, game.Player.Zoom)
	game.Broker.Register(keyboard.TopicInputCursorMoved, game.Igo.HandleMouseEvent)
	game.Broker.Register(ui.TopicUiHotbarInteraction, game.Player.HandleHotbarInteraction)

	for !rl.WindowShouldClose() {
		keyboard.HandleInput(game.Broker, game.RenderContext)
		game.Update()
		rl.BeginTextureMode(*game.RenderContext.RenderTexture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*game.GameCamera.Camera)
		game.Draw()
		rl.EndMode2D()
		game.Igo.Draw()
		rl.EndTextureMode()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(game.RenderContext.RenderTexture.Texture, source, game.RenderContext.Viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

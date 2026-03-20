package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/system/controller"
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

	game.Broker.Register(controller.TopicScreenToggle, game.ToggleScreenSize)
	game.Broker.Register(controller.TopicInputRotate, game.Player.Rotate)
	game.Broker.Register(controller.TopicInputRotateReset, game.Player.RotateReset)
	game.Broker.Register(controller.TopicInputMove, game.Player.Move)
	game.Broker.Register(controller.TopicInputZoom, game.Player.Zoom)
	game.Broker.Register(controller.TopicInputCursorMoved, game.Igo.CheckIntersection)
	game.Broker.Register(ui.TopicUiHotbarInteraction, game.Player.HandleHotbarInteraction)

	for !rl.WindowShouldClose() {
		controller.HandleInput(game.Broker, game.RenderContext)
		game.Update()
		rl.BeginTextureMode(*game.RenderContext.RenderTexture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*game.GameCamera.Camera)
		game.Draw()
		rl.EndMode2D()
		game.Igo.Draw(&game)
		rl.EndTextureMode()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexturePro(game.RenderContext.RenderTexture.Texture, source, game.RenderContext.Viewport, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

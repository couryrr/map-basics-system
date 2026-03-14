package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SystemSettings struct {
	ScreenSetting system.ScreenSetting
}

type Game struct {
	Broker         *system.Broker
	GameCamera     *system.GameCamera
	Player         *system.Player
	RenderContext  *system.RenderContext
	SystemSettings SystemSettings
	IsFullScreen   bool
}

func (game *Game) LoadResources() {
	renderContext := system.CreateRenderContext(
		config.VirtualWidth,
		config.VirtualHeight,
		game.SystemSettings.ScreenSetting.ScreenSize,
	)
	game.RenderContext = &renderContext

	player := system.CreatePlayer(rl.NewVector2(renderContext.VirtualWidth/2, renderContext.VirtualHeight/2))
	game.Player = &player

	camera := system.CreateGameCamera(rl.NewVector2(player.Position.X, player.Position.Y), rl.NewVector2(float32(renderContext.VirtualWidth/2), float32(renderContext.VirtualHeight/2)), 0.0, 1.0)
	game.GameCamera = &camera
}

func (game *Game) ToggleScreenSize(message system.Message) {
	rl.ToggleFullscreen()
	newScreenSize := rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	if !game.IsFullScreen {
		newScreenSize = game.SystemSettings.ScreenSetting.WindowedScreenSize
	}
	game.SystemSettings.ScreenSetting.ScreenSize = newScreenSize
	game.IsFullScreen = !game.IsFullScreen
	game.RenderContext.Update(game.SystemSettings.ScreenSetting.ScreenSize)
}

func (game *Game) Update() {
	game.GameCamera.Camera.Rotation = game.Player.Rotation
	game.GameCamera.Camera.Target = game.Player.Position
	game.GameCamera.Camera.Zoom = game.Player.ZoomLevel
}

func (game *Game) Unload() {
	rl.UnloadRenderTexture(*game.RenderContext.RenderTexture)
}

func CreateGame(windowedScreenSize, screenSize rl.Vector2) Game {
	// TODO: This will load setting from files.
	broker := system.CreateBroker()
	return Game{
		Broker:        &broker,
		Player:        nil,
		GameCamera:    nil,
		RenderContext: nil,
		SystemSettings: SystemSettings{
			ScreenSetting: system.CreateScreenSetting(screenSize, windowedScreenSize),
		},
	}
}

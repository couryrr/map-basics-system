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
	SystemSettings SystemSettings
	RenderContext  *system.RenderContext
	GameCamera     *system.GameCamera
}

func (game *Game) Init() {
	renderContext := system.CreateRenderContext(
		config.VirtualWidth,
		config.VirtualHeight,
		game.SystemSettings.ScreenSetting.ScreenSize,
	)
	game.RenderContext = &renderContext

}

func (game *Game) Unload() {
	rl.UnloadRenderTexture(*game.RenderContext.RenderTexture)
}

func CreateGame(windowedScreenSize, screenSize rl.Vector2) Game {
	//TODO: This is a hold over until real game loading happens
	camera := system.CreateGameCamera(rl.NewVector2(float32(512), float32(512)), rl.NewVector2(config.VirtualWidth/2, config.VirtualHeight/2), 0.0, 1.0)
	return Game{
		GameCamera: &camera,
		RenderContext: nil,
		SystemSettings: SystemSettings{
			ScreenSetting: system.CreateScreenSetting(screenSize, windowedScreenSize),
		},
	}
}

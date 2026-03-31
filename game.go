package main

import (
	"github.com/couryrr/map-basics-system/config"
	"github.com/couryrr/map-basics-system/entity/player"
	"github.com/couryrr/map-basics-system/framework/queue"
	framework "github.com/couryrr/map-basics-system/framework/ui"
	"github.com/couryrr/map-basics-system/system/camera"
	"github.com/couryrr/map-basics-system/system/renderer"
	"github.com/couryrr/map-basics-system/system/setting"
	"github.com/couryrr/map-basics-system/system/ui"
	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SystemSettings struct {
	ScreenSetting setting.ScreenSetting
}

type Game struct {
	EventQueue         *queue.EventQueue
	GameCamera     *camera.GameCamera
	RenderContext  *renderer.RenderContext
	Player         *player.Player
	World          *world.World
	Ui             *framework.Root
	SystemSettings SystemSettings
	IsFullScreen   bool
}

func (game *Game) LoadResources() {
	renderContext := renderer.NewRenderContext(
		config.VirtualWidth,
		config.VirtualHeight,
		game.SystemSettings.ScreenSetting.ScreenSize,
	)
	game.RenderContext = &renderContext

	p1 := player.NewPlayer(rl.NewVector2(renderContext.VirtualWidth/2, renderContext.VirtualHeight/2))
	game.Player = &p1

	cam := camera.NewGameCamera(rl.NewVector2(p1.Position.X, p1.Position.Y), rl.NewVector2(float32(renderContext.VirtualWidth/2), float32(renderContext.VirtualHeight/2)), 0.0, 1.0)
	game.GameCamera = &cam

	root := ui.NewInGameOverlay(game)
	game.Ui = &root
}

func (game *Game) ToggleScreenSize(event *queue.Event) {
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

func (game *Game) Draw() {
	game.World.Draw(game.GameCamera.Camera.Target)
}

func (game *Game) LoadWorld() {
	w := world.NewWorld()
	game.World = &w
}

// TODO: should the game match the DrawState? Is there a better way to pass this data in?
func (game *Game) GetRenderContext() *renderer.RenderContext { return game.RenderContext }
func (game *Game) GetHotbarState() ui.HotbarState {
	return &game.Player.Hotbar
}

func (game *Game) GetRegistryState() ui.RegistryState {
	return game.World.Registry
}

func (game *Game) Unload() {
	game.World.UnloadWorld()
	rl.UnloadRenderTexture(*game.RenderContext.RenderTexture)
}

func NewGame(windowedScreenSize, screenSize rl.Vector2) Game {
	// TODO: This will load setting from files.
	broker := queue.NewEventQueue()
	return Game{
		EventQueue:        &broker,
		Player:        nil,
		GameCamera:    nil,
		RenderContext: nil,
		SystemSettings: SystemSettings{
			ScreenSetting: setting.NewScreenSetting(screenSize, windowedScreenSize),
		},
	}
}

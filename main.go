package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ScreenSetting struct {
	IsFullScreen bool
	destWidth    int32
	destHeight   int32
	destX        int32
	destY        int32
}

func CreateScreenSetting() ScreenSetting {
	scaleX := ScreenWidth / VirtualWidth
	scaleY := ScreenHeight / VirtualHeight
	scale := min(scaleX, scaleY)

	destWidth := VirtualWidth * scale
	destHeight := VirtualHeight * scale
	destX := (ScreenWidth - destWidth) / 2
	destY := (ScreenHeight - destHeight) / 2
	return ScreenSetting{
		IsFullScreen: false,
		destX:        destX,
		destY:        destY,
		destWidth:    destWidth,
		destHeight:   destHeight,
	}
}

func (ss *ScreenSetting) HandleScreenToggle() {
	rl.ToggleFullscreen()
	if ss.IsFullScreen {
		ScreenWidth = DefaultScreenWidth
		ScreenHeight = DefaultScreenHeight
	} else {
		ScreenWidth = int32(rl.GetScreenWidth())
		ScreenHeight = int32(rl.GetScreenHeight())
	}
	ss.IsFullScreen = !ss.IsFullScreen
	ss.CalculateViewport()
}

func (ss *ScreenSetting) CalculateViewport() {
	scaleX := ScreenWidth / VirtualWidth
	scaleY := ScreenHeight / VirtualHeight
	scale := min(scaleX, scaleY)

	destWidth := VirtualWidth * scale
	destHeight := VirtualHeight * scale
	destX := (ScreenWidth - destWidth) / 2
	destY := (ScreenHeight - destHeight) / 2

	ss.destWidth = destWidth
	ss.destHeight = destHeight
	ss.destX = destX
	ss.destY = destY
}

type TerrainLevel struct {
	Threshold float32
	Color     rl.Color
}

const (
	DefaultScreenWidth  int32 = 1920
	DefaultScreenHeight int32 = 1080
	VirtualWidth        int32 = 320
	VirtualHeight       int32 = 180
)

var (
	ScreenWidth  int32 = DefaultScreenWidth
	ScreenHeight int32 = DefaultScreenHeight

	tileSize      int32   = 2
	playerSize    int32   = 2
	playerSpeed   float32 = 100.0
	TerrainColors         = []TerrainLevel{
		{0.30, rl.NewColor(10, 50, 100, 255)},   // Deep Water
		{0.40, rl.NewColor(30, 100, 160, 255)},  // Shallow Water
		{0.45, rl.NewColor(210, 190, 140, 255)}, // Sand
		{0.65, rl.NewColor(90, 160, 70, 255)},   // Grass
		{0.80, rl.NewColor(40, 100, 40, 255)},   // Forest
		{0.90, rl.NewColor(130, 120, 110, 255)}, // Mountain
		{1.00, rl.NewColor(220, 225, 230, 255)}, // Snow Cap
	}
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Map Basics")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	perlin := rl.GenImagePerlinNoise(int(ScreenWidth)*2, int(ScreenHeight)*2, 0, 0, 10)
	defer rl.UnloadImage(perlin)

	perlinTexture := rl.LoadTextureFromImage(perlin)
	defer rl.UnloadTexture(perlinTexture)

	texture := rl.LoadRenderTexture(VirtualWidth, VirtualHeight)
	defer rl.UnloadRenderTexture(texture)

	player := rl.NewRectangle(float32(ScreenWidth/2), float32(ScreenHeight/2), float32(playerSize), float32(playerSize))
	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(player.X), float32(player.Y))
	camera.Offset = rl.NewVector2(float32(VirtualWidth/2), float32(VirtualHeight/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	screenSetting := CreateScreenSetting()

	for !rl.WindowShouldClose() {
		HandleInput(&player, &screenSetting)

		rl.BeginTextureMode(texture)
		rl.ClearBackground(rl.White)
		camera.Target = rl.NewVector2(player.X+player.Width/2, player.Y+player.Height/2)
		rl.BeginMode2D(camera)
		rl.DrawTexture(perlinTexture, 0, 0, rl.White)
		for x := int32(0); x < ScreenWidth; x += tileSize {
			for y := int32(0); y < ScreenHeight; y += tileSize {
				rl.DrawRectangle(x, y, tileSize, tileSize, DetermineTile(x, y, perlin))
			}
		}
		rl.DrawRectangleRec(player, rl.Red)

		rl.EndMode2D()
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		source := rl.NewRectangle(0, 0, float32(VirtualWidth), float32(-VirtualHeight))
		dest := rl.NewRectangle(float32(screenSetting.destX), float32(screenSetting.destY), float32(screenSetting.destWidth), float32(screenSetting.destHeight))
		rl.DrawTexturePro(texture.Texture, source, dest, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

func HandleInput(player *rl.Rectangle, screenSetting *ScreenSetting) {
	if rl.IsKeyDown(rl.KeyF11) {
		screenSetting.HandleScreenToggle()
	}
	delta := rl.GetFrameTime()

	dx, dy := float32(0), float32(0)

	if rl.IsKeyDown(rl.KeyW) {
		dy -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		dy += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		dx -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		dx += 1
	}

	if dx != 0 && dy != 0 {
		dx *= 0.7071
		dy *= 0.7071
	}

	player.X += dx * playerSpeed * delta
	player.Y += dy * playerSpeed * delta
}
func DetermineTile(x, y int32, perlin *rl.Image) color.RGBA {
	n := float32(rl.GetImageColor(*perlin, x, y).R) / 255
	for _, t := range TerrainColors {
		if n <= t.Threshold {
			return t.Color
		}
	}
	return rl.Black
}

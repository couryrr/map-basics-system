package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultScreenWidth  float32 = 1920
	DefaultScreenHeight float32 = 1080
	VirtualWidth        float32 = 320
	VirtualHeight       float32 = 180
)

var (
	ScreenWidth  float32 = DefaultScreenWidth
	ScreenHeight float32 = DefaultScreenHeight

	tileSize      float32 = 2
	playerSize    float32 = 2
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

	currentMark = 0
	marks       = [6]rl.Vector2{}
)

type Point interface {
	GetSize() rl.Vector2
	GetPostition() rl.Vector2
}

type Player struct {
	Size     rl.Vector2
	Position rl.Vector2
}

func (p *Player) GetSize() rl.Vector2 {
	return p.Size
}

func (p *Player) GetPostition() rl.Vector2 {
	return p.Position
}

type ScreenSetting struct {
	IsFullScreen bool
	scale        float32
	destWidth    float32
	destHeight   float32
	destX        float32
	destY        float32
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
		scale:        scale,
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
		ScreenWidth = float32(rl.GetScreenWidth())
		ScreenHeight = float32(rl.GetScreenHeight())
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

	ss.scale = scale
	ss.destWidth = destWidth
	ss.destHeight = destHeight
	ss.destX = destX
	ss.destY = destY
}

type GameCamera struct {
	Camera *rl.Camera2D
}

func CreateGameCamera(target Point, offSet rl.Vector2, rotation float32, zoom float32) GameCamera {
	return GameCamera{
		Camera: &rl.Camera2D{
			Target:   target.GetPostition(),
			Offset:   offSet,
			Rotation: rotation,
			Zoom:     zoom,
		},
	}
}

type TerrainLevel struct {
	Threshold float32
	Color     rl.Color
}

func main() {
	rl.InitWindow(int32(ScreenWidth), int32(ScreenHeight), "Map Basics")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	perlin := rl.GenImagePerlinNoise(int(ScreenWidth)*2, int(ScreenHeight)*2, 0, 0, 10)
	defer rl.UnloadImage(perlin)

	texture := rl.LoadRenderTexture(int32(VirtualWidth), int32(VirtualHeight))
	defer rl.UnloadRenderTexture(texture)

	player := Player{
		Size:     rl.NewVector2(playerSize, playerSize),
		Position: rl.NewVector2(ScreenWidth/2, ScreenHeight/2),
	}

	gameCamera := CreateGameCamera(&player, rl.NewVector2(VirtualWidth/2, VirtualHeight/2), 0.0, 1.0)
	screenSetting := CreateScreenSetting()

	for !rl.WindowShouldClose() {
		HandleInput(&player, &screenSetting, &gameCamera)

		rl.BeginTextureMode(texture)
		rl.ClearBackground(rl.White)
		gameCamera.Camera.Target = rl.NewVector2(player.Position.X+player.Size.X/2, player.Position.Y+player.Size.Y/2)
		rl.BeginMode2D(*gameCamera.Camera)
		for x := int32(0); x < int32(ScreenWidth); x += int32(tileSize) {
			for y := int32(0); y < int32(ScreenHeight); y += int32(tileSize) {
				rl.DrawRectangle(x, y, int32(tileSize), int32(tileSize), DetermineTile(x, y, perlin))
			}
		}
		for _, mark := range marks {
			rl.DrawRectangleLinesEx(rl.NewRectangle(mark.X, mark.Y, float32(tileSize), float32(tileSize)), 1, rl.White)
		}

		rl.DrawRectangleRec(rl.NewRectangle(player.Position.X, player.Position.Y, player.Size.X, player.Size.Y), rl.Red)

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

func HandleInput(player *Player, screenSetting *ScreenSetting, gameCamera *GameCamera) {
	if rl.IsKeyPressed(rl.KeyF11) {
		screenSetting.HandleScreenToggle()
	}
	if rl.IsKeyPressed(rl.KeyE) {
		gameCamera.Camera.Rotation += 90
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		gameCamera.Camera.Rotation -= 90
	}
	if rl.IsKeyPressed(rl.KeyC) {
		gameCamera.Camera.Rotation = 0
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mark := rl.GetMousePosition()
		mark.X = (mark.X - float32(screenSetting.destX)) / float32(screenSetting.scale)
		mark.Y = (mark.Y - float32(screenSetting.destY)) / float32(screenSetting.scale)
		marks[currentMark] = rl.GetScreenToWorld2D(mark, *gameCamera.Camera)
		currentMark = (currentMark + 1) % 6
	}
	for key := rl.KeyOne; key <= rl.KeySix; key++ {
		if rl.IsKeyPressed(int32(key)) {
			selected := int(key - rl.KeyOne)
			mark := marks[selected]
			player.Position.X = mark.X
			player.Position.Y = mark.Y
		}
	}
	if rl.IsKeyPressed(rl.KeyOne) {
		mark := marks[0]
		player.Position.Y = mark.Y
		player.Position.X = mark.X
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
	angle := -gameCamera.Camera.Rotation * rl.Deg2rad
	movement := rl.NewVector2(dx, dy)
	rotated := rl.Vector2Rotate(movement, angle)

	player.Position.X += rotated.X * playerSpeed * delta
	player.Position.Y += rotated.Y * playerSpeed * delta
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

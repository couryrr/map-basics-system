package main

import (
	"math"
	"os"

	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowedScreenWidth  float32 = 1920
	WindowededScreenHeight float32 = 1080
	VirtualWidth          float32 = 320
	VirtualHeight         float32 = 180
)

var (
	ScreenHeight   float32 = WindowededScreenHeight
	ScreenWidth    float32 = WindowedScreenWidth
	chunksToRender float32 = 3
	chunkSize      float32 = 15
	chunkWorldSize float32 = chunkSize * tileSize
	tileSize       float32 = 5
)

const perlinPath = "perlin.png"

func GetPerlin(size rl.Vector2) *rl.Image {
	if _, err := os.Stat(perlinPath); os.IsNotExist(err) {
		perlin := rl.GenImagePerlinNoise(int(size.X)*5, int(size.Y)*5, 0, 0, 4.0)
		rl.ExportImage(*perlin, perlinPath)
		return perlin
	}
	return rl.LoadImage(perlinPath)
}

func main() {
	screenSetting := system.CreateScreenSetting(rl.NewVector2(ScreenWidth, ScreenHeight), rl.NewVector2(VirtualWidth, VirtualHeight), rl.NewVector2(WindowedScreenWidth, WindowededScreenHeight))

	rl.InitWindow(int32(screenSetting.ScreenSize.X), int32(screenSetting.ScreenSize.Y), "Map Basics")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	perlin := GetPerlin(screenSetting.ScreenSize)
	defer rl.UnloadImage(perlin)

	world := system.CreateWorld(tileSize, perlin)

	texture := rl.LoadRenderTexture(int32(screenSetting.VirtualScreenSize.X), int32(screenSetting.VirtualScreenSize.Y))
	defer rl.UnloadRenderTexture(texture)

	gameCamera := system.CreateGameCamera(rl.Vector2Scale(screenSetting.ScreenSize, 0.5), rl.Vector2Scale(screenSetting.VirtualScreenSize, 0.5), 0.0, 1.0)

	for !rl.WindowShouldClose() {
		system.HandleInput(&screenSetting, &gameCamera)
		target := gameCamera.Camera.Target
		rl.BeginTextureMode(texture)
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(*gameCamera.Camera)
		chunkX := float32(math.Floor(float64(target.X / chunkWorldSize)))
		chunkY := float32(math.Floor(float64(target.Y / chunkWorldSize)))
		for dx := float32(-chunksToRender); dx <= chunksToRender; dx++ {
			for dy := float32(-chunksToRender); dy <= chunksToRender; dy++ {
				for x := float32(0); x < chunkSize; x++ {
					for y := float32(0); y < chunkSize; y++ {
						worldX := (chunkX+dx)*chunkWorldSize + x*tileSize
						worldY := (chunkY+dy)*chunkWorldSize + y*tileSize
						rl.DrawRectangleRec(rl.NewRectangle(worldX, worldY, tileSize, tileSize), world.DetermineTile(worldX, worldY, perlin))
					}
				}
				worldRX := float32(chunkX+dx) * chunkWorldSize
				worldRY := float32(chunkY+dy) * chunkWorldSize
				rl.DrawRectangleLinesEx(rl.NewRectangle(worldRX, worldRY, chunkWorldSize, chunkWorldSize), 1, rl.Green)
			}
		}

		// TODO: Fix these loops seems unneeded... guard mix/max of width/height...
		// for x := float32(0); x < ScreenWidth; x += tileSize {
		// 	rl.DrawLine(int32(x), int32(0), int32(x), int32(ScreenHeight), rl.Black)
		// }
		// for y := float32(0); y < ScreenHeight; y += tileSize {
		// 	rl.DrawLine(int32(0), int32(y), int32(ScreenWidth), int32(y), rl.Black)
		// }

		for _, mark := range system.Marks {
			rl.DrawRectangleLinesEx(rl.NewRectangle(mark.X, mark.Y, float32(tileSize), float32(tileSize)), 1, rl.White)
		}

		rl.EndMode2D()
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		source := rl.NewRectangle(0, 0, screenSetting.VirtualScreenSize.X, -screenSetting.VirtualScreenSize.Y)
		dest := rl.NewRectangle(screenSetting.DestinationPosition.X, screenSetting.DestinationPosition.Y, screenSetting.DestinationSize.X, screenSetting.DestinationSize.Y)
		rl.DrawTexturePro(texture.Texture, source, dest, rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}

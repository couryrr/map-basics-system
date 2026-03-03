package system

import (
	"image/color"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	worldSize      rl.Vector2 = rl.NewVector2(1920*5, 1080*5)
	tileSize       float32    = 5
	chunksToRender float32    = 3
	chunkSize      float32    = 15
	chunkWorldSize float32    = chunkSize * tileSize
	terrainColors             = [7]TerrainLevel{
		{0.30, rl.NewColor(10, 50, 100, 255)},   // Deep Water
		{0.40, rl.NewColor(30, 100, 160, 255)},  // Shallow Water
		{0.45, rl.NewColor(210, 190, 140, 255)}, // Sand
		{0.65, rl.NewColor(90, 160, 70, 255)},   // Grass
		{0.80, rl.NewColor(40, 100, 40, 255)},   // Forest
		{0.90, rl.NewColor(130, 120, 110, 255)}, // Mountain
		{1.00, rl.NewColor(220, 225, 230, 255)}, // Snow Cap
	}
)

type TerrainLevel struct {
	Threshold float32
	Color     rl.Color
}

type World struct {
	perlin        *rl.Image
	RenderTexture *rl.RenderTexture2D
	tileSize      float32
	Camera        GameCamera
}

func (w *World) DetermineTile(x, y float32, perlin *rl.Image) color.RGBA {
	px := int32(x/w.tileSize) % perlin.Width
	py := int32(y/w.tileSize) % perlin.Height
	if px < 0 {
		px += perlin.Width
	}
	if py < 0 {
		py += perlin.Height
	}
	n := float32(rl.GetImageColor(*perlin, px, py).R) / 255.0
	for _, t := range terrainColors {
		if n <= t.Threshold {
			return t.Color
		}
	}
	return rl.Black
}

func (w *World) Draw() {
	target := w.Camera.Camera.Target
	rl.BeginTextureMode(*w.RenderTexture)
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(*w.Camera.Camera)
	chunkX := float32(math.Floor(float64(target.X / chunkWorldSize)))
	chunkY := float32(math.Floor(float64(target.Y / chunkWorldSize)))
	for dx := float32(-chunksToRender); dx <= chunksToRender; dx++ {
		for dy := float32(-chunksToRender); dy <= chunksToRender; dy++ {
			for x := float32(0); x < chunkSize; x++ {
				for y := float32(0); y < chunkSize; y++ {
					worldX := (chunkX+dx)*chunkWorldSize + x*tileSize
					worldY := (chunkY+dy)*chunkWorldSize + y*tileSize
					rl.DrawRectangleRec(rl.NewRectangle(worldX, worldY, tileSize, tileSize), w.DetermineTile(worldX, worldY, w.perlin))
				}
			}
			worldRX := float32(chunkX+dx) * chunkWorldSize
			worldRY := float32(chunkY+dy) * chunkWorldSize
			rl.DrawRectangleLinesEx(rl.NewRectangle(worldRX, worldRY, chunkWorldSize, chunkWorldSize), 1, rl.Green)
		}
	}
	rl.EndMode2D()
	rl.EndTextureMode()

}

func (w *World) UnloadWorld(){
	rl.UnloadImage(w.perlin)
	rl.UnloadRenderTexture(*w.RenderTexture)
}

func GetPerlin(size rl.Vector2) *rl.Image {
	perlinPath := "perlin.png"
	if _, err := os.Stat(perlinPath); os.IsNotExist(err) {
		perlin := rl.GenImagePerlinNoise(int(size.X)*5, int(size.Y)*5, 0, 0, 4.0)
		rl.ExportImage(*perlin, perlinPath)
		return perlin
	}
	return rl.LoadImage(perlinPath)
}

func CreateWorld(camera GameCamera, texture rl.RenderTexture2D) World {
	perlin := GetPerlin(worldSize)
	defer rl.UnloadImage(perlin)
	return World{
		perlin:        GetPerlin(worldSize),
		RenderTexture: &texture,
		tileSize:      tileSize,
		Camera:        camera,
	}
}


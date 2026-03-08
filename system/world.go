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
	Camera          GameCamera
	WorldScreenSize rl.Vector2
	RenderTexture   *rl.RenderTexture2D
	perlin          *rl.Image
	tileSize        float32
}

func (w *World) DetermineTile(x, y float32) color.RGBA {
	px := int32(x/w.tileSize) % w.perlin.Width
	py := int32(y/w.tileSize) % w.perlin.Height
	if px < 0 {
		px += w.perlin.Width
	}
	if py < 0 {
		py += w.perlin.Height
	}
	n := float32(rl.GetImageColor(*w.perlin, px, py).R) / 255.0
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
	// Where am I to start in the chunks
	chunkX := float32(math.Floor(float64(target.X / chunkWorldSize)))
	chunkY := float32(math.Floor(float64(target.Y / chunkWorldSize)))
	for dx := float32(-chunksToRender); dx <= chunksToRender; dx++ {
		for dy := float32(-chunksToRender); dy <= chunksToRender; dy++ {
			// at my target chunk go out - and + and draw
			worldRX := float32(chunkX+dx) * chunkWorldSize
			worldRY := float32(chunkY+dy) * chunkWorldSize
			for x := range int32(chunkSize) {
				for y := range int32(chunkSize) {
					worldX := worldRX + float32(x)*tileSize
					worldY := worldRY + float32(y)*tileSize
					rl.DrawRectangleRec(rl.NewRectangle(worldX, worldY, tileSize, tileSize), w.DetermineTile(worldX, worldY))
				}
			}
			rl.DrawRectangleLinesEx(rl.NewRectangle(worldRX, worldRY, chunkWorldSize, chunkWorldSize), 1, rl.Green)
		}
	}
	rl.DrawRectangleLinesEx(rl.NewRectangle(chunkX*chunkWorldSize, chunkY*chunkWorldSize, chunkWorldSize, chunkWorldSize), 1, rl.Red)
	rl.EndMode2D()
	rl.EndTextureMode()

}

func (w *World) UnloadWorld() {
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

func CreateWorld(worldScreenSize rl.Vector2) World {
	perlin := GetPerlin(worldSize)
	texture := rl.LoadRenderTexture(int32(worldScreenSize.X), int32(worldScreenSize.Y))
	camera := CreateGameCamera(rl.NewVector2(float32(perlin.Width)/2, float32(perlin.Height)/2), rl.Vector2Scale(worldScreenSize, 0.5), 0.0, 1.0)
	return World{
		Camera:          camera,
		WorldScreenSize: worldScreenSize,
		RenderTexture:   &texture,
		perlin:          perlin,
		tileSize:        tileSize,
	}
}

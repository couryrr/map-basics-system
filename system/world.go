package system

import (
	"image/color"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	tileSize       float32    = 32
	chunkSize      float32    = 32
	chunksToRender float32    = 5
	chunkWorldSize float32    = chunkSize * tileSize
	terrainColors             = [7]TerrainLevel{
		{0.10, rl.NewColor(26, 16, 8, 255)},     // Deep Crater
		{0.25, rl.NewColor(58, 35, 18, 255)},    // Low Basin
		{0.45, rl.NewColor(122, 68, 32, 255)},   // Dark Oxide Flat
		{0.65, rl.NewColor(178, 95, 45, 255)},   // Rust Plain
		{0.78, rl.NewColor(196, 148, 88, 255)},  // Scoured Slope
		{0.88, rl.NewColor(200, 176, 140, 255)}, // Silicate Rock
		{1.00, rl.NewColor(216, 203, 178, 255)}, // Bleached Summit
	}
	// terrainColors             = [7]TerrainLevel{
	// 	{0.30, rl.NewColor(10, 50, 100, 255)},   // Deep Water
	// 	{0.40, rl.NewColor(30, 100, 160, 255)},  // Shallow Water
	// 	{0.45, rl.NewColor(210, 190, 140, 255)}, // Sand
	// 	{0.65, rl.NewColor(90, 160, 70, 255)},   // Grass
	// 	{0.80, rl.NewColor(40, 100, 40, 255)},   // Forest
	// 	{0.90, rl.NewColor(130, 120, 110, 255)}, // Mountain
	// 	{1.00, rl.NewColor(220, 225, 230, 255)}, // Snow Cap
	// }
)

type TerrainLevel struct {
	Threshold float32
	Color     rl.Color
}

type World struct {
	perlin1         *rl.Image
	perlin2         *rl.Image
	perlin3         *rl.Image
	tileSize        float32
}

func (w *World) sampleFBM(worldX, worldY float32) float32 {
	v1 := samplePerlin(worldX, worldY, w.perlin1)

	rx2, ry2 := rotatePoint(worldX, worldY, 45)
	v2 := samplePerlin(rx2, ry2, w.perlin2)

	rx3, ry3 := rotatePoint(worldX, worldY, 90)
	v3 := samplePerlin(rx3, ry3, w.perlin3)

	return ((v1 * 0.5) + (v2 * 0.25) + (v3 * 0.125)) / 0.875
}

func rotatePoint(x, y, angle float32) (float32, float32) {
	rad := angle * rl.Deg2rad
	cos := float32(math.Cos(float64(rad)))
	sin := float32(math.Sin(float64(rad)))
	return x*cos - y*sin, x*sin + y*cos
}

func samplePerlin(worldX, worldY float32, img *rl.Image) float32 {
	px := int32(math.Floor(float64(worldX/tileSize))) % img.Width
	py := int32(math.Floor(float64(worldY/tileSize))) % img.Height
	if px < 0 {
		px += img.Width
	}
	if py < 0 {
		py += img.Height
	}
	return float32(rl.GetImageColor(*img, px, py).R) / 255.0
}

func (w *World) DetermineTile(x, y float32) color.RGBA {
	n := w.sampleFBM(x, y)
	for _, t := range terrainColors {
		if n <= t.Threshold {
			return t.Color
		}
	}
	return rl.Black
}

func (w *World) Draw(target rl.Vector2) {
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
					rl.DrawRectangleLines(int32(worldX), int32(worldY), int32(tileSize), int32(tileSize), rl.Gray)
				}
			}
			rl.DrawRectangleLinesEx(rl.NewRectangle(worldRX, worldRY, chunkWorldSize, chunkWorldSize), 1, rl.Green)
		}
	}
	rl.DrawRectangleLinesEx(rl.NewRectangle(chunkX*chunkWorldSize, chunkY*chunkWorldSize, chunkWorldSize, chunkWorldSize), 1, rl.Red)
}

func (w *World) UnloadWorld() {
	rl.UnloadImage(w.perlin1)
	rl.UnloadImage(w.perlin2)
	rl.UnloadImage(w.perlin3)
}

func GetPerlin() (*rl.Image, *rl.Image, *rl.Image) {
	perlinWidth := 1024
	perlinHeight := 1024

	paths := []string{"perlin1.png", "perlin2.png", "perlin3.png"}
	params := [][2]int{{0, 0}, {143, 143}, {313, 313}}
	scales := []float32{3.0, 6.0, 12.0}

	for i, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			img := rl.GenImagePerlinNoise(perlinWidth, perlinHeight, params[i][0], params[i][1], scales[i])
			rl.ExportImage(*img, path)
			rl.UnloadImage(img)
		}
	}
	return rl.LoadImage(paths[0]), rl.LoadImage(paths[1]), rl.LoadImage(paths[2])
}

func CreateWorld() World {
	perlin1, perlin2, perlin3 := GetPerlin()
	return World{
		perlin1:         perlin1,
		perlin2:         perlin2,
		perlin3:         perlin3,
		tileSize:        tileSize,
	}
}

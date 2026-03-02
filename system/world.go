package system

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	terrainColors = [7]TerrainLevel{
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
	perlin   *rl.Image
	tileSize float32
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
func CreateWorld(tileSize float32, noise *rl.Image) World {
	return World{
		perlin:   noise,
		tileSize: tileSize,
	}
}

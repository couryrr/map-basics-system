package system

import rl "github.com/gen2brain/raylib-go/raylib"

type Item struct {
	Position rl.Vector2
	Size     rl.Vector2
}

const (
	zoomStep = float32(0.25)
	zoomMin  = float32(0.25)
	zoomMax  = float32(3.0)
)

type Player struct {
	Position  rl.Vector2
	Rotation  float32
	Speed     float32
	ZoomLevel float32
}

func CreatePlayer(start rl.Vector2) Player {
	return Player{
		Position:  start,
		Rotation:  0,
		Speed:     400,
		ZoomLevel: 1.0,
	}
}

func (player *Player) Rotate(message Message) {
	if rotation, ok := message.Data.(float32); ok {
		player.Rotation += rotation
	}
}

func (player *Player) RotateReset(message Message) {
	player.Rotation = 0
}

func (player *Player) Zoom(message Message) {
	if delta, ok := message.Data.(float32); ok {
		player.ZoomLevel += delta * zoomStep
		player.ZoomLevel = rl.Clamp(player.ZoomLevel, zoomMin, zoomMax)
	}
}

func (player *Player) Move(message Message) {
	if movement, ok := message.Data.(rl.Vector2); ok {
		delta := rl.GetFrameTime()
		angle := -player.Rotation * rl.Deg2rad
		rotated := rl.Vector2Rotate(movement, angle)
		player.Position.X += rotated.X * player.Speed * delta
		player.Position.Y += rotated.Y * player.Speed * delta
	}

}

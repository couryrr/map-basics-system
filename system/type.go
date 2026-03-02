package system

import rl "github.com/gen2brain/raylib-go/raylib"

type Point interface {
	GetSize() rl.Vector2
	GetPostition() rl.Vector2
}

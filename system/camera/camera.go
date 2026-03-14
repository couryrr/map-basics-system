package camera

import rl "github.com/gen2brain/raylib-go/raylib"

type GameCamera struct {
	Camera     *rl.Camera2D
}

func NewGameCamera(target rl.Vector2, offSet rl.Vector2, rotation float32, zoom float32) GameCamera {
	return GameCamera{
		Camera: &rl.Camera2D{
			Target:   target,
			Offset:   offSet,
			Rotation: rotation,
			Zoom:     zoom,
		},
	}
}

func (gc *GameCamera) Update(){
	
}

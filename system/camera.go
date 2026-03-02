package system

import rl "github.com/gen2brain/raylib-go/raylib"

type GameCamerMode int

const (
	CameraModePlanning GameCamerMode = iota
	CameraModeBuild
)

var GameCameraName = map[GameCamerMode]string{
	CameraModePlanning: "planning",
	CameraModeBuild:   "build",
}

type GameCamera struct {
	CameraMode GameCamerMode
	Camera     *rl.Camera2D
}

func CreateGameCamera(target rl.Vector2, offSet rl.Vector2, rotation float32, zoom float32) GameCamera {
	return GameCamera{
		CameraMode: CameraModeBuild,
		Camera: &rl.Camera2D{
			Target:   target,
			Offset:   offSet,
			Rotation: rotation,
			Zoom:     zoom,
		},
	}
}

func (gc *GameCamera) ChangeMode(mode GameCamerMode) {
	gc.CameraMode = mode
}

func (gc *GameCamera) GetPostition() rl.Vector2 {
	return gc.Camera.Target
}

package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// FIXME: Just making some lazy decisions here.
var (
	cameraSpeed float32 = 400.0
)

func HandleInput(screenSetting *ScreenSetting, gameCamera *GameCamera, rCtx *RenderContext) {
	if rl.IsKeyPressed(rl.KeyF11) {
		screenSetting.ToggleScreenSize()
		rCtx.Update(screenSetting.ScreenSize)

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
	if rl.IsKeyPressed(rl.KeyTab) {
		switch gameCamera.CameraMode {
		case CameraModePlanning:
			gameCamera.ChangeMode(CameraModeBuild)
			gameCamera.Camera.Zoom = 1.0
		case CameraModeBuild:
			gameCamera.ChangeMode(CameraModePlanning)
			gameCamera.Camera.Zoom = 0.5
		default:
			gameCamera.ChangeMode(CameraModeBuild)
			gameCamera.Camera.Zoom = 1.0
		}
	}

	// gameCamera.Camera.Zoom += rl.GetMouseWheelMove();

	target := gameCamera.Camera.Target
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

	target.X += rotated.X * cameraSpeed * delta
	target.Y += rotated.Y * cameraSpeed * delta

	gameCamera.Camera.Target = target
}

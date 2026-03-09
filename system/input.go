package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// FIXME: Just making some lazy decisions here.
var (
	cameraSpeed float32 = 400.0
)

func HandleInput(screenSetting *ScreenSetting, world *World) {
	if rl.IsKeyPressed(rl.KeyF11) {
		screenSetting.ToggleScreenSize(world.WorldScreenSize)
	}
	if rl.IsKeyPressed(rl.KeyE) {
		world.Camera.Camera.Rotation += 90
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		world.Camera.Camera.Rotation -= 90
	}
	if rl.IsKeyPressed(rl.KeyC) {
		world.Camera.Camera.Rotation = 0
	}
	if rl.IsKeyPressed(rl.KeyTab) {
		switch world.Camera.CameraMode {
		case CameraModePlanning:
			world.Camera.ChangeMode(CameraModeBuild)
			world.Camera.Camera.Zoom = 1.0
		case CameraModeBuild:
			world.Camera.ChangeMode(CameraModePlanning)
			world.Camera.Camera.Zoom = 0.5
		default:
			world.Camera.ChangeMode(CameraModeBuild)
			world.Camera.Camera.Zoom = 1.0
		}
	}

	// world.Camera.Camera.Zoom += rl.GetMouseWheelMove();

	target := world.Camera.Camera.Target
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
	angle := -world.Camera.Camera.Rotation * rl.Deg2rad
	movement := rl.NewVector2(dx, dy)
	rotated := rl.Vector2Rotate(movement, angle)

	target.X += rotated.X * cameraSpeed * delta
	target.Y += rotated.Y * cameraSpeed * delta

	world.Camera.Camera.Target = target
}

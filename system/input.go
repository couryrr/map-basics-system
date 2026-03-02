package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// FIXME: Just making some lazy decisions here.
var (
	Marks               = [6]rl.Vector2{}
	currentMark         = 0
	playerSpeed float32 = 100.0
)

func HandleInput(screenSetting *ScreenSetting, gameCamera *GameCamera) {
	if rl.IsKeyPressed(rl.KeyF11) {
		screenSetting.ToggleScreenSize()
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

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mark := rl.GetMousePosition()
		mark.X = (mark.X - float32(screenSetting.DestinationPosition.X)) / float32(screenSetting.scale)
		mark.Y = (mark.Y - float32(screenSetting.DestinationPosition.Y)) / float32(screenSetting.scale)
		Marks[currentMark] = rl.GetScreenToWorld2D(mark, *gameCamera.Camera)
		currentMark = (currentMark + 1) % 6
	}

	target := gameCamera.Camera.Target
	for key := rl.KeyOne; key <= rl.KeySix; key++ {
		if rl.IsKeyPressed(int32(key)) {
			selected := int(key - rl.KeyOne)
			mark := Marks[selected]
			target.X = mark.X
			target.Y = mark.Y
		}
	}

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

	target.X += rotated.X * playerSpeed * delta
	target.Y += rotated.Y * playerSpeed * delta

	gameCamera.Camera.Target = target
}

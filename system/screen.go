package system

import rl "github.com/gen2brain/raylib-go/raylib"

type ScreenSetting struct {
	IsFullScreen       bool
	scale              float32
	WindowedScreenSize rl.Vector2
	ScreenSize         rl.Vector2
}

func CreateScreenSetting(screenSize, windowedScreenSize rl.Vector2) ScreenSetting {
	return ScreenSetting{
		IsFullScreen:       false,
		ScreenSize:         screenSize,
		WindowedScreenSize: windowedScreenSize,
	}
}

func (ss *ScreenSetting) CalculateViewport(virtualScreenSize rl.Vector2) rl.Rectangle {
	scaleX := ss.ScreenSize.X / virtualScreenSize.X
	scaleY := ss.ScreenSize.Y / virtualScreenSize.Y
	scale := min(scaleX, scaleY)

	destWidth := virtualScreenSize.X * scale
	destHeight := virtualScreenSize.Y * scale
	destX := (ss.ScreenSize.X - destWidth) / 2
	destY := (ss.ScreenSize.Y - destHeight) / 2
	return rl.NewRectangle(destX, destY, destWidth, destHeight)
}

func (ss *ScreenSetting) ToggleScreenSize(virtualScreenSize rl.Vector2) {
	rl.ToggleFullscreen()
	newScreenSize := rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	if !ss.IsFullScreen {
		newScreenSize = ss.WindowedScreenSize
	}
	ss.ScreenSize = newScreenSize
	ss.IsFullScreen = !ss.IsFullScreen
}

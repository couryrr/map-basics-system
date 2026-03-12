package system

import rl "github.com/gen2brain/raylib-go/raylib"

type ScreenSetting struct {
	IsFullScreen       bool
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

func (ss *ScreenSetting) ToggleScreenSize() {
	rl.ToggleFullscreen()
	newScreenSize := rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	if !ss.IsFullScreen {
		newScreenSize = ss.WindowedScreenSize
	}
	ss.ScreenSize = newScreenSize
	ss.IsFullScreen = !ss.IsFullScreen
}

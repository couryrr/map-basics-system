package system

import rl "github.com/gen2brain/raylib-go/raylib"

type ScreenSetting struct {
	WindowedScreenSize rl.Vector2
	ScreenSize         rl.Vector2
}

func CreateScreenSetting(screenSize, windowedScreenSize rl.Vector2) ScreenSetting {
	return ScreenSetting{
		ScreenSize:         screenSize,
		WindowedScreenSize: windowedScreenSize,
	}
}


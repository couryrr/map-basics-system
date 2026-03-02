package system

import rl "github.com/gen2brain/raylib-go/raylib"

type ScreenSetting struct {
	IsFullScreen        bool
	scale               float32
	WindowedScreenSize  rl.Vector2
	DestinationSize     rl.Vector2
	DestinationPosition rl.Vector2
	ScreenSize          rl.Vector2
	VirtualScreenSize   rl.Vector2
}

func CreateScreenSetting(screenSize, virtualScreenSize, windowedScreenSize rl.Vector2) ScreenSetting {
	scaleX := screenSize.X / virtualScreenSize.X
	scaleY := screenSize.Y / virtualScreenSize.Y
	scale := min(scaleX, scaleY)

	destWidth := virtualScreenSize.X * scale
	destHeight := virtualScreenSize.Y * scale
	destX := (screenSize.X - destWidth) / 2
	destY := (screenSize.Y - destHeight) / 2
	return ScreenSetting{
		IsFullScreen:        false,
		scale:               scale,
		ScreenSize:          screenSize,
		WindowedScreenSize:  windowedScreenSize,
		VirtualScreenSize:   virtualScreenSize,
		DestinationSize:     rl.NewVector2(destWidth, destHeight),
		DestinationPosition: rl.NewVector2(destX, destY),
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
	ss.CalculateViewport()
}

func (ss *ScreenSetting) CalculateViewport() {
	scaleX := ss.ScreenSize.X / ss.VirtualScreenSize.X
	scaleY := ss.ScreenSize.Y / ss.VirtualScreenSize.Y
	scale := min(scaleX, scaleY)

	destWidth := ss.VirtualScreenSize.X * scale
	destHeight := ss.VirtualScreenSize.Y * scale
	destX := (ss.ScreenSize.X - destWidth) / 2
	destY := (ss.ScreenSize.Y - destHeight) / 2

	ss.scale = scale
	ss.DestinationSize = rl.NewVector2(destWidth, destHeight)
	ss.DestinationPosition = rl.NewVector2(destX, destY)
}

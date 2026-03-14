package renderer

import rl "github.com/gen2brain/raylib-go/raylib"

type RenderContext struct {
	VirtualWidth, VirtualHeight float32
	Viewport                    rl.Rectangle
	Scale                       float32

	RenderTexture *rl.RenderTexture2D
}

func (rc *RenderContext) Update(screenSize rl.Vector2) {
	scaleX := screenSize.X / rc.VirtualWidth
	scaleY := screenSize.Y / rc.VirtualHeight
	rc.Scale = min(scaleX, scaleY)

	destWidth := rc.VirtualWidth * rc.Scale
	destHeight := rc.VirtualHeight * rc.Scale
	destX := (screenSize.X - destWidth) / 2
	destY := (screenSize.Y - destHeight) / 2

	rc.Viewport = rl.NewRectangle(destX, destY, destWidth, destHeight)
}

func (rc *RenderContext) ScreenToVirtual(screenPos rl.Vector2) rl.Vector2 {
	return rl.NewVector2(
		(screenPos.X-rc.Viewport.X)/rc.Scale,
		(screenPos.Y-rc.Viewport.Y)/rc.Scale,
	)
}

func (rc *RenderContext) VirtualToScreen(virtualPos rl.Vector2) rl.Vector2 {
	return rl.NewVector2(
		virtualPos.X*rc.Scale+rc.Viewport.X,
		virtualPos.Y*rc.Scale+rc.Viewport.Y,
	)
}

func NewRenderContext(virtualWidth, virtualHeight float32, screenSize rl.Vector2) RenderContext {
	renderTexture := rl.LoadRenderTexture(int32(virtualWidth), int32(virtualHeight))

	ctx := RenderContext{
		VirtualWidth:  virtualWidth,
		VirtualHeight: virtualHeight,
		RenderTexture: &renderTexture,
	}
	ctx.Update(screenSize)
	return ctx
}

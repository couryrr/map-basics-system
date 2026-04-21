package ui

import (
	"github.com/couryrr/map-basics-system/framework/ui"
	"github.com/couryrr/map-basics-system/system/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InGameOverlayState interface {
	GetHotbarState() HotbarState
	GetRegistryState() RegistryState

	GetRenderContext() *renderer.RenderContext
}

type InGameOverlay struct {
	ui.Element
}

func NewInGameOverlay(state InGameOverlayState) *ui.Element {
	igo := ui.NewElement()
	igo.SetBounds(rl.NewRectangle(0, 0, state.GetRenderContext().VirtualWidth, state.GetRenderContext().VirtualHeight))
	igo.WithPropFn(func() ui.Prop {
		return ui.Prop{
			Style: ui.NewStyle().Border(1, rl.DarkGray).Build(),
		}
	})


	hotbar := NewHotbarElement(igo.Bounds(), state.GetHotbarState())
	registry := NewRegistryElement(igo.Bounds(), state.GetRegistryState())

	igo.AddChild(&hotbar)
	igo.AddChild(&registry)

	return &igo
}

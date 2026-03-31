package ui

import (
	"github.com/couryrr/map-basics-system/framework/queue"
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
	broker *queue.EventQueue
	ui.Element
}

func NewInGameOverlay(state InGameOverlayState) ui.Root {
	root := ui.NewRoot(rl.NewRectangle(0, 0, state.GetRenderContext().VirtualWidth, state.GetRenderContext().VirtualHeight))

	igo := ui.NewElement()
	igo.WithPropFn(func() ui.Prop {
		return ui.Prop{
			Style: ui.NewStyle().Border(1, rl.DarkGray).Build(),
		}
	})

	root.SetChild(&igo)

	hotbar := NewHotbarElement(igo.Bounds(), state.GetHotbarState())
	registry := NewRegistryElement(igo.Bounds(), state.GetRegistryState())

	igo.AddChild(&hotbar)
	igo.AddChild(&registry)

	return root
}

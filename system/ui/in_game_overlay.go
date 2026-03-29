package ui

import (
	"github.com/couryrr/map-basics-system/framework/pubsub"
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
	broker *pubsub.Broker
	framework.Element
}

func NewInGameOverlay(state InGameOverlayState) framework.Root {
	root := framework.NewRoot(rl.NewRectangle(0, 0, state.GetRenderContext().VirtualWidth, state.GetRenderContext().VirtualHeight))

	igo := framework.NewElement()
	igo.WithPropFn(func() framework.Prop {
		return framework.Prop{
			Style: framework.NewStyle().Border(1, rl.DarkGray).Build(),
		}
	})

	root.AddChild(&igo)

	hotbar := NewHotbarElement(igo.Bounds(), state.GetHotbarState())
	registry := NewRegistryElement(igo.Bounds(), state.GetRegistryState())

	igo.AddChild(&hotbar)
	igo.AddChild(&registry)

	return root
}

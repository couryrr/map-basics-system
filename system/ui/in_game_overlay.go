package ui

import (
	"github.com/couryrr/map-basics-system/framework/pubsub"
	"github.com/couryrr/map-basics-system/framework/ui"
	"github.com/couryrr/map-basics-system/system/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DrawContext interface {
	GetHotbarState() HotbarState
	GetRenderContext() *renderer.RenderContext
}

type InteractionResult struct {
	Topic   pubsub.Topic
	Message pubsub.Message
}

// TODO: Another state I am just not sure of...
type InGameOverlayState interface {
	GetHotbarState() HotbarState
	GetRegistryState() RegistryState
}

type InGameOverlay struct {
	broker *pubsub.Broker
	framework.Element
}

func (igo *InGameOverlay) HandleMouseEvent(messge pubsub.Message) {
	if event, ok := messge.Data.(framework.InputEvent); ok {
		for _, child := range igo.Children() {
			if rl.CheckCollisionPointRec(event.Position, child.Bounds()) {
				child.HandleEvents(event)
			} else {
				igo.broker.Send(TopicUiHotbarInteraction, pubsub.Message{
					Data: HotbarInteractionMessage{Action: HotbarActionLeave},
				})
			}
		}
	}
}

func NewInGameOverlay(broker *pubsub.Broker, rctx renderer.RenderContext, state InGameOverlayState) InGameOverlay {
	root := rl.NewRectangle(0, 0, rctx.VirtualWidth, rctx.VirtualHeight)
	igo := InGameOverlay{
		broker:  broker,
		Element: framework.NewRoot(root),
	}
	igo.WithStyle(framework.NewStyle().Border(1, rl.DarkGray).Build())

	hotbar := NewHotbarElement(igo.Bounds(), state.GetHotbarState())
	registry := NewRegistryElement(igo.Bounds(), state.GetRegistryState())

	igo.AddChild(&hotbar)
	igo.AddChild(&registry)
	return igo
}

package ui

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
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
	Container
}

func (igo *InGameOverlay) HandleMouseEvent(messge pubsub.Message) {
	if event, ok := messge.Data.(InputEvent); ok {
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

func (igo *InGameOverlay) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(igo.bounds, igo.Style.Border.Thickness, igo.Style.Border.Color)
	for _, child := range igo.Children() {
		child.Draw(ctx)
	}
}

func NewInGameOverlay(broker *pubsub.Broker, rctx renderer.RenderContext, state InGameOverlayState) InGameOverlay {
	root := rl.NewRectangle(0, 0, rctx.VirtualWidth, rctx.VirtualHeight)
	igo := InGameOverlay{
		broker:    broker,
		Container: NewContainer(root, WithBorder(1, rl.DarkGray)),
	}

	parentBounds := igo.Bounds()
	hotbar := NewHotbarElement(parentBounds, state.GetHotbarState())
	registry := NewRegistryElement(parentBounds, state.GetRegistryState())

	igo.AddChild(&hotbar)
	igo.AddChild(&registry)
	return igo
}

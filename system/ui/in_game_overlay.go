package ui

import (
	"iter"

	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/renderer"
	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DrawContext interface {
	GetHotbarState() HotbarState
	GetRenderContext() *renderer.RenderContext
	GetItemByIdFromRegistry(itemId string) (*world.GameItem, error)
	GetRegistryItems() iter.Seq2[string, world.GameItem]
}

type InteractionResult struct {
	Topic   pubsub.Topic
	Message pubsub.Message
}

type InGameOverlay struct {
	broker *pubsub.Broker
	Container
}

func (igo *InGameOverlay) HandleMouseEvent(messge pubsub.Message) {
	if event, ok := messge.Data.(MouseEvent); ok {
		for _, child := range igo.Children() {
			if rl.CheckCollisionPointRec(event.Position, child.Bounds()) {

			} else {
				igo.broker.Send(TopicUiHotbarInteraction, pubsub.Message{
					Data: HotbarInteractionMessage{Action: HotbarActionLeave},
				})
			}
		}
	}
}

func NewInGameOverlay(broker *pubsub.Broker, rctx renderer.RenderContext, items iter.Seq2[string, world.GameItem]) InGameOverlay {
	root := rl.NewRectangle(0, 0, rctx.VirtualWidth, rctx.VirtualHeight)
	igo := InGameOverlay{
		broker:    broker,
		Container: NewContainer(root, WithBorder(1, rl.DarkGray)),
	}

	parentBounds := igo.Bounds()
	hotbar := NewHotbarElement(parentBounds)
	registry := NewRegistryElement(parentBounds, items)

	igo.AddChild(&registry)
	igo.AddChild(&hotbar)
	return igo
}

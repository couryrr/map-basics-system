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
	broker    *pubsub.Broker
	container Container
}

func (igo *InGameOverlay) CheckIntersection(messge pubsub.Message) {
	if point, ok := messge.Data.(rl.Vector2); ok {
		_ = point
		// if rl.CheckCollisionPointRec(point, igo.hotbar.Bound) {
		// 	ir := igo.hotbar.HandleIntersection(point)
		// 	if ir != nil {
		// 		igo.broker.Send(ir.Topic, ir.Message)
		// 	}
		// } else {
		// 	igo.broker.Send(TopicUiHotbarInteraction, pubsub.Message{
		// 		Data: HotbarInteractionMessage{Action: HotbarActionLeave},
		// 	})
		// }
	}
}

func (igo *InGameOverlay) Draw(ctx DrawContext) {
	igo.container.Draw(ctx)
}

func NewInGameOverlay(broker *pubsub.Broker, rctx renderer.RenderContext) InGameOverlay {
	root := rl.NewRectangle(0, 0, rctx.VirtualWidth, rctx.VirtualHeight)
	igo := InGameOverlay{
		broker:    broker,
		container: NewContainer(root, LayoutNone, WithBorder(1, rl.DarkGray)),
	}

	parentBounds := igo.container.Bounds()
	hotbar := NewHotbarElement(parentBounds)
	registry := NewRegistryElement(parentBounds)

	igo.container.AddChild(registry)
	igo.container.AddChild(hotbar)
	return igo
}

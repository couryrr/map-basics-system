package ui

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/renderer"
	"github.com/couryrr/map-basics-system/system/resource"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DrawContext interface {
	GetHotbarState() HotbarState
	GetRenderContext() *renderer.RenderContext
	GetItemFromDirectory(itemId string) (*resource.GameItem, error)
}

type InteractionResult struct {
	Topic   pubsub.Topic
	Message pubsub.Message
}

type InGameOverlay struct {
	broker *pubsub.Broker
	hotbar HotbarElement
}

func (igo *InGameOverlay) CheckIntersection(messge pubsub.Message) {
	if point, ok := messge.Data.(rl.Vector2); ok {
		if rl.CheckCollisionPointRec(point, igo.hotbar.Bound) {
			ir := igo.hotbar.HandleIntersection(point)
			if ir != nil {
				igo.broker.Send(ir.Topic, ir.Message)
			}
		} else {
			igo.broker.Send(TopicUiHotbarInteraction, pubsub.Message{
				Data: HotbarInteractionMessage{Action: HotbarActionLeave},
			})
		}
	}
}

func (igo *InGameOverlay) Draw(ctx DrawContext) {
	igo.hotbar.Draw(ctx)
}

func NewInGameOverlay(broker *pubsub.Broker, rCtx *renderer.RenderContext) InGameOverlay {
	hotbar := NewHotbar(rCtx)

	return InGameOverlay{
		broker: broker,
		hotbar: hotbar,
	}
}

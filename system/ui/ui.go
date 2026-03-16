package ui

import (
	"fmt"

	"github.com/couryrr/map-basics-system/entity/player"
	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TopicUiHotbarElementHovered pubsub.Topic = "ui.hotbar.element.hovered"
)

type HotbarItemElement struct {
	Bound rl.Rectangle
}

type HotbarElement struct {
	Bound rl.Rectangle
	Slots [6]HotbarItemElement
}

func (hb *HotbarElement) Draw(player player.Player, rCtx *renderer.RenderContext) {
	rl.DrawRectangleLinesEx(hb.Bound, 1, rl.DarkGray)
	for i, slot := range hb.Slots {
		rl.DrawText(fmt.Sprintf("%s", player.Hotbar.Slots[i]), int32(slot.Bound.X+2), int32(slot.Bound.Y+2), 12, rl.DarkBlue)
		rl.DrawRectangleLinesEx(slot.Bound, float32(1), rl.DarkBlue)
		ii := int32(i)
		if player.Hotbar.ActiveSlot != nil {
			if *player.Hotbar.ActiveSlot == ii {
				rl.DrawRectangleLinesEx(slot.Bound, float32(5), rl.Red)
			}
		}
	}
}

type InGameOverlay struct {
	broker *pubsub.Broker
	hotbar HotbarElement
}

func (igo *InGameOverlay) CheckIntersection(messge pubsub.Message) {
	if point, ok := messge.Data.(rl.Vector2); ok {
		//TODO: this is just sloppy for now. Needs to be more propagated.
		for i := range igo.hotbar.Slots {
			if rl.CheckCollisionPointRec(point, igo.hotbar.Slots[i].Bound) {
				ii := int32(i)
				igo.broker.Send(TopicUiHotbarElementHovered, pubsub.Message{Data: &ii})
				break
			}
		}
	}
}

func (igo *InGameOverlay) Draw(player player.Player, rCtx *renderer.RenderContext) {
	igo.hotbar.Draw(player, rCtx)
}

func NewInGameOverlay(broker *pubsub.Broker, rCtx *renderer.RenderContext) InGameOverlay {
	posY := rCtx.VirtualHeight - 64
	sizeX := rCtx.VirtualWidth - 128

	//TODO: This should be on the HotbarElement
	slots := new([6]HotbarItemElement)
	for i := range 6 {
		pos := rl.NewVector2(float32(64+i*32), float32(posY))
		slots[i] = HotbarItemElement{
			Bound: rl.NewRectangle(pos.X, posY, 32, 32),
		}
	}
	hotbar := HotbarElement{
		Bound: rl.NewRectangle(64, float32(posY), float32(sizeX), 64),
		Slots: *slots,
	}

	return InGameOverlay{
		broker: broker,
		hotbar: hotbar,
	}
}

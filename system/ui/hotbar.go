package ui

import (
	"fmt"

	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HotbarAction string

const (
	TopicUiHotbarInteraction pubsub.Topic = "ui.hotbar.interaction"
	HotbarActionHover        HotbarAction = "hover"
	HotbarActionLeave        HotbarAction = "leave"
)

type HotbarInteractionMessage struct {
	Slot   int32
	ItemId string
	Action HotbarAction
}

type HotbarState interface {
	SlotItem(i int) string
	GetActiveSlot() *int32
}

type HotbarItemElement struct {
	Bound rl.Rectangle
}

type HotbarElement struct {
	Bound rl.Rectangle
	Slots [6]HotbarItemElement
}

func (hb HotbarElement) HandleIntersection(point rl.Vector2) *InteractionResult {
	for i := range hb.Slots {
		if rl.CheckCollisionPointRec(point, hb.Slots[i].Bound) {
			return &InteractionResult{
				Topic: TopicUiHotbarInteraction,
				Message: pubsub.Message{Data: HotbarInteractionMessage{
					Slot:   int32(i),
					Action: HotbarActionHover,
				}},
			}
		}
	}
	return nil
}

func NewHotbar(rCtx *renderer.RenderContext) HotbarElement {
	posY := rCtx.VirtualHeight - 64
	sizeX := rCtx.VirtualWidth - 128

	slots := new([6]HotbarItemElement)
	for i := range 6 {
		pos := rl.NewVector2(float32(64+i*32), float32(posY))
		slots[i] = HotbarItemElement{
			Bound: rl.NewRectangle(pos.X, posY, 32, 32),
		}
	}
	return HotbarElement{
		Bound: rl.NewRectangle(64, float32(posY), float32(sizeX), 64),
		Slots: *slots,
	}
}

func (hb *HotbarElement) Draw(state HotbarState, rCtx *renderer.RenderContext) {
	rl.DrawRectangleLinesEx(hb.Bound, 1, rl.DarkGray)
	for i, slot := range hb.Slots {
		rl.DrawText(fmt.Sprintf("%s", state.SlotItem(i)), int32(slot.Bound.X+2), int32(slot.Bound.Y+2), 12, rl.DarkBlue)
		rl.DrawRectangleLinesEx(slot.Bound, float32(1), rl.DarkBlue)
		ii := int32(i)
		if state.GetActiveSlot() != nil {
			if *state.GetActiveSlot() == ii {
				rl.DrawRectangleLinesEx(slot.Bound, float32(5), rl.Red)
			}
		}
	}
}

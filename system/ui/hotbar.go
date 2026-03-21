package ui

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
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
	SlotItem(i int32) string
	GetActiveSlot() *int32
}

type HotbarItemElement struct {
	Container
	slotId int32
	state  HotbarState
}

func (hbie *HotbarItemElement) Draw(ctx DrawContext) {
	//TODO: Move to hover
	// why did I even do this it looks so awful...
	ai := hbie.state.GetActiveSlot()
	if ai != nil {
		if *ai == hbie.slotId {
			rl.DrawRectangleLinesEx(hbie.bounds, hbie.Style.Border.Thickness+2, hbie.Style.Border.Color)
		} else {
			rl.DrawRectangleLinesEx(hbie.bounds, hbie.Style.Border.Thickness, hbie.Style.Border.Color)
		}
	} else {
		rl.DrawRectangleLinesEx(hbie.bounds, hbie.Style.Border.Thickness, hbie.Style.Border.Color)
	}
	name := hbie.state.SlotItem(hbie.slotId)
	rl.DrawText(name, int32(hbie.Bounds().X), int32(hbie.Bounds().Y), 10, rl.DarkGray)
	for _, child := range hbie.Children() {
		child.Draw(ctx)
	}
}

type HotbarElement struct {
	Container
}

func (hbe *HotbarElement) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(hbe.bounds, hbe.Style.Border.Thickness, hbe.Style.Border.Color)
	for _, child := range hbe.Children() {
		child.Draw(ctx)
	}
}

// TODO: Not a fan of state being on the struct.
func NewHotbarItemElement(bounds rl.Rectangle, slotId int32, state HotbarState) HotbarItemElement {
	return HotbarItemElement{
		Container: NewContainer(bounds, WithBorder(1, rl.DarkBlue)),
		slotId:    slotId,
		state:     state,
	}
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) HotbarElement {
	e := HotbarElement{
		Container: NewContainer(bounds,
			WithLayout(LayoutHorizontal),
			WithWidth(bounds.Width-197),
			WithHeight(48),
			WithOffset(197, bounds.Height-48),
			WithGap(2),
			WithPadding(2),
			WithBorder(1, rl.DarkGray)),
	}

	for i := range 10 {
		ce := NewHotbarItemElement(e.Bounds(), int32(i), state)
		e.AddChild(&ce)
	}

	return e
}

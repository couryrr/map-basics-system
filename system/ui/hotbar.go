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
	SlotItem(i int) string
	GetActiveSlot() *int32
}

type HotbarItemElement struct {
	Container
}

func (hbie *HotbarItemElement) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(hbie.Bounds(), 1, rl.DarkGray)

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

func NewHotbarItemElement(bounds rl.Rectangle, state HotbarState) HotbarItemElement {
	return HotbarItemElement{
		Container: NewContainer(bounds, WithBorder(1, rl.DarkBlue)),
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

	for range 6 {
		ce := NewHotbarItemElement(e.Bounds(), state)
		e.AddChild(&ce)
	}

	return e
}

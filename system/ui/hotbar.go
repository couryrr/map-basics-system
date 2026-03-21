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
	Index int
}

type HotbarElement struct {
	Container
}

func NewHotbarItemElement(bounds rl.Rectangle, index int) HotbarItemElement {
	return HotbarItemElement{
		Container: NewContainer(bounds, WithBorder(1, rl.DarkBlue)),
	}
}

func NewHotbarElement(bounds rl.Rectangle) HotbarElement {
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
	for i := range 6 {
		ce := NewHotbarItemElement(e.Bounds(), i)
		e.AddChild(&ce)
	}
	return e
}

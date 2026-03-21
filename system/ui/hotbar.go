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
	*Container
	slotId int32
	state  HotbarState
}

func (hbie *HotbarItemElement) Draw(ctx DrawContext) {
	borderThickness := hbie.Style.Border.Thickness
	rl.TraceLog(rl.LogInfo, "The elem state is %v", hbie.ElementState())
	if hbie.ElementState() == ElementStateHovered {
		borderThickness += 2
	}
	rl.DrawRectangleLinesEx(hbie.bounds, borderThickness, hbie.Style.Border.Color)
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

// TODO: Containers should have a prop (yes like react (I like solidjs more)).
func NewHotbarItemElement(bounds rl.Rectangle, slotId int32, state HotbarState) HotbarItemElement {
	container := NewContainer(bounds, NewStyle().Border(1, rl.DarkBlue).Build())
	hbie := HotbarItemElement{
		Container: &container,
		slotId:    slotId,
		state:     state,
	}

	//TODO: The container should manage the state not the caller. All the caller should do is set Styles based on the state.
	hbie.AddEventListener(MouseHoverEvent, func(event InputEvent) {
		if rl.CheckCollisionPointRec(event.Position, hbie.Bounds()) {
			hbie.SetElementState(ElementStateHovered)
		} else {
			hbie.SetElementState(ElementStateNormal)
		}
	})

	return hbie
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) HotbarElement {
	e := HotbarElement{
		Container: NewContainer(bounds, NewStyle().
			Layout(LayoutHorizontal).
			Width(bounds.Width-197).
			Height(48).
			Offset(197, bounds.Height-48).
			Gap(2).
			Padding(2).
			Border(1, rl.DarkGray).
			Build()),
	}

	for i := range 10 {
		ce := NewHotbarItemElement(e.Bounds(), int32(i), state)
		e.AddChild(&ce)
	}

	return e
}

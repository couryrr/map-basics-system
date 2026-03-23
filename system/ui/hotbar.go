package ui

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/ui/framework"
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

type HotbarItem struct {
	SlotId int32
	state  HotbarState
}

func NewHotbarItemElement(bounds rl.Rectangle, prop HotbarItem) framework.TypedElement[HotbarItem] {
	element := framework.NewTypedElement(bounds,
		framework.NewStyle().
			Border(1, rl.DarkBlue).
			Font(framework.DefaultFont(10, rl.DarkGray, framework.TextAlignCenter)).
			Build(),
		prop.state.SlotItem(prop.SlotId), prop)

	//TODO: The container should manage the state not the caller. All the caller should do is set Styles based on the state.
	element.AddEventListener(framework.MouseHoverEvent, func(event framework.InputEvent) {
		if rl.CheckCollisionPointRec(event.Position, element.Bounds()) {
			element.SetElementState(framework.ElementStateHovered)
		} else {
			element.SetElementState(framework.ElementStateNone)
		}
	})

	return element
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) framework.Element {
	element := framework.NewElement(bounds,
		framework.NewStyle().Layout(framework.LayoutHorizontal).
			Width(bounds.Width-197).
			Height(48).
			Offset(197, bounds.Height-48).
			Gap(2).
			Padding(2).
			Border(1, rl.DarkGray).
			Build(), "")

	for i := range 10 {
		ce := NewHotbarItemElement(element.Bounds(), HotbarItem{
			SlotId: int32(i),
			state:  state,
		})
		element.AddChild(&ce)
	}

	return element
}

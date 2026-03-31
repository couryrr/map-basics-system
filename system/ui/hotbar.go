package ui

import (
	"github.com/couryrr/map-basics-system/framework/queue"
	"github.com/couryrr/map-basics-system/framework/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HotbarAction string

const (
	TopicUiHotbarInteraction queue.Topic = "ui.hotbar.interaction"
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
	GetActiveSlot() int32
}

type HotbarItem struct {
	SlotId int32
	state  HotbarState
}

func NewHotbarItemElement(bounds rl.Rectangle, state *HotbarItem) ui.TypedElement[HotbarItem] {
	element := ui.NewTypedElement(bounds, state)
	element.WithPropFn(func() ui.Prop {
		prop := ui.Prop{}

		prop.Text = state.state.SlotItem(state.SlotId)
		prop.Style = ui.NewStyle().
			Border(1, rl.DarkBlue).
			Font(ui.DefaultFont(10, rl.DarkGray, ui.TextAlignCenter)).
			Build()

		if state.state.GetActiveSlot() == element.Type.SlotId {
			prop.Style = ui.NewStyle().
				Border(1, rl.Red).
				Font(ui.DefaultFont(10, rl.DarkGray, ui.TextAlignCenter)).
				Build()
		}
		return prop
	})

	element.OnClick(func(e *ui.UiEvent) {
		rl.TraceLog(rl.LogInfo, "the value is: %v", element.Type.SlotId)
	})

	return element
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) ui.Element {
	element := ui.NewElement()
	element.WithPropFn(func() ui.Prop {
		return ui.Prop{
			Style: ui.NewStyle().Layout(ui.LayoutHorizontal).
				Width(bounds.Width-197).
				Height(48).
				Offset(197, bounds.Height-48).
				Gap(2).
				Padding(2).
				Border(1, rl.DarkGray).
				Build(),
		}
	})

	for i := range 10 {
		ce := NewHotbarItemElement(element.Bounds(), &HotbarItem{
			SlotId: int32(i),
			state:  state,
		})
		element.AddChild(&ce)
	}

	return element
}

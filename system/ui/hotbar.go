package ui

import (
	"github.com/couryrr/map-basics-system/framework/pubsub"
	"github.com/couryrr/map-basics-system/framework/ui"
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
	GetActiveSlot() int32
}

type HotbarItem struct {
	SlotId int32
	state  HotbarState
}

func NewHotbarItemElement(bounds rl.Rectangle, state *HotbarItem) framework.TypedElement[HotbarItem] {
	element := framework.NewTypedElement(bounds, state)

	element.WithPropFn(func() framework.Prop {
		prop := framework.Prop{}

		prop.Text = state.state.SlotItem(state.SlotId)
		prop.Style = framework.NewStyle().
			Border(1, rl.DarkBlue).
			Font(framework.DefaultFont(10, rl.DarkGray, framework.TextAlignCenter)).
			Build()

		if state.state.GetActiveSlot() == element.Props.SlotId {
			prop.Style = framework.NewStyle().
				Border(1, rl.Red).
				Font(framework.DefaultFont(10, rl.DarkGray, framework.TextAlignCenter)).
				Build()
		}
		return prop
	})

	return element
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) framework.Element {
	element := framework.NewElement()
	element.WithPropFn(func() framework.Prop {
		return framework.Prop{
			Style: framework.NewStyle().Layout(framework.LayoutHorizontal).
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

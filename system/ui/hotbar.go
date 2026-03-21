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

type HotbarItemElement struct {
	*framework.TypedContainer[HotbarItem]
}

func (hbie *HotbarItemElement) Draw() {
	borderThickness := hbie.Style.Border.Thickness
	rl.TraceLog(rl.LogInfo, "The elem state is %v", hbie.ElementState())
	if hbie.ElementState() == framework.ElementStateHovered {
		borderThickness += 2
	}
	rl.DrawRectangleLinesEx(hbie.Bounds(), borderThickness, hbie.Style.Border.Color)
	name := hbie.Props.state.SlotItem(hbie.Props.SlotId)
	if fs := hbie.Style.Font; fs != nil {
		pos := fs.Position(name, hbie.Bounds())
		rl.DrawTextEx(fs.Font, name, pos, fs.Size, fs.Spacing, fs.Color)
	}
	for _, child := range hbie.Children() {
		child.Draw()
	}
}

type HotbarElement struct {
	framework.Container
}

func (hbe *HotbarElement) Draw() {
	rl.DrawRectangleLinesEx(hbe.Bounds(), hbe.Style.Border.Thickness, hbe.Style.Border.Color)
	for _, child := range hbe.Children() {
		child.Draw()
	}
}

func NewHotbarItemElement(bounds rl.Rectangle, prop HotbarItem) HotbarItemElement {
	container := framework.NewTypedContainer(bounds, framework.NewStyle().Border(1, rl.DarkBlue).Font(framework.DefaultFont(10, rl.DarkGray, framework.TextAlignCenter)).Build(), prop)
	hbie := HotbarItemElement{
		TypedContainer: &container,
	}

	//TODO: The container should manage the state not the caller. All the caller should do is set Styles based on the state.
	hbie.AddEventListener(framework.MouseHoverEvent, func(event framework.InputEvent) {
		if rl.CheckCollisionPointRec(event.Position, hbie.Bounds()) {
			hbie.SetElementState(framework.ElementStateHovered)
		} else {
			hbie.SetElementState(framework.ElementStateNormal)
		}
	})

	return hbie
}

func NewHotbarElement(bounds rl.Rectangle, state HotbarState) HotbarElement {
	e := HotbarElement{
		Container: framework.NewContainer(bounds, framework.NewStyle().
			Layout(framework.LayoutHorizontal).
			Width(bounds.Width-197).
			Height(48).
			Offset(197, bounds.Height-48).
			Gap(2).
			Padding(2).
			Border(1, rl.DarkGray).
			Build()),
	}

	for i := range 10 {
		ce := NewHotbarItemElement(e.Bounds(), HotbarItem{
			SlotId: int32(i),
			state: state,
		})
		e.AddChild(&ce)
	}

	return e
}

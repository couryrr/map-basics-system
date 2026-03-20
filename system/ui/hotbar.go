package ui

import (
	"fmt"

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
	Bound rl.Rectangle
	Index int
}

func (e *HotbarItemElement) Bounds() rl.Rectangle     { return e.Bound }
func (e *HotbarItemElement) SetBounds(b rl.Rectangle) { e.Bound = b }
func (e *HotbarItemElement) Children() []Element      { return nil }
func (e *HotbarItemElement) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(e.Bound, 1, rl.DarkBlue)

	state := ctx.GetHotbarState()
	itemId := state.SlotItem(e.Index)
	if itemId != "" {
		item, err := ctx.GetItemByIdFromRegistry(itemId)
		if err != nil {
			rl.TraceLog(rl.LogInfo, err.Error())
		} else {
			rl.DrawText(fmt.Sprintf("%s", item.Name), int32(e.Bound.X+2), int32(e.Bound.Y+2), 12, item.Color)
		}
	}

	if active := state.GetActiveSlot(); active != nil && *active == int32(e.Index) {
		rl.DrawRectangleLinesEx(e.Bound, 5, rl.Red)
	}
}

type HotbarElement struct {
	container Container
}

func (e *HotbarElement) Bounds() rl.Rectangle     { return e.container.Bounds() }
func (e *HotbarElement) SetBounds(b rl.Rectangle) { e.container.SetBounds(b) }
func (e *HotbarElement) Children() []Element      { return e.container.Children() }
func (e *HotbarElement) AddChild(ce Element)      { e.container.AddChild(ce) }
func (e *HotbarElement) Draw(ctx DrawContext)     { e.container.Draw(ctx) }

func (e *HotbarElement) HandleIntersection(point rl.Vector2) *InteractionResult {
	return nil
}

func NewHotbarElement(bounds rl.Rectangle) *HotbarElement {
	e := &HotbarElement{
		container: NewContainer(bounds, WithLayout(LayoutHorizontal),
			WithWidth(bounds.Width-197), WithHeight(48), WithOffset(197, bounds.Height-48),
			WithGap(2), WithPadding(2),
			WithBorder(1, rl.DarkGray)),
	}
	for i := range 6 {
		e.container.AddChild(&HotbarItemElement{Index: i})
	}
	return e
}

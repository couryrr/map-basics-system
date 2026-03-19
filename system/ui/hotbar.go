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
	Bound rl.Rectangle
}

type HotbarElement struct {
	container Container
}

func (e HotbarElement) Bounds() rl.Rectangle         { return e.container.bounds }
func (e HotbarElement) SetBounds(bound rl.Rectangle) { e.container.bounds = bound }
func (e HotbarElement) Children() []Element          { return e.container.Children() }
func (e HotbarElement) AddChild(ce Element)          { e.container.AddChild(ce) }
func (e HotbarElement) Draw(ctx DrawContext)         { e.container.Draw(ctx) }

func (e HotbarElement) HandleIntersection(point rl.Vector2) *InteractionResult {
	return nil
}

func NewHotbarElement(bounds rl.Rectangle) *HotbarElement {
	return &HotbarElement{
		container: NewContainer(bounds, LayoutNone, WithHeight(48), WithOffset(0, bounds.Height-48), WithBorder(1, rl.DarkGray)),
	}
}

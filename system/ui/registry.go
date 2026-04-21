package ui

import (
	"iter"

	"github.com/couryrr/map-basics-system/framework/ui"
	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryState interface {
	GetItems() iter.Seq2[string, world.GameItem]
}

// TODO: Do not pass GameItem directly could be an interface
func NewRegistryItemElement(bounds rl.Rectangle, gameItem *world.GameItem) ui.TypedElement[world.GameItem] {
	element := ui.NewTypedElement(bounds, gameItem)
	element.WithPropFn(func(ctx *ui.UiContext) ui.Prop {
		borderColor := rl.DarkGray
		if element.Id == ctx.HoveredId {
			borderColor = rl.Red
		}
		return ui.Prop{
			Style: ui.NewStyle().
				Layout(ui.LayoutGrid).
				Width(200).
				Border(1, borderColor).
				Gap(2).
				Padding(4).
				CellHeight(48).
				Columns(2).
				Font(ui.DefaultFont(10, rl.DarkGray, ui.TextAlignCenter)).
				Build(),
			Text: gameItem.Name,
		}
	})
	return element
}

func NewRegistryElement(bounds rl.Rectangle, state RegistryState) ui.Element {
	element := ui.NewElement()
	element.WithPropFn(func(ctx *ui.UiContext) ui.Prop {
		return ui.Prop{
			Style: ui.NewStyle().
				Layout(ui.LayoutGrid).
				Width(200).
				Border(1, rl.DarkGray).
				Gap(2).
				Padding(4).
				CellHeight(48).
				Columns(2).
				Build(),
		}
	})

	for _, item := range state.GetItems() {
		c := NewRegistryItemElement(bounds, &item)
		element.AddChild(&c)
	}

	return element
}

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
func NewRegistryItemElement(bounds rl.Rectangle, gameItem *world.GameItem) framework.TypedElement[world.GameItem] {
	element := framework.NewTypedElement(bounds, gameItem)
	element.WithStyleFn(func() framework.Style {
		return framework.NewStyle().
			Layout(framework.LayoutGrid).
			Width(200).
			Border(1, rl.DarkGray).
			Gap(2).
			Padding(4).
			CellHeight(48).
			Columns(2).
			Font(framework.DefaultFont(10, rl.DarkGray, framework.TextAlignCenter)).
			Build()
	})
	element.WithTextFn(func() string { return gameItem.Name })
	return element
}

func NewRegistryElement(bounds rl.Rectangle, state RegistryState) framework.Element {
	element := framework.NewElement()
	element.WithStyleFn(func() framework.Style {
		return framework.NewStyle().
			Layout(framework.LayoutGrid).
			Width(200).
			Border(1, rl.DarkGray).
			Gap(2).
			Padding(4).
			CellHeight(48).
			Columns(2).
			Build()
	})

	for _, item := range state.GetItems() {
		c := NewRegistryItemElement(bounds, &item)
		element.AddChild(&c)
	}

	return element
}

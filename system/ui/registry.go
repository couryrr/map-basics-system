package ui

import (
	"iter"

	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryState interface {
	GetItems() iter.Seq2[string, world.GameItem]
}

type RegistryItemElement struct {
	Container
	gameItem *world.GameItem
}

func (rie *RegistryItemElement) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(rie.bounds, rie.Style.Border.Thickness, rie.Style.Border.Color)
	rl.DrawText(rie.gameItem.Name, int32(rie.Bounds().X), int32(rie.Bounds().Y), 10, rie.gameItem.Color)

	//TODO: the system should handle the children draw some how.
	for _, child := range rie.Children() {
		child.Draw(ctx)
	}
}

type RegistryElement struct {
	Container
}

func (re *RegistryElement) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(re.bounds, re.Style.Border.Thickness, re.Style.Border.Color)
	for _, child := range re.Children() {
		child.Draw(ctx)
	}
}

// TODO: Do not pass GameItem directly
func NewRegistryItemElement(bounds rl.Rectangle, gameItem world.GameItem) RegistryItemElement {
	e := RegistryItemElement{
		gameItem: &gameItem,
		Container: NewContainer(bounds,
			WithLayout(LayoutGrid),
			WithWidth(200),
			WithBorder(1, rl.DarkGray),
			WithGap(2),
			WithPadding(4),
			WithCellHeight(48),
			WithColumns(2)),
	}

	return e
}

func NewRegistryElement(bounds rl.Rectangle, state RegistryState) RegistryElement {
	e := RegistryElement{Container: NewContainer(bounds,
		WithLayout(LayoutGrid),
		WithWidth(200),
		WithBorder(1, rl.DarkGray),
		WithGap(2),
		WithPadding(4),
		WithCellHeight(48),
		WithColumns(2)),
	}

	for _, item := range state.GetItems() {
		c := NewRegistryItemElement(bounds, item)
		e.AddChild(&c)
	}

	return e
}

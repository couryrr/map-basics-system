package ui

import (
	"image/color"
	"iter"

	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryItemElement struct {
	Bound rl.Rectangle
	Name  string
	Color color.RGBA
}

func (e *RegistryItemElement) Bounds() rl.Rectangle        { return e.Bound }
func (e *RegistryItemElement) SetBounds(b rl.Rectangle)   { e.Bound = b }
func (e *RegistryItemElement) Children() []Element         { return nil }
func (e *RegistryItemElement) Draw(_ DrawContext) {
	rl.DrawRectangleLinesEx(e.Bound, 1, rl.DarkGray)
	rl.DrawText(e.Name, int32(e.Bound.X+2), int32(e.Bound.Y+2), 10, e.Color)
}

type RegistryElement struct {
	container Container
}

func (e *RegistryElement) Bounds() rl.Rectangle        { return e.container.Bounds() }
func (e *RegistryElement) SetBounds(b rl.Rectangle)   { e.container.SetBounds(b) }
func (e *RegistryElement) Children() []Element         { return e.container.Children() }
func (e *RegistryElement) AddChild(ce Element)         { e.container.AddChild(ce) }
func (e *RegistryElement) Draw(ctx DrawContext)         { e.container.Draw(ctx) }

func NewRegistryElement(bounds rl.Rectangle, items iter.Seq2[string, world.GameItem]) *RegistryElement {
	e := &RegistryElement{
		container: NewContainer(bounds, WithLayout(LayoutGrid),
			WithWidth(200), WithBorder(1, rl.DarkGray),
			WithGap(2), WithPadding(4), WithCellHeight(48), WithColumns(2)),
	}
	for _, item := range items {
		e.container.AddChild(&RegistryItemElement{Name: item.Name, Color: item.Color})
	}
	return e
}

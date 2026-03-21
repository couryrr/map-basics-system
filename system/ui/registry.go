package ui

import (
	"image/color"
	"iter"

	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryItemElement struct {
	Container
	Name  string
	Color color.RGBA
}

type RegistryElement struct {
	Container
}

func NewRegistryElement(bounds rl.Rectangle, items iter.Seq2[string, world.GameItem]) RegistryElement {
	e := RegistryElement{Container: NewContainer(bounds,
		WithLayout(LayoutGrid),
		WithWidth(200),
		WithBorder(1, rl.DarkGray),
		WithGap(2),
		WithPadding(4),
		WithCellHeight(48),
		WithColumns(2)),
	}

	for _, item := range items {
		e.AddChild(&RegistryItemElement{
			Container: NewContainer(e.bounds, WithBorder(1, rl.DarkGray)),
			Name:      item.Name,
			Color:     item.Color,
		})
	}

	return e
}

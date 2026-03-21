package ui

import (
	"iter"

	"github.com/couryrr/map-basics-system/system/ui/framework"
	"github.com/couryrr/map-basics-system/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryState interface {
	GetItems() iter.Seq2[string, world.GameItem]
}

type RegistryItemElement struct {
	framework.Container
	gameItem *world.GameItem
}

func (rie *RegistryItemElement) Draw() {
	rl.DrawRectangleLinesEx(rie.Bounds(), rie.Style.Border.Thickness, rie.Style.Border.Color)
	if fs := rie.Style.Font; fs != nil {
		textSize := rl.MeasureTextEx(fs.Font, rie.gameItem.Name, fs.Size, fs.Spacing)
		x := rie.Bounds().X + (rie.Bounds().Width-textSize.X)/2
		y := rie.Bounds().Y + (rie.Bounds().Height-textSize.Y)/2
		rl.DrawTextEx(fs.Font, rie.gameItem.Name, rl.NewVector2(x, y), fs.Size, fs.Spacing, rie.gameItem.Color)
	}

	//TODO: the system should handle the children draw.
	//TODO: this draw could return a method that is used by container.
	for _, child := range rie.Children() {
		child.Draw()
	}
}

type RegistryElement struct {
	framework.Container
}

func (re *RegistryElement) Draw() {
	rl.DrawRectangleLinesEx(re.Bounds(), re.Style.Border.Thickness, re.Style.Border.Color)
	for _, child := range re.Children() {
		child.Draw()
	}
}

// TODO: Do not pass GameItem directly could be an interface
func NewRegistryItemElement(bounds rl.Rectangle, gameItem world.GameItem) RegistryItemElement {
	e := RegistryItemElement{
		gameItem: &gameItem,
		Container: framework.NewContainer(bounds, framework.NewStyle().
			Layout(framework.LayoutGrid).
			Width(200).
			Border(1, rl.DarkGray).
			Gap(2).
			Padding(4).
			CellHeight(48).
			Columns(2).
			Font(framework.DefaultFont(10, rl.DarkGray)).
			Build()),
	}

	return e
}

func NewRegistryElement(bounds rl.Rectangle, state RegistryState) RegistryElement {
	e := RegistryElement{Container: framework.NewContainer(bounds, framework.NewStyle().
		Layout(framework.LayoutGrid).
		Width(200).
		Border(1, rl.DarkGray).
		Gap(2).
		Padding(4).
		CellHeight(48).
		Columns(2).
		Build()),
	}

	for _, item := range state.GetItems() {
		c := NewRegistryItemElement(bounds, item)
		e.AddChild(&c)
	}

	return e
}

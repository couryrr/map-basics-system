package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegistryElement struct {
	container Container
	Bound rl.Rectangle
}

func (e RegistryElement) Bounds() rl.Rectangle         { return e.container.bounds }
func (e RegistryElement) SetBounds(bound rl.Rectangle) { e.container.bounds = bound }
func (e RegistryElement) Children() []Element          { return e.container.Children() }
func (e RegistryElement) AddChild(ce Element)           { e.container.AddChild(ce) }
func (e RegistryElement) Draw(ctx DrawContext)         { e.container.Draw(ctx) }

func NewRegistryElement(bounds rl.Rectangle) RegistryElement {
    return RegistryElement{
		container: NewContainer(bounds, LayoutNone, WithWidth(200), WithBorder(1, rl.DarkGray)),
	}
}

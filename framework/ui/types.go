package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout int
type TextAlign int
type InputEventType int
type ElementState int

const (
	LayoutNone Layout = iota
	LayoutHorizontal
	LayoutVertical
	LayoutGrid
)

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)

const (
	MouseClickEvent InputEventType = iota
	MouseHoverEvent
	MouseDragEvent
)

type UiContext struct{
	Hovered Drawable
}

type Drawable interface {
	Bounds() rl.Rectangle
	ComputeBounds(rl.Rectangle)
	SetBounds(rl.Rectangle)
	Children() []Drawable
	Parent() Drawable
	SetParent(parent Drawable)
	draw(ctx *UiContext)
	hitTest(point *rl.Vector2) Drawable
	bubble(event UiEvent)
}

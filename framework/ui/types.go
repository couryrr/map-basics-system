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

const (
	ElementStateNone ElementState = iota
	ElementStateHovered
	ElementStateActive
	ElementStateFocused
)

type InputEvent struct {
	Position  rl.Vector2
	EventType InputEventType
}

type Drawable interface {
	Draw()
	Bounds() rl.Rectangle
	ComputeBounds(rl.Rectangle)
	SetBounds(rl.Rectangle)
	Children() []Drawable
	Parent() Drawable
	SetParent(parent Drawable)
	hitTest(point rl.Vector2) Drawable
	bubble(uiEvent *UiEvent)
}

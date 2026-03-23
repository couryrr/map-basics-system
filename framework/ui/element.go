package framework

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Element struct {
	bounds       rl.Rectangle
	style        Style
	text         string
	children     []Drawable
	inputEvents  map[InputEventType][]func(event InputEvent)
	elementState ElementState
}

func NewElement(bound rl.Rectangle, style Style, text string) Element {
	c := Element{
		style:        style,
		text:         text,
		elementState: ElementStateNone,
		inputEvents:  make(map[InputEventType][]func(event InputEvent)),
	}
	c.ComputeBounds(bound)
	return c
}

func (elm *Element) HandleEvents(event InputEvent) {
	if elm.inputEvents != nil {
		callbacks := elm.inputEvents[event.EventType]
		for _, cb := range callbacks {
			cb(event)
		}
	}

	for _, child := range elm.children {
		child.HandleEvents(event)
	}
}

func (elm *Element) AddEventListener(eventType InputEventType, cb func(event InputEvent)) {
	elm.inputEvents[eventType] = append(elm.inputEvents[eventType], cb)
}

func (elm *Element) ElementState() ElementState {
	return elm.elementState
}

func (elm *Element) SetElementState(es ElementState) {
	elm.elementState = es
}

func (elm *Element) Bounds() rl.Rectangle     { return elm.bounds }
func (elm *Element) SetBounds(b rl.Rectangle) { elm.bounds = b; elm.applyLayout() }
func (elm *Element) ComputeBounds(b rl.Rectangle) {
	inset := elm.style.Margin
	if elm.style.Border != nil {
		inset += elm.style.Border.Thickness
	}
	w := b.Width
	h := b.Height
	if elm.style.Width != 0 {
		w = elm.style.Width
	}
	if elm.style.Height != 0 {
		h = elm.style.Height
	}
	elm.bounds = rl.NewRectangle(
		b.X+inset+elm.style.OffsetX,
		b.Y+inset+elm.style.OffsetY,
		w-inset*2,
		h-inset*2,
	)
	elm.applyLayout()
}

func (elm *Element) Children() []Drawable { return elm.children }
func (elm *Element) AddChild(e Drawable) {
	elm.children = append(elm.children, e)
	elm.applyLayout()
}

func (elm *Element) Draw() {
	rl.DrawRectangleLinesEx(elm.Bounds(), elm.style.Border.Thickness, elm.style.Border.Color)
	if fs := elm.style.Font; fs != nil && elm.text != "" {
		pos := fs.Position(elm.text, elm.bounds)
		rl.DrawTextEx(fs.Font, elm.text, pos, fs.Size, fs.Spacing, fs.Color)
	}
	for _, child := range elm.Children() {
		child.Draw()
	}
}

func (elm *Element) applyLayout() {
	n := len(elm.children)
	if n == 0 || elm.style.Layout == LayoutNone {
		return
	}

	p := elm.style.Padding
	g := elm.style.Gap

	switch elm.style.Layout {
	case LayoutHorizontal:
		slotW := (elm.bounds.Width - p*2 - g*float32(n-1)) / float32(n)
		slotH := elm.bounds.Height - p*2
		x := elm.bounds.X + p
		for _, child := range elm.children {
			child.SetBounds(rl.NewRectangle(x, elm.bounds.Y+p, slotW, slotH))
			x += slotW + g
		}
	case LayoutVertical:
		slotW := elm.bounds.Width - p*2
		slotH := (elm.bounds.Height - p*2 - g*float32(n-1)) / float32(n)
		y := elm.bounds.Y + p
		for _, child := range elm.children {
			child.SetBounds(rl.NewRectangle(elm.bounds.X+p, y, slotW, slotH))
			y += slotH + g
		}
	case LayoutGrid:
		cols := elm.style.Columns
		if cols <= 0 {
			cols = 1
		}
		rows := (n + cols - 1) / cols
		slotW := (elm.bounds.Width - p*2 - g*float32(cols-1)) / float32(cols)
		slotH := elm.style.CellHeight
		if slotH == 0 {
			slotH = (elm.bounds.Height - p*2 - g*float32(rows-1)) / float32(rows)
		}
		for i, child := range elm.children {
			col := i % cols
			row := i / cols
			x := elm.bounds.X + p + float32(col)*(slotW+g)
			y := elm.bounds.Y + p + float32(row)*(slotH+g)
			child.SetBounds(rl.NewRectangle(x, y, slotW, slotH))
		}
	}
}

type TypedElement[T any] struct {
	Element
	Props T
}

func NewTypedElement[T any](bound rl.Rectangle, style Style, text string, prop T) TypedElement[T] {
	container := NewElement(bound, style, text)
	return TypedElement[T]{
		Element: container,
		Props:   prop,
	}
}

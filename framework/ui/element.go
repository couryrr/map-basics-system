package framework

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type StyleFn func() Style
type TextFn func() string

type Element struct {
	bounds       rl.Rectangle
	styleFn      StyleFn //Current thought is only change style not the container
	textFn       TextFn  // Same idea for the text?
	children     []Drawable
	inputEvents  map[InputEventType][]func(event InputEvent)
	elementState ElementState
}

func NewRoot(rootBound rl.Rectangle) Element {
	elm := Element{
		bounds:       rootBound,
		styleFn:      func() Style { return DefaultStyle() },
		textFn:       func() string { return "" },
		elementState: ElementStateNone,
		inputEvents:  make(map[InputEventType][]func(event InputEvent)),
	}
	return elm

}

func NewElement() Element {
	elm := Element{
		styleFn:      func() Style { return DefaultStyle() },
		textFn:       func() string { return "" },
		elementState: ElementStateNone,
		inputEvents:  make(map[InputEventType][]func(event InputEvent)),
	}
	return elm
}

func (elm *Element) WithStyleFn(styleFn StyleFn) {
	elm.styleFn = styleFn
}

func (elm *Element) WithTextFn(textFn TextFn) {
	elm.textFn = textFn
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

// TODO: this math is from an LLM verify
func (elm *Element) ComputeBounds(b rl.Rectangle) {
	style := elm.styleFn()
	inset := style.Margin
	if style.Border != nil {
		inset += style.Border.Thickness
	}
	w := b.Width
	h := b.Height
	if style.Width != 0 {
		w = style.Width
	}
	if style.Height != 0 {
		h = style.Height
	}
	elm.bounds = rl.NewRectangle(
		b.X+inset+style.OffsetX,
		b.Y+inset+style.OffsetY,
		w-inset*2,
		h-inset*2,
	)
	elm.applyLayout()
}

func (elm *Element) Children() []Drawable { return elm.children }
func (elm *Element) AddChild(e Drawable) {
	e.ComputeBounds(elm.Bounds())
	elm.children = append(elm.children, e)
	elm.applyLayout()
}

func (elm *Element) Draw() {
	style := elm.styleFn()
	rl.DrawRectangleLinesEx(elm.Bounds(), style.Border.Thickness, style.Border.Color)

	text := elm.textFn()
	if fs := style.Font; fs != nil && text != "" {
		pos := fs.Position(text, elm.bounds)
		rl.DrawTextEx(fs.Font, text, pos, fs.Size, fs.Spacing, fs.Color)
	}
	for _, child := range elm.Children() {
		child.Draw()
	}
}

// TODO: this math is from an LLM verify
func (elm *Element) applyLayout() {
	style := elm.styleFn()
	n := len(elm.children)
	if n == 0 || style.Layout == LayoutNone {
		return
	}

	p := style.Padding
	g := style.Gap

	switch style.Layout {
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
		cols := style.Columns
		if cols <= 0 {
			cols = 1
		}
		rows := (n + cols - 1) / cols
		slotW := (elm.bounds.Width - p*2 - g*float32(cols-1)) / float32(cols)
		slotH := style.CellHeight
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
	Props *T
}

func NewTypedElement[T any](bound rl.Rectangle, prop *T) TypedElement[T] {
	telm := TypedElement[T]{
		Element: NewElement(),
		Props:   prop,
	}
	return telm
}

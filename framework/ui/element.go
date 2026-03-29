package framework

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Prop struct {
	Style Style
	Text  string
}

type PropFn func() Prop

type Root struct {
	Element
	events []InputEvent
}

func NewRoot(rootBound rl.Rectangle) Root {
	root := Root{
		Element: NewElement(),
		events:  make([]InputEvent, 100),
	}
	root.bounds = rootBound
	root.WithPropFn(func() Prop {
		return Prop{
			Style: DefaultStyle(),
		}
	})
	return root
}

type Element struct {
	bounds       rl.Rectangle
	propFn       PropFn
	children     []Drawable
	inputEvents  map[InputEventType][]func(event InputEvent)
	elementState ElementState
}

func NewElement() Element {
	elm := Element{
		elementState: ElementStateNone,
		inputEvents:  make(map[InputEventType][]func(event InputEvent)),
	}
	return elm
}

func (elm *Element) WithPropFn(propFn PropFn) {
	elm.propFn = propFn
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
	if elm.propFn != nil {
		props := elm.propFn()
		inset := props.Style.Margin
		if props.Style.Border != nil {
			inset += props.Style.Border.Thickness
		}
		w := b.Width
		h := b.Height
		if props.Style.Width != 0 {
			w = props.Style.Width
		}
		if props.Style.Height != 0 {
			h = props.Style.Height
		}
		elm.bounds = rl.NewRectangle(
			b.X+inset+props.Style.OffsetX,
			b.Y+inset+props.Style.OffsetY,
			w-inset*2,
			h-inset*2,
		)
		elm.applyLayout()
	}
}

func (elm *Element) Children() []Drawable { return elm.children }
func (elm *Element) AddChild(e Drawable) {
	e.ComputeBounds(elm.Bounds())
	elm.children = append(elm.children, e)
	elm.applyLayout()
}

func (elm *Element) Draw() {
	if elm.propFn != nil {
		props := elm.propFn()
		rl.DrawRectangleLinesEx(elm.Bounds(), props.Style.Border.Thickness, props.Style.Border.Color)

		if fs := props.Style.Font; fs != nil && props.Text != "" {
			pos := fs.Position(props.Text, elm.bounds)
			rl.DrawTextEx(fs.Font, props.Text, pos, fs.Size, fs.Spacing, fs.Color)
		}

		for _, child := range elm.Children() {
			child.Draw()
		}
	}
}

// TODO: this math is from an LLM verify
func (elm *Element) applyLayout() {
	if elm.propFn != nil {
		props := elm.propFn()
		n := len(elm.children)
		if n == 0 || props.Style.Layout == LayoutNone {
			return
		}

		p := props.Style.Padding
		g := props.Style.Gap

		switch props.Style.Layout {
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
			cols := props.Style.Columns
			if cols <= 0 {
				cols = 1
			}
			rows := (n + cols - 1) / cols
			slotW := (elm.bounds.Width - p*2 - g*float32(cols-1)) / float32(cols)
			slotH := props.Style.CellHeight
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

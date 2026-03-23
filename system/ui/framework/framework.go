package framework

import (
	"image/color"

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

	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight

	MouseClickEvent InputEventType = iota
	MouseHoverEvent
	MouseDragEvent

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
	SetBounds(rl.Rectangle)
	Children() []Drawable
	AddEventListener(eventType InputEventType, cb func(event InputEvent))
	HandleEvents(event InputEvent)
	ElementState() ElementState
}

type Style struct {
	Padding      float32
	Margin       float32
	Gap          float32
	Width        float32
	Height       float32
	CellHeight   float32
	Columns      int
	Layout       Layout
	OffsetX      float32
	OffsetY      float32
	BGColor      *color.RGBA
	Border       *Border
	BorderImage  *BorderImage
	Font         *FontStyle
	StyleVariant map[ElementState]Style
}

type Border struct {
	Thickness float32
	Color     color.RGBA
}

type BorderImage struct {
	Texture rl.Texture2D
	Tint    color.RGBA
}

type FontStyle struct {
	Font    rl.Font
	Size    float32
	Spacing float32
	Color   color.RGBA
	Align   TextAlign
}

type StyleBuilder struct {
	s Style
}

func DefaultFont(size float32, c color.RGBA, align TextAlign) FontStyle {
	return FontStyle{Font: rl.GetFontDefault(), Size: size, Spacing: 1, Color: c, Align: align}
}

func (fs FontStyle) Position(text string, bounds rl.Rectangle) rl.Vector2 {
	textSize := rl.MeasureTextEx(fs.Font, text, fs.Size, fs.Spacing)
	y := bounds.Y + (bounds.Height-textSize.Y)/2
	var x float32
	switch fs.Align {
	case TextAlignCenter:
		x = bounds.X + (bounds.Width-textSize.X)/2
	case TextAlignRight:
		x = bounds.X + bounds.Width - textSize.X
	default:
		x = bounds.X
	}
	return rl.NewVector2(x, y)
}

func NewStyle() StyleBuilder { return StyleBuilder{} }

func (b StyleBuilder) Build() Style { return b.s }

func (b StyleBuilder) Padding(p float32) StyleBuilder    { b.s.Padding = p; return b }
func (b StyleBuilder) Margin(m float32) StyleBuilder     { b.s.Margin = m; return b }
func (b StyleBuilder) Gap(g float32) StyleBuilder        { b.s.Gap = g; return b }
func (b StyleBuilder) Width(w float32) StyleBuilder      { b.s.Width = w; return b }
func (b StyleBuilder) Height(h float32) StyleBuilder     { b.s.Height = h; return b }
func (b StyleBuilder) CellHeight(h float32) StyleBuilder { b.s.CellHeight = h; return b }
func (b StyleBuilder) Columns(n int) StyleBuilder        { b.s.Columns = n; return b }
func (b StyleBuilder) Layout(l Layout) StyleBuilder      { b.s.Layout = l; return b }
func (b StyleBuilder) Offset(x, y float32) StyleBuilder  { b.s.OffsetX = x; b.s.OffsetY = y; return b }
func (b StyleBuilder) BGColor(c color.RGBA) StyleBuilder { b.s.BGColor = &c; return b }
func (b StyleBuilder) Border(thickness float32, c color.RGBA) StyleBuilder {
	b.s.Border = &Border{Thickness: thickness, Color: c}
	return b
}
func (b StyleBuilder) Font(f FontStyle) StyleBuilder { b.s.Font = &f; return b }

type Element struct {
	bounds       rl.Rectangle
	style        Style
	text         string
	children     []Drawable
	inputEvents  map[InputEventType][]func(event InputEvent)
	elementState ElementState
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
	// Apply any state variant style
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

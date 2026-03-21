package framework

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout int
type InputEventType int
type ElementState int
type TextAlign int

const (
	TextAlignLeft   TextAlign = iota
	TextAlignCenter
	TextAlignRight
)

const (
	LayoutNone Layout = iota
	LayoutHorizontal
	LayoutVertical
	LayoutGrid

	MouseClickEvent InputEventType = iota
	MouseHoverEvent
	MouseDragEvent

	ElementStateNormal ElementState = iota
	ElementStateHovered
	ElementStateSelected
)

type InputEvent struct {
	Position  rl.Vector2
	EventType InputEventType
}

type Element interface {
	Draw()
	Bounds() rl.Rectangle
	SetBounds(rl.Rectangle)
	Children() []Element
	AddEventListener(eventType InputEventType, cb func(event InputEvent))
	HandleEvents(event InputEvent)
	ElementState() ElementState
}

type Style struct {
	Padding     float32
	Margin      float32
	Gap         float32
	Width       float32
	Height      float32
	CellHeight  float32
	Columns     int
	Layout      Layout
	OffsetX     float32
	OffsetY     float32
	BGColor     *color.RGBA
	Border      *Border
	BorderImage *BorderImage
	Font        *FontStyle
}

type StyleBuilder struct {
	s Style
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

func DefaultFont(size float32, c color.RGBA, align TextAlign) FontStyle {
	return FontStyle{Font: rl.GetFontDefault(), Size: size, Spacing: 1, Color: c, Align: align}
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

type Container struct {
	bounds       rl.Rectangle
	Style        Style
	Layout       Layout
	Columns      int
	children     []Element
	inputEvents  map[InputEventType][]func(event InputEvent)
	elementState ElementState
}

func (c *Container) HandleEvents(event InputEvent) {
	if c.inputEvents != nil {
		callbacks := c.inputEvents[event.EventType]
		for _, cb := range callbacks {
			cb(event)
		}
	}

	for _, child := range c.children {
		child.HandleEvents(event)
	}
}

func (c *Container) AddEventListener(eventType InputEventType, cb func(event InputEvent)) {
	c.inputEvents[eventType] = append(c.inputEvents[eventType], cb)
}

func (c *Container) ElementState() ElementState {
	return c.elementState
}

func (c *Container) SetElementState(es ElementState) {
	c.elementState = es
}

func (c *Container) Bounds() rl.Rectangle     { return c.bounds }
func (c *Container) SetBounds(b rl.Rectangle) { c.bounds = b; c.applyLayout() }

func (c *Container) ComputeBounds(b rl.Rectangle) {
	inset := c.Style.Margin
	if c.Style.Border != nil {
		inset += c.Style.Border.Thickness
	}
	w := b.Width
	h := b.Height
	if c.Style.Width != 0 {
		w = c.Style.Width
	}
	if c.Style.Height != 0 {
		h = c.Style.Height
	}
	c.bounds = rl.NewRectangle(
		b.X+inset+c.Style.OffsetX,
		b.Y+inset+c.Style.OffsetY,
		w-inset*2,
		h-inset*2,
	)
	c.applyLayout()
}
func (c *Container) Children() []Element      { return c.children }
func (c *Container) AddChild(e Element) {
	c.children = append(c.children, e)
	c.applyLayout()
}

func (c *Container) applyLayout() {
	n := len(c.children)
	if n == 0 || c.Layout == LayoutNone {
		return
	}

	p := c.Style.Padding
	g := c.Style.Gap

	switch c.Layout {
	case LayoutHorizontal:
		slotW := (c.bounds.Width - p*2 - g*float32(n-1)) / float32(n)
		slotH := c.bounds.Height - p*2
		x := c.bounds.X + p
		for _, child := range c.children {
			child.SetBounds(rl.NewRectangle(x, c.bounds.Y+p, slotW, slotH))
			x += slotW + g
		}
	case LayoutVertical:
		slotW := c.bounds.Width - p*2
		slotH := (c.bounds.Height - p*2 - g*float32(n-1)) / float32(n)
		y := c.bounds.Y + p
		for _, child := range c.children {
			child.SetBounds(rl.NewRectangle(c.bounds.X+p, y, slotW, slotH))
			y += slotH + g
		}
	case LayoutGrid:
		cols := c.Columns
		if cols <= 0 {
			cols = 1
		}
		rows := (n + cols - 1) / cols
		slotW := (c.bounds.Width - p*2 - g*float32(cols-1)) / float32(cols)
		slotH := c.Style.CellHeight
		if slotH == 0 {
			slotH = (c.bounds.Height - p*2 - g*float32(rows-1)) / float32(rows)
		}
		for i, child := range c.children {
			col := i % cols
			row := i / cols
			x := c.bounds.X + p + float32(col)*(slotW+g)
			y := c.bounds.Y + p + float32(row)*(slotH+g)
			child.SetBounds(rl.NewRectangle(x, y, slotW, slotH))
		}
	}
}

type TypedContainer[T any] struct {
	Container
	Props T
}

func NewTypedContainer[T any](bound rl.Rectangle, style Style, prop T) TypedContainer[T] {
	container := NewContainer(bound, style)
	return TypedContainer[T]{
		Container: container,
		Props: prop,
	}
}

func NewContainer(bound rl.Rectangle, style Style) Container {
	c := Container{
		Layout:       style.Layout,
		Style:        style,
		Columns:      style.Columns,
		elementState: ElementStateNormal,
		inputEvents:  make(map[InputEventType][]func(event InputEvent)),
	}
	c.ComputeBounds(bound)
	return c
}

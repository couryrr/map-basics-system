package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout int
type MouseEventType int

const (
	LayoutNone Layout = iota
	LayoutHorizontal
	LayoutVertical
	LayoutGrid

	MouseClickEvent MouseEventType = iota
	MouseHoveEvent
	MouseDragEvent
)

type MouseEvent struct {
	Position rl.Vector2
	Event    MouseEventType
}

type Element interface {
	Draw(ctx DrawContext)
	Bounds() rl.Rectangle
	SetBounds(rl.Rectangle)
	Children() []Element
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

func DefaultStyle() Style {
	return Style{
		Padding: 0,
		Margin:  0,
		Gap:     0,
		Width:   0,
		Height:  0,
	}
}

func WithPadding(p float32) func(*Style) {
	return func(s *Style) {
		s.Padding = p
	}
}

func WithMargin(m float32) func(*Style) {
	return func(s *Style) {
		s.Margin = m
	}
}

func WithGap(g float32) func(*Style) {
	return func(s *Style) {
		s.Gap = g
	}
}

func WithWidth(w float32) func(*Style) {
	return func(s *Style) {
		s.Width = w
	}
}

func WithHeight(h float32) func(*Style) {
	return func(s *Style) {
		s.Height = h
	}
}

func WithOffset(x, y float32) func(*Style) {
	return func(s *Style) {
		s.OffsetX = x
		s.OffsetY = y
	}
}

func WithBGColor(c color.RGBA) func(*Style) {
	return func(s *Style) {
		s.BGColor = &c
	}
}

func WithLayout(l Layout) func(*Style) {
	return func(s *Style) {
		s.Layout = l
	}
}

func WithColumns(n int) func(*Style) {
	return func(s *Style) {
		s.Columns = n
	}
}

func WithCellHeight(h float32) func(*Style) {
	return func(s *Style) {
		s.CellHeight = h
	}
}

func WithBorder(thickness float32, c color.RGBA) func(*Style) {
	return func(s *Style) {
		s.Border = &Border{Thickness: thickness, Color: c}
	}
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
}

type Container struct {
	bounds   rl.Rectangle
	Style    Style
	Layout   Layout
	Columns  int
	children []Element
	mouseEvents map[MouseEventType][]func(event MouseEvent)
}

func (igo *Container) AddMouseEventHandler(eventType MouseEventType, cb func(event MouseEvent)) {   
	igo.mouseEvents[eventType] = append(igo.mouseEvents[eventType], cb)
}
func (igo *Container) Bounds() rl.Rectangle     { return igo.bounds }
func (igo *Container) SetBounds(b rl.Rectangle) { igo.bounds = b; igo.applyLayout() }
func (igo *Container) Children() []Element      { return igo.children }
func (igo *Container) AddChild(e Element) {
	igo.children = append(igo.children, e)
	igo.applyLayout()
}

func (igo *Container) applyLayout() {
	n := len(igo.children)
	if n == 0 || igo.Layout == LayoutNone {
		return
	}

	p := igo.Style.Padding
	g := igo.Style.Gap

	switch igo.Layout {
	case LayoutHorizontal:
		slotW := (igo.bounds.Width - p*2 - g*float32(n-1)) / float32(n)
		slotH := igo.bounds.Height - p*2
		x := igo.bounds.X + p
		for _, child := range igo.children {
			child.SetBounds(rl.NewRectangle(x, igo.bounds.Y+p, slotW, slotH))
			x += slotW + g
		}
	case LayoutVertical:
		slotW := igo.bounds.Width - p*2
		slotH := (igo.bounds.Height - p*2 - g*float32(n-1)) / float32(n)
		y := igo.bounds.Y + p
		for _, child := range igo.children {
			child.SetBounds(rl.NewRectangle(igo.bounds.X+p, y, slotW, slotH))
			y += slotH + g
		}
	case LayoutGrid:
		cols := igo.Columns
		if cols <= 0 {
			cols = 1
		}
		rows := (n + cols - 1) / cols
		slotW := (igo.bounds.Width - p*2 - g*float32(cols-1)) / float32(cols)
		slotH := igo.Style.CellHeight
		if slotH == 0 {
			slotH = (igo.bounds.Height - p*2 - g*float32(rows-1)) / float32(rows)
		}
		for i, child := range igo.children {
			col := i % cols
			row := i / cols
			x := igo.bounds.X + p + float32(col)*(slotW+g)
			y := igo.bounds.Y + p + float32(row)*(slotH+g)
			child.SetBounds(rl.NewRectangle(x, y, slotW, slotH))
		}
	}
}

func NewContainer(bound rl.Rectangle, opts ...func(*Style)) Container {
	s := DefaultStyle()
	for _, opt := range opts {
		opt(&s)
	}

	inset := s.Margin
	if s.Border != nil {
		inset += s.Border.Thickness
	}

	w := bound.Width
	h := bound.Height
	if s.Width != 0 {
		w = s.Width
	}
	if s.Height != 0 {
		h = s.Height
	}

	adjusted := rl.NewRectangle(
		bound.X+inset+s.OffsetX,
		bound.Y+inset+s.OffsetY,
		w-inset*2,
		h-inset*2,
	)
	return Container{bounds: adjusted, Layout: s.Layout, Style: s, Columns: s.Columns}
}

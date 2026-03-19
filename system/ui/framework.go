package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout int

const (
	LayoutNone Layout = iota
	LayoutHorizontal
	LayoutVertical
	LayoutGrid
)

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
}

func (c *Container) Draw(ctx DrawContext) {
	rl.DrawRectangleLinesEx(c.bounds, c.Style.Border.Thickness, c.Style.Border.Color)
	for _, child := range c.Children() {
		child.Draw(ctx)
	}
}

func (c *Container) Bounds() rl.Rectangle     { return c.bounds }
func (c *Container) SetBounds(b rl.Rectangle) { c.bounds = b; c.applyLayout() }
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
		slotH := (c.bounds.Height - p*2 - g*float32(rows-1)) / float32(rows)
		for i, child := range c.children {
			col := i % cols
			row := i / cols
			x := c.bounds.X + p + float32(col)*(slotW+g)
			y := c.bounds.Y + p + float32(row)*(slotH+g)
			child.SetBounds(rl.NewRectangle(x, y, slotW, slotH))
		}
	}
}

func NewContainer(bound rl.Rectangle, layout Layout, opts ...func(*Style)) Container {
	s := DefaultStyle()
	for _, opt := range opts {
		opt(&s)
	}

	inset := s.Margin + s.Padding
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
	return Container{bounds: adjusted, Layout: layout, Style: s}
}

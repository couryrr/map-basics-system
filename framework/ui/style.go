package framework

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func DefaultStyle() Style {
	return Style{
		Padding: 1,
		Margin:  1,
		Width:   1,
		Height:  1,
		Border: &Border{
			Thickness: 1,
			Color: rl.Gray,
		},
	}
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

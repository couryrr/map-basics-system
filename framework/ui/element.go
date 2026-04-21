package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type UiEvent interface {
	IsConsumed() bool
	Consume()
	GetPosition() *rl.Vector2
}

type Prop struct {
	Style Style
	Text  string
}

type PropFn func(ctx *UiContext) Prop

type Element struct {
	Id       string
	bounds   rl.Rectangle
	propFn   PropFn
	parent   Drawable
	children []Drawable
	onClick  func(e UiEvent)
}

func NewElement() Element {
	elm := Element{Id: uuid.NewString()}
	return elm
}

type TypedElement[T any] struct {
	Element
	Type *T
}

func NewTypedElement[T any](bound rl.Rectangle, prop *T) TypedElement[T] {
	telm := TypedElement[T]{
		Element: NewElement(),
		Type:    prop,
	}
	return telm
}

func (elm *Element) WithPropFn(propFn PropFn) {
	elm.propFn = propFn
}

func (elm *Element) Bounds() rl.Rectangle { return elm.bounds }
func (elm *Element) SetBounds(bounds rl.Rectangle) {
	elm.bounds = bounds
}

// TODO: this math is from an LLM verify
func (elm *Element) ComputeBounds(ctx *UiContext) {
	container := elm.Bounds()
	if elm.Parent() != nil {
		container = elm.Parent().Bounds()
	}
	if elm.propFn != nil {
		props := elm.propFn(ctx)
		inset := props.Style.Margin
		if props.Style.Border != nil {
			inset += props.Style.Border.Thickness
		}
		w := container.Width
		h := container.Height
		if props.Style.Width != 0 {
			w = props.Style.Width
		}
		if props.Style.Height != 0 {
			h = props.Style.Height
		}
		elm.bounds = rl.NewRectangle(
			container.X+inset+props.Style.OffsetX,
			container.Y+inset+props.Style.OffsetY,
			w-inset*2,
			h-inset*2,
		)
	}
	for _, child := range elm.Children() {
		child.ComputeBounds(ctx)
	}
	elm.applyLayout(ctx)
}

func (elm *Element) Parent() Drawable          { return elm.parent }
func (elm *Element) SetParent(parent Drawable) { elm.parent = parent }
func (elm *Element) Children() []Drawable      { return elm.children }
func (elm *Element) AddChild(e Drawable) {
	e.SetParent(elm)
	elm.children = append(elm.children, e)
}

func (elm *Element) OnClick(fn func(e UiEvent)) {
	elm.onClick = fn
}

func (elm *Element) hitTest(point *rl.Vector2) string {
	for i := len(elm.Children()) - 1; i >= 0; i-- {
		if hit := elm.children[i].hitTest(point); hit != "" {
			return hit
		}
	}
	if rl.CheckCollisionPointRec(*point, elm.bounds) {
		return elm.Id
	}
	return ""
}

func (elm *Element) bubble(e UiEvent) {
	if elm.onClick != nil {
		elm.onClick(e)
	}

	if !e.IsConsumed() && elm.parent != nil {
		elm.parent.bubble(e)
	}
}

func (elm *Element) draw(ctx *UiContext) {
	if elm.propFn != nil {
		props := elm.propFn(ctx)
		rl.DrawRectangleLinesEx(elm.Bounds(), props.Style.Border.Thickness, props.Style.Border.Color)

		if fs := props.Style.Font; fs != nil && props.Text != "" {
			pos := fs.Position(props.Text, elm.bounds)
			rl.DrawTextEx(fs.Font, props.Text, pos, fs.Size, fs.Spacing, fs.Color)
		}

		for _, child := range elm.Children() {
			child.draw(ctx)
		}
	}
}

// TODO: this math is from an LLM verify
func (elm *Element) applyLayout(ctx *UiContext) {
	if elm.propFn != nil {
		props := elm.propFn(ctx)
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

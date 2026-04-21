package ui

type UiManager struct {
	root    Drawable
	Hovered Drawable
}

func (um *UiManager) Update(event UiEvent) {
	if um.root != nil && event != nil {
		if event.GetPosition() != nil {
			hovered := um.root.hitTest(event.GetPosition())
			if hovered != nil {
				um.Hovered = hovered
				event.Consume()
			}
		}
	}
}

func (um *UiManager) Draw() {
	ctx := &UiContext{
		Hovered: um.Hovered,
	}

	um.root.draw(ctx)
}

func NewUiManager(root Drawable) *UiManager {
	return &UiManager{
		root:    root,
		Hovered: nil,
	}
}

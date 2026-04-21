package ui

type UiManager struct {
	root    Drawable
	HoveredId string
	dirty bool
}

func (um *UiManager) Update(event UiEvent) {
	if um.dirty {
		um.root.ComputeBounds(&UiContext{
			HoveredId: um.HoveredId,
		})
		um.dirty = false
	}
	if um.root != nil && event != nil {
		if event.GetPosition() != nil {
			hoveredId := um.root.hitTest(event.GetPosition())
			if hoveredId != "" {
				um.HoveredId = hoveredId
				event.Consume()
			}
		}
	}
}

func (um *UiManager) Draw() {
	ctx := &UiContext{
		HoveredId: um.HoveredId,
	}

	um.root.draw(ctx)
}

func NewUiManager(root Drawable) *UiManager {
	return &UiManager{
		root:    root,
		dirty: true,
	}
}

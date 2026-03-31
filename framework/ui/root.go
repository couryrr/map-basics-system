package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Root struct {
	bounds rl.Rectangle
	child  Drawable
}

func (r *Root) SetChild(child Drawable){
	child.SetBounds(r.bounds)
	r.child = child
}

func (r *Root)Draw(){
	r.child.Draw()
}

func (r *Root) Click(point rl.Vector2) {
    hit := r.child.hitTest(point)
    if hit == nil {
        return
    }
    hit.bubble(&UiEvent{point: point})
}

func NewRoot(screen rl.Rectangle) Root {
	root := Root{}
	root.bounds = screen
	return root
}

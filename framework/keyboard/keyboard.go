package keyboard

import (
	"github.com/couryrr/map-basics-system/framework/queue"
	"github.com/couryrr/map-basics-system/system/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MouseMoved queue.EventKind = iota
	MouseClicked
	KeyPressed
)

const (
	TopicScreenToggle     queue.Topic = "screen.toggle"
	TopicInputMove        queue.Topic = "input.move"
	TopicInputRotate      queue.Topic = "input.rotate"
	TopicInputRotateReset queue.Topic = "input.rotate.reset"
	TopicInputZoom        queue.Topic = "input.zoom"
)

type InputEvent struct {
	Type     string
	State    string
	AltKey   string
	Key      string
	Position *rl.Vector2
	Consumed bool
}

func (ie *InputEvent) GetPosition() *rl.Vector2 { return ie.Position }
func (ie *InputEvent) IsConsumed() bool         { return ie.Consumed }
func (ie *InputEvent) Consume()                 { ie.Consumed = true }

func HandleInput(rCtx *renderer.RenderContext) *InputEvent {
	event := &InputEvent{
		Position: rCtx.ScreenToVirtual(rl.GetMousePosition()),
		Consumed: false,
	}
	if rl.IsKeyPressed(rl.KeyF11) {
		return event
	}
	if rl.IsKeyPressed(rl.KeyE) {
		return event
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		return event
	}
	if rl.IsKeyPressed(rl.KeyC) {
		return event
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		event.Type = "mouse"
		event.State = "clicked"
		// event.Key = rl.MouseButtonLeft
		event.Position = rCtx.ScreenToVirtual(rl.GetMousePosition())
		return event
	}

	if wheel := rl.GetMouseWheelMove(); wheel != 0 {
		return event
	}

	direction := rl.Vector2Zero()

	if rl.IsKeyDown(rl.KeyW) {
		direction.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		direction.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		direction.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		direction.X += 1
	}

	if direction.X != 0 && direction.Y != 0 {
		direction = rl.Vector2Scale(direction, 0.7071)
	}

	if !rl.Vector2Equals(direction, rl.Vector2Zero()) {
		return event
	}

	return event
}

package keyboard

import (
	"github.com/couryrr/map-basics-system/framework/queue"
	"github.com/couryrr/map-basics-system/framework/ui"
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
	TopicInputCursorMoved queue.Topic = "input.cursor.moved"
)

func HandleInput(ui *ui.Root, broker *queue.EventQueue, rCtx *renderer.RenderContext) {
	if rl.IsKeyPressed(rl.KeyF11) {
		broker.Push(&queue.Event{})
	}
	if rl.IsKeyPressed(rl.KeyE) {
		broker.Push(&queue.Event{})
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		broker.Push(&queue.Event{})
	}
	if rl.IsKeyPressed(rl.KeyC) {
		broker.Push(&queue.Event{
			Kind: 1,
			Position:  rCtx.ScreenToVirtual(rl.GetMousePosition()),
		})
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft){
		ui.Click(rCtx.ScreenToVirtual(rl.GetMousePosition()))
	}
	delta := rl.GetMouseDelta()
	if !rl.Vector2Equals(delta, rl.Vector2Zero()) {
		// broker.Send(TopicInputCursorMoved, pubsub.Event{Data: framework.InputEvent{
		// 	Position:  rCtx.ScreenToVirtual(rl.GetMousePosition()),
		// 	EventType: framework.MouseHoverEvent,}})
	}

	if wheel := rl.GetMouseWheelMove(); wheel != 0 {
		broker.Push(&queue.Event{})
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
		broker.Push(&queue.Event{})
	}
}

package controller

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
	"github.com/couryrr/map-basics-system/system/renderer"
	"github.com/couryrr/map-basics-system/system/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TopicScreenToggle     pubsub.Topic = "screen.toggle"
	TopicInputMove        pubsub.Topic = "input.move"
	TopicInputRotate      pubsub.Topic = "input.rotate"
	TopicInputRotateReset pubsub.Topic = "input.rotate.reset"
	TopicInputZoom        pubsub.Topic = "input.zoom"
	TopicInputCursorMoved pubsub.Topic = "input.cursor.moved"
)

func HandleInput(broker *pubsub.Broker, rCtx *renderer.RenderContext) {
	if rl.IsKeyPressed(rl.KeyF11) {
		broker.Send(TopicScreenToggle, pubsub.Message{})
	}
	if rl.IsKeyPressed(rl.KeyE) {
		broker.Send(TopicInputRotate, pubsub.Message{Data: float32(90)})
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		broker.Send(TopicInputRotate, pubsub.Message{Data: float32(-90)})
	}
	if rl.IsKeyPressed(rl.KeyC) {
		broker.Send(TopicInputRotateReset, pubsub.Message{})
	}

	delta := rl.GetMouseDelta()
	if !rl.Vector2Equals(delta, rl.Vector2Zero()) {
		broker.Send(TopicInputCursorMoved, pubsub.Message{ Data: ui.MouseEvent{
			Position: rCtx.ScreenToVirtual(rl.GetMousePosition()),
			Event:    ui.MouseHoveEvent,
		}})
	}

	if wheel := rl.GetMouseWheelMove(); wheel != 0 {
		broker.Send(TopicInputZoom, pubsub.Message{Data: wheel})
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
		broker.Send(TopicInputMove, pubsub.Message{Data: direction})
	}
}

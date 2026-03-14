package system

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TopicScreenToggle      Topic = "screen.toggle"
	TopicPlayerMove        Topic = "player.move"
	TopicPlayerRotate      Topic = "player.rotate"
	TopicPlayerRotateReset Topic = "player.rotate.reset"
	TopicPlayerZoom        Topic = "player.zoom"
)

func HandleInput(broker Broker) {
	if rl.IsKeyPressed(rl.KeyF11) {
		broker.Send(TopicScreenToggle, Message{})
	}
	if rl.IsKeyPressed(rl.KeyE) {
		broker.Send(TopicPlayerRotate, Message{Data: float32(90)})
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		broker.Send(TopicPlayerRotate, Message{Data: float32(-90)})
	}
	if rl.IsKeyPressed(rl.KeyC) {
		broker.Send(TopicPlayerRotateReset, Message{})
	}
	if wheel := rl.GetMouseWheelMove(); wheel != 0 {
		broker.Send(TopicPlayerZoom, Message{Data: wheel})
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
		broker.Send(TopicPlayerMove, Message{Data: direction})
	}
}

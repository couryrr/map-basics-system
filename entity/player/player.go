package player

import (
	"github.com/couryrr/map-basics-system/framework/queue"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	zoomStep                                   = float32(0.25)
	zoomMin                                    = float32(0.25)
	zoomMax                                    = float32(3.0)
	TopicPlayerHotbarSlotSelected queue.Topic = "player.hotbar.slot.selected"
)

type Hotbar struct {
	Slots      [10]string
	ActiveSlot int32
}

func (h *Hotbar) SlotItem(i int32) string { return h.Slots[i] }
func (h *Hotbar) GetActiveSlot() int32   { return h.ActiveSlot }

func (h *Hotbar) SetActive(slot int32) {
	h.ActiveSlot = slot
}
func (h *Hotbar) Assign(slot int32, itemID string) {
	h.Slots[slot] = itemID
}
func (h *Hotbar) Clear() {
	h.ActiveSlot = -1
}

type Player struct {
	Position  rl.Vector2
	Rotation  float32
	Speed     float32
	ZoomLevel float32
	Hotbar    Hotbar
}

func NewPlayer(start rl.Vector2) Player {
	slots := [10]string{
		"drone",
		"stockpile",
	}
	slots[5] = "sieve"
	i := int32(2)
	return Player{
		Position:  start,
		Rotation:  0,
		Speed:     400,
		ZoomLevel: 1.0,
		Hotbar: Hotbar{
			ActiveSlot: i,
			Slots:      slots,
		},
	}
}

// func (player *Player) Rotate(message pubsub.Event) {
// 	if rotation, ok := message.Data.(float32); ok {
// 		player.Rotation += rotation
// 	}
// }

// func (player *Player) RotateReset(message queue.Event) {
// 	player.Rotation = 0
// }

// func (player *Player) Zoom(message pubsub.Event) {
// 	if delta, ok := message.Data.(float32); ok {
// 		player.ZoomLevel += delta * zoomStep
// 		player.ZoomLevel = rl.Clamp(player.ZoomLevel, zoomMin, zoomMax)
// 	}
// }

// func (player *Player) Move(message pubsub.Event) {
// 	if movement, ok := message.Data.(rl.Vector2); ok {
// 		delta := rl.GetFrameTime()
// 		angle := -player.Rotation * rl.Deg2rad
// 		rotated := rl.Vector2Rotate(movement, angle)
// 		player.Position.X += rotated.X * player.Speed * delta
// 		player.Position.Y += rotated.Y * player.Speed * delta
// 	}
// }

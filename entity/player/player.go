package player

import (
	"github.com/couryrr/map-basics-system/system/pubsub"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	zoomStep                                   = float32(0.25)
	zoomMin                                    = float32(0.25)
	zoomMax                                    = float32(3.0)
	TopicPlayerHotbarSlotSelected pubsub.Topic = "player.hotbar.slot.selected"
)

type HotbarInteractionMessage struct {
	Slot   int32
	ItemId string
}

type Hotbar struct {
	Slots      [6]string
	ActiveSlot *int32
}

func (h *Hotbar) SetActive(slot *int32) {
	h.ActiveSlot = slot
}
func (h *Hotbar) Assign(slot int32, itemID string) {
	h.Slots[slot] = itemID
}
func (h *Hotbar) Clear() {
	h.ActiveSlot = nil
}
func (h *Hotbar) ActiveItem() string {
	return h.ActiveItem()
}

type Player struct {
	Position  rl.Vector2
	Rotation  float32
	Speed     float32
	ZoomLevel float32
	Hotbar    Hotbar
}

func NewPlayer(start rl.Vector2) Player {
	slots := [6]string{
		"1x1",
		"1x2",
	}
	slots[5] = "2x2"
	return Player{
		Position:  start,
		Rotation:  0,
		Speed:     400,
		ZoomLevel: 1.0,
		Hotbar: Hotbar{
			ActiveSlot: nil,
			Slots:      slots,
		},
	}
}

func (player *Player) AddHotbarItem(message pubsub.Message) {
	if hbim, ok := message.Data.(HotbarInteractionMessage); ok {
		player.Hotbar.Assign(hbim.Slot, hbim.ItemId)
	}
}

func (player *Player) SelectHotbarItem(message pubsub.Message) {
	if hbim, ok := message.Data.(HotbarInteractionMessage); ok {
		player.Hotbar.SetActive(&hbim.Slot)
	}
}

func (player *Player) HighlightHotbarItem(message pubsub.Message) {
	if index, ok := message.Data.(*int32); ok {
		rl.TraceLog(rl.LogInfo, "AHHHHHHHHHHH: %d", *index)
		player.Hotbar.SetActive(index)
	}
}

func (player *Player) Rotate(message pubsub.Message) {
	if rotation, ok := message.Data.(float32); ok {
		player.Rotation += rotation
	}
}

func (player *Player) RotateReset(message pubsub.Message) {
	player.Rotation = 0
}

func (player *Player) Zoom(message pubsub.Message) {
	if delta, ok := message.Data.(float32); ok {
		player.ZoomLevel += delta * zoomStep
		player.ZoomLevel = rl.Clamp(player.ZoomLevel, zoomMin, zoomMax)
	}
}

func (player *Player) Move(message pubsub.Message) {
	if movement, ok := message.Data.(rl.Vector2); ok {
		delta := rl.GetFrameTime()
		angle := -player.Rotation * rl.Deg2rad
		rotated := rl.Vector2Rotate(movement, angle)
		player.Position.X += rotated.X * player.Speed * delta
		player.Position.Y += rotated.Y * player.Speed * delta
	}
}

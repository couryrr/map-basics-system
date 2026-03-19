package world

import (
	"encoding/json"
	"fmt"
	"image/color"
	"iter"
	"maps"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameItem struct {
	Name       string     `json:"name"`
	Category   string     `json:"category"`
	Shape      string     `json:"shape"`
	Dimensions rl.Vector2 `json:"dimensions"`
	Color      color.RGBA `json:"color"`
	Sprite     string     `json:"sprite"`
}

type Registry struct {
	items map[string]GameItem
}

func (r *Registry) GetItems() iter.Seq2[string, GameItem] {
	return maps.All(r.items)
}

func (r *Registry) GetItemById(itemId string) (*GameItem, error) {
	if gameItem, ok := r.items[itemId]; ok {
		return &gameItem, nil
	}

	return nil, fmt.Errorf("Item not found for: %s", itemId)
}

func NewRegistry() Registry {
	directoryJson := "./assets/directory.json"
	file, err := os.ReadFile(directoryJson)
	if err != nil {
		rl.TraceLog(rl.LogInfo, "This is it: %v", err)
		panic(err)
	}
	registry := Registry{}
	err = json.Unmarshal(file, &registry.items)
	if err != nil {
		rl.TraceLog(rl.LogInfo, "This is it 2: %v", err)
		panic(err)
	}

	return registry
}

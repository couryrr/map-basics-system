package resource

import (
	"encoding/json"
	"fmt"
	"image/color"
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

type Directory struct {
	Items map[string]GameItem `json:"items"`
}

func (d *Directory) GetItemById(itemId string) (*GameItem, error) {
	if gameItem, ok := d.Items[itemId]; ok {
		return &gameItem, nil
	}

	return nil, fmt.Errorf("Item not found for: %s", itemId)
}

func NewDirectory() *Directory {

	//TODO: This is going to be read from a directory based on os.
	directoryJson := "./assets/directory.json"
	file, err := os.ReadFile(directoryJson)
	if err != nil {
		rl.TraceLog(rl.LogInfo, "This is it: %v", err)
		panic(err)
	}

	directory := &Directory{}
	err = json.Unmarshal(file, directory)
	if err != nil {
		rl.TraceLog(rl.LogInfo, "This is it 2: %v", err)
		panic(err)
	}
	rl.TraceLog(rl.LogInfo, "%v", directory)

	return directory
}

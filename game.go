package main

import (
	"github.com/couryrr/map-basics-system/system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSettings struct {
	ScreenSetting system.ScreenSetting
}

type Game struct {
	GameSettings GameSettings
}

func (game Game) LoadGame(){}
func (game Game) SaveGame(){}

func CreateGame(windowedScreenSize, screenSize rl.Vector2) Game {
	return Game{
		GameSettings: GameSettings{
			ScreenSetting: system.CreateScreenSetting(screenSize, windowedScreenSize),
		},
	}
}

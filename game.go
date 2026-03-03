package main

type GameSettings struct {
	WindowedScreenWidth    float32
	WindowededScreenHeight float32
	ScreenHeight           float32
	ScreenWidth            float32
}

type Game struct {
	settings GameSettings
}

func CreateGamePlease() Game {
	return Game{
		settings: GameSettings{
			WindowedScreenWidth: 1920,
			WindowededScreenHeight: 1080,
			ScreenWidth: 1920,
			ScreenHeight: 1080,
		},
	}
}

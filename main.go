package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"bodies"
)

const (
	SCREEN_W = 1200
	SCREEN_H = 1000
)

func main() {
	rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Tankers")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)

	var players []*bodies.Tank

	tank := bodies.NewTank(len(players), rl.NewVector2(0, 0))
	players = append(players, tank)
	var dt float32

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()
		for i := 0; i < len(players); i++ {
			players[i].Update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)

		// mouseV := rl.GetMousePosition()
		// println(mouseV.X, mouseV.Y)

		for i := 0; i < len(players); i++ {
			players[i].Draw()
		}

		rl.EndDrawing()
	}
}

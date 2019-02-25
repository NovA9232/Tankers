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
	bodies.SCREEN_W = SCREEN_W
	bodies.SCREEN_H = SCREEN_H

	rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Tankers")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)

	var players []*bodies.Tank

	tank := bodies.NewTank(len(players), rl.NewVector2(400, 500))
	players = append(players, tank)
	var dt float32

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()
		for i := 0; i < len(players); i++ {
			players[i].Update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for i := 0; i < len(players); i++ {
			players[i].Draw()
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

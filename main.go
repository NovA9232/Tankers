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

	camera := rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 2)

	var players []*bodies.Tank

	tank := bodies.NewTank(len(players), rl.NewVector2(200, 200))
	players = append(players, tank)

	var dt float32

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()
		for i := 0; i < len(players); i++ {
			players[i].Update(dt)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		for i := 0; i < len(players); i++ {
			players[i].Draw()
		}
		rl.EndMode2D()  // Draw players etc using camera.

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

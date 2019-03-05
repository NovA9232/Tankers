package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	crosshairL = 7   // Length of each part of the crosshair
	crosshairW = 2    // Width of each part of the crosshair. Works best with multiples of 2
	crosshairGap = 3 // Gap between each part of crosshair
)

var (
	crosshairCol rl.Color = rl.Green
)

func drawCrosshair() {
	mouseX, mouseY := rl.GetMouseX(), rl.GetMouseY()
	posYMid := mouseY - crWDiv2
	posXMid := mouseX - crWDiv2

	rl.DrawRectangle(mouseX - crosshairGap - crosshairL, posYMid, crosshairL, crosshairW, crosshairCol)   // Left part
  rl.DrawRectangle(mouseX + crosshairGap, posYMid, crosshairL, crosshairW, crosshairCol)								// Right part

	rl.DrawRectangle(posXMid, mouseY - crosshairGap - crosshairL, crosshairW, crosshairL, crosshairCol)   // Top part
	rl.DrawRectangle(posXMid, mouseY + crosshairGap, crosshairW, crosshairL, crosshairCol)								// Bottom part
}



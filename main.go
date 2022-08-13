package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	GRID_WIDTH  int32 = 48
	GRID_HEIGHT int32 = 48
	CELL_SIZE   int32 = 12

	SCREEN_WIDTH  int32 = CELL_SIZE * GRID_WIDTH
	SCREEN_HEIGHT int32 = CELL_SIZE * GRID_HEIGHT

	FRAME_RATE int32 = 60

	NEXT_GEN_INTERVAL_INITIAL float32 = 1.0 / 6.0
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Game of Life")

	rl.SetTargetFPS(FRAME_RATE)

	camera := rl.Camera2D{
		Offset:   rl.Vector2{X: 0.0, Y: 0.0},
		Target:   rl.Vector2{X: 0.0, Y: 0.0},
		Rotation: 0,
		Zoom:     1,
	}

	board := InitBoard()

	frame := 0
	pause := true

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		if rl.IsKeyPressed(rl.KeyR) {
			board = nil // marking the old board as nil, not needed
			board = InitBoard()
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsMouseButtonDown(rl.MouseRightButton) {
			x := rl.GetMouseX() / CELL_SIZE
			y := rl.GetMouseY() / CELL_SIZE

			if !(x < 0 || x >= GRID_WIDTH || y < 0 || y >= GRID_HEIGHT) {
				if rl.IsMouseButtonDown(rl.MouseRightButton) {
					// kill cell that was clicked on
					board.cells[((y * GRID_WIDTH) + x)].alive = false
					board.SetNextStatus(x, y, false)
				} else {
					// paint cells with left mouse button
					board.cells[((y * GRID_WIDTH) + x)].alive = true
				}
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		var y int32 = 0
		var x int32 = 0
		for y = 0; y < GRID_HEIGHT; y++ {
			for x = 0; x < GRID_WIDTH; x++ {
				if board.CheckStatus(x, y) {
					rl.DrawRectangle(x*CELL_SIZE, y*CELL_SIZE, CELL_SIZE, CELL_SIZE, rl.White)
				}
				rl.DrawRectangleLines(x*CELL_SIZE, y*CELL_SIZE, 1, 1, rl.Gray)
			}
		}
		rl.EndMode2D()

		rl.DrawFPS(SCREEN_WIDTH-80, 0)

		if pause {
			rl.DrawText("Paused", 2, 2, 40, rl.Red)
		} else {
			if frame == int(float32(FRAME_RATE)*NEXT_GEN_INTERVAL_INITIAL) {
				board.PlayRound()
				frame = 1
			}
			frame++
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	gravity      = 400.0
	jumpSpeed    = -300.0
	pipeSpeed    = 200.0
	pipeWidth    = 50
	pipeGap      = 150
	frameSpeed   = 0.1 // How fast the bird animation changes frames
	saveFile     = "savegame.json"
)

type Bird struct {
	Position rl.Vector2
	Speed    float32
	Frame    int // Current frame for bird animation
}

type Pipe struct {
	Position rl.Vector2
	GapY     float32
}

type GameState struct {
	Bird  Bird
	Pipe  Pipe
	Score int
}

func NewBird() Bird {
	return Bird{
		Position: rl.NewVector2(100, screenHeight/2),
		Speed:    0,
		Frame:    0, // Start at the first frame
	}
}

func NewPipe() Pipe {
	gapY := float32(rand.Intn(screenHeight-pipeGap)) + float32(pipeGap)/2
	return Pipe{
		Position: rl.NewVector2(screenWidth, 0),
		GapY:     gapY,
	}
}

func ResetGame(b *Bird, p *Pipe, score *int) {
	*b = NewBird()
	*p = NewPipe()
	*score = 0
}

// Save game state to a file
func SaveGame(bird Bird, pipe Pipe, score int) {
	state := GameState{
		Bird:  bird,
		Pipe:  pipe,
		Score: score,
	}

	file, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		fmt.Printf("Error saving game: %v\n", err)
		return
	}

	err = ioutil.WriteFile(saveFile, file, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

// Load game state from a file
func LoadGame(b *Bird, p *Pipe, score *int) {
	if _, err := os.Stat(saveFile); os.IsNotExist(err) {
		fmt.Println("No save file found")
		return
	}

	file, err := ioutil.ReadFile(saveFile)
	if err != nil {
		fmt.Printf("Error reading save file: %v\n", err)
		return
	}

	var state GameState
	err = json.Unmarshal(file, &state)
	if err != nil {
		fmt.Printf("Error loading game: %v\n", err)
		return
	}

	*b = state.Bird
	*p = state.Pipe
	*score = state.Score
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Flappy Bird with Animation")
	defer rl.CloseWindow()

	rand.Seed(time.Now().UnixNano())
	rl.SetTargetFPS(60)

	// Load bird frames
	frames := []rl.Texture2D{
		rl.LoadTexture("frame-1.png"),
		rl.LoadTexture("frame-2.png"),
		rl.LoadTexture("frame-3.png"),
		rl.LoadTexture("frame-4.png"),
		rl.LoadTexture("frame-5.png"),
		rl.LoadTexture("frame-6.png"),
		rl.LoadTexture("frame-7.png"),
		rl.LoadTexture("frame-8.png"),
	}
	defer func() {
		for _, frame := range frames {
			rl.UnloadTexture(frame)
		}
	}()

	// Initialize bird, pipe, and score
	bird := NewBird()
	pipe := NewPipe()
	score := 0
	gameOver := false

	var frameCounter float32     // Used to track time for animation frame switching
	scaleFactor := float32(0.07) // Scale down the bird texture to 7%

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if gameOver {
			// Game over screen
			rl.DrawText("Game Over!", screenWidth/2-rl.MeasureText("Game Over!", 40)/2, screenHeight/2-50, 40, rl.Red)
			rl.DrawText("Press 'R' to restart", screenWidth/2-rl.MeasureText("Press 'R' to restart", 20)/2, screenHeight/2, 20, rl.Black)
			rl.DrawText("Press 'S' to save", screenWidth/2-rl.MeasureText("Press 'S' to save", 20)/2, screenHeight/2+30, 20, rl.Black)
			rl.DrawText("Press 'L' to load", screenWidth/2-rl.MeasureText("Press 'L' to load", 20)/2, screenHeight/2+60, 20, rl.Black)

			if rl.IsKeyPressed(rl.KeyR) {
				ResetGame(&bird, &pipe, &score)
				gameOver = false
			}
		} else {
			// Control bird movement
			if rl.IsKeyPressed(rl.KeyW) {
				bird.Speed = jumpSpeed
			}
			bird.Speed += gravity * rl.GetFrameTime()
			bird.Position.Y += bird.Speed * rl.GetFrameTime()

			if bird.Position.Y < 0 {
				bird.Position.Y = 0
			} else if bird.Position.Y > screenHeight {
				gameOver = true
			}

			pipe.Position.X -= pipeSpeed * rl.GetFrameTime()

			if pipe.Position.X < -pipeWidth {
				pipe = NewPipe()
				score++
			}

			if bird.Position.X+20 > pipe.Position.X && bird.Position.X < pipe.Position.X+pipeWidth {
				if bird.Position.Y < pipe.GapY-pipeGap/2 || bird.Position.Y > pipe.GapY+pipeGap/2 {
					gameOver = true
				}
			}

			// Animate bird
			frameCounter += rl.GetFrameTime()
			if frameCounter >= frameSpeed {
				bird.Frame++
				if bird.Frame >= len(frames) {
					bird.Frame = 0 // Loop back to the first frame
				}
				frameCounter = 0
			}

			// Draw the bird with the current frame, scaled
			rl.DrawTextureEx(frames[bird.Frame], bird.Position, 0, scaleFactor, rl.White)

			// Draw pipe
			rl.DrawRectangle(int32(pipe.Position.X), 0, int32(pipeWidth), int32(pipe.GapY-pipeGap/2), rl.Green)
			rl.DrawRectangle(int32(pipe.Position.X), int32(pipe.GapY+pipeGap/2), int32(pipeWidth), int32(screenHeight-int(pipe.GapY+pipeGap/2)), rl.Green)

			// Draw score
			rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 10, 20, rl.Black)
		}

		// Save game state
		if rl.IsKeyPressed(rl.KeyS) {
			SaveGame(bird, pipe, score)
			rl.DrawText("Game Saved!", screenWidth/2-rl.MeasureText("Game Saved!", 20)/2, screenHeight/2+90, 20, rl.DarkGreen)
		}

		// Load game state
		if rl.IsKeyPressed(rl.KeyL) {
			LoadGame(&bird, &pipe, &score)
			gameOver = false
			rl.DrawText("Game Loaded!", screenWidth/2-rl.MeasureText("Game Loaded!", 20)/2, screenHeight/2+90, 20, rl.DarkGreen)
		}

		rl.EndDrawing()
	}
}

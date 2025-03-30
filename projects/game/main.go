package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ----------------------------------------------
// Game Constants
// ----------------------------------------------
const (
	windowWidth  = 800
	windowHeight = 450
)

// ----------------------------------------------
// Global State
// ----------------------------------------------
var (
	gameIsRunning   = true
	backgroundColor = rl.NewColor(147, 211, 196, 255)

	// Textures
	grassTexture rl.Texture2D
	playerSprite rl.Texture2D

	// player rellated
	playerSrc   rl.Rectangle
	playerDest  rl.Rectangle
	playerSpeed float32 = 5.0

	// Music
	bgmPaused bool
	bgm       rl.Music

	// Camera
	cam rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(grassTexture, 100, 50, rl.White)
	// rl.DrawTexture(playerSprite, 200, 200, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyH) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyK) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyL) {
		playerDest.X += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyJ) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		bgmPaused = !bgmPaused
	}
}

func update() {
	gameIsRunning = !rl.WindowShouldClose()

	rl.UpdateMusicStream(bgm)
	if bgmPaused {
		rl.PauseMusicStream(bgm)
	} else {
		rl.ResumeMusicStream(bgm)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()

	rl.EndDrawing()
}

func gameInit() {
	rl.InitWindow(windowWidth, windowHeight, "Sproutlings")
	rl.SetExitKey('Q')
	rl.SetTargetFPS(60)

	// Load Textures
	grassTexture = rl.LoadTexture("res/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("res/Characters/BasicCharakterSpritesheet.png")

	// Player Related
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	// Music
	rl.InitAudioDevice()
	bgm = rl.LoadMusicStream("res/bgm.mp3")
	bgmPaused = false
	rl.PlayMusicStream(bgm)

	// Camera
	cam = rl.NewCamera2D(rl.NewVector2(float32(windowWidth/2), float32(windowHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.0)
}

func gameQuit() {
	// Textures
	rl.UnloadTexture(grassTexture)
	rl.UnloadTexture(playerSprite)

	// Music
	rl.UnloadMusicStream(bgm)
	rl.CloseAudioDevice()

	// Window
	rl.CloseWindow()
}

// ----------------------------------------------
// Entry Point
// ----------------------------------------------
func main() {
	gameInit()

	for gameIsRunning {
		input()
		update()
		render()
	}
	gameQuit()
}

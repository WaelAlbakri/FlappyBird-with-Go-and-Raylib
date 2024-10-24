# Flappy Bird with Go and Raylib

This repository contains a recreation of the classic **Flappy Bird** game using **Raylib** and **Go**. The game involves controlling a bird to pass through randomly placed gaps in moving pipes, while tracking the score and handling game-over conditions.

## Game Preview

Watch the game in action: [Flappy Bird Gameplay](https://www.youtube.com/watch?v=RwHirAC7rYs)

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [How to Run](#how-to-run)
- [Controls](#controls)
- [Scoring](#scoring)
- [Save and Load Feature](#save-and-load-feature)

## Features

- **Character Control**: 
  - Control the bird using 'W' to jump (move up) and 'S' to move down.
  
- **Dynamic Obstacles**: 
  - Pipes with randomly placed gaps appear on the right side of the screen and move to the left.
  
- **Game Over & Restart**: 
  - Colliding with a pipe ends the game with a "Game Over" screen. Players can press 'R' to restart the game.

- **Score Tracking**: 
  - Score increases by 1 for each set of pipes the player successfully passes through. The score is displayed in the top-left corner of the screen.

- **Save/Load Feature**: 
  - Players can save their current game state (bird's position, pipe's position, and score) and load it later.
  - Press `S` to save the game state, and press `L` to load the saved state.

## Requirements

- **Go** programming language.
- **Raylib** library for Go.

## Controls
W: Move bird up (jump)
S: Move bird down
R: Restart the game (on Game Over)
S: Save the game state
L: Load the saved game state

## Scoring
- Pass through the pipes to earn +1 point.
- The score is displayed at the top-left corner of the screen.
- Save and Load Feature
- Press S to save the current game state (bird position, pipe position, score).
- Press L to load the saved game state.

# Golang Scratch

[![Go Reference](https://pkg.go.dev/badge/github.com/GaryBrownEEngr/scratch.svg)](https://pkg.go.dev/github.com/GaryBrownEEngr/scratch)
[![Go CI](https://github.com/GaryBrownEEngr/scratch/actions/workflows/go.yml/badge.svg)](https://github.com/GaryBrownEEngr/scratch/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaryBrownEEngr/scratch)](https://goreportcard.com/report/github.com/GaryBrownEEngr/scratch)
[![Coverage Badge](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/GaryBrownEEngr/0a036dc69ea9afb4202e2d262fec1e1d/raw/GaryBrownEEngr_scratch_main.json)](https://github.com/GaryBrownEEngr/scratch/actions)

## A Scratch Game Engine for Golang

Golang Scratch is a reimagining of [Scratch](https://scratch.mit.edu/) created by MIT Media Lab. Every entity on the screen is a sprite that can be programmed independently.

The goals are many:

- Remove the complexity of the game loop handling
- A super easy game engine
  - See the falling turtles example
- A package for visualizing a simulation
  - See the double pendulum example

## Install

Go 1.21 or later is required.<br>
Ebitengine is the main dependency. [Check here for the system specific instructions](https://ebitengine.org/en/documents/install.html).

## Example

### Random Walk

A really simple example is to create turtles that each perform a random walk.

`main(...)` start the sim with the window size parameters. Ebiten eats the main thread, so the function to run once Ebiten starts is passed as an additional parameter.

`simStartFunc(...)` Adds one sprite image and then starts go routines for each sprite. A new sprite will be started every half second.

`turtleRandomWalk(...)` is the main function for each sprite. A new sprite is created and setup. Then it enters an infinite loop where a random velocity is added in the x and y directions. That velocity is integrated into position. The position and angle are updated. And finally, the sprite sleeps. If the sprite gets too far off screen then it is deleted.

```go
package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/GaryBrownEEngr/scratch"
	"github.com/GaryBrownEEngr/scratch/models"
	"github.com/GaryBrownEEngr/scratch/sprite"
)

func main() {
	params := scratch.ScratchParams{Width: 1000, Height: 1000}
	scratch.Start(params, simStartFunc)
}

func simStartFunc(sim models.Scratch) {
	sim.AddCostume(sprite.DecodeCodedSprite(sprite.TurtleImage), "t")
	for {
		go turtleRandomWalk(sim)
		time.Sleep(time.Millisecond * 500)
	}
}

func turtleRandomWalk(sim models.Scratch) {
	s := sim.AddSprite("")
	s.Costume("t")
	s.Visible(true)

	x, y := 0.0, 0.0
	velX, velY := 0.0, 0.0

	for {
		velX += (rand.Float64()*2 - 1) * .05
		velY += (rand.Float64()*2 - 1) * .05
		x += velX
		y += velY
		s.Pos(x, y)
		s.Angle(180.0 / math.Pi * math.Atan2(velY, velX))

		if max(math.Abs(x), math.Abs(y)) > 1000 {
			s.DeleteSprite()
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}
```

### Double Pendulum

Here is a simulation of a double pendulum.

```bash
go run github.com/GaryBrownEEngr/scratch/examples/DoublePendulum@latest
```

![Example Picture](https://github.com/GaryBrownEEngr/scratch/blob/main/examples/DoublePendulum/DoublePendulum.gif)

### Falling Turtles

Here is a really simple game where you have to click on each turtle before it reaches the bottom of the screen. At the end, your final score is displayed.

```bash
go run github.com/GaryBrownEEngr/scratch/examples/fallingturtles@latest
```

## Build Executable for Window

```bash
GOOS=windows go build ./examples/fallingturtles/
```

## Things to Research

* https://github.com/jakecoffman/cp
* https://github.com/jakecoffman/cp-examples
* https://github.com/jakecoffman/cp-ebiten
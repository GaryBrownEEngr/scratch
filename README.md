# Golang Scratch

[![Go Reference](https://pkg.go.dev/badge/github.com/GaryBrownEEngr/scratch.svg)](https://pkg.go.dev/github.com/GaryBrownEEngr/scratch)
[![Go CI](https://github.com/GaryBrownEEngr/scratch/actions/workflows/go.yml/badge.svg)](https://github.com/GaryBrownEEngr/scratch/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaryBrownEEngr/scratch)](https://goreportcard.com/report/github.com/GaryBrownEEngr/scratch)
[![Coverage Badge](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/GaryBrownEEngr/0a036dc69ea9afb4202e2d262fec1e1d/raw/GaryBrownEEngr_scratch_main.json)](https://github.com/GaryBrownEEngr/scratch/actions)

## A Scratch Game Engine for Golang

Golang Scratch is a reimagining of [Scratch](https://scratch.mit.edu/) created by MIT Media Lab. Every entity on the screen is a sprite that can be programmed independently.

## Example

### Double Pendulum

```bash
go run github.com/GaryBrownEEngr/scratch/examples/DoublePendulum@latest
```

![Example Picture](https://github.com/GaryBrownEEngr/scratch/blob/main/examples/DoublePendulum/DoublePendulum.gif)

## Build for Window

```bash
GOOS=windows go build ./examples/fallingturtles/
```

## Things to Research

* https://github.com/jakecoffman/cp
* https://github.com/jakecoffman/cp-examples
* https://github.com/jakecoffman/cp-ebiten
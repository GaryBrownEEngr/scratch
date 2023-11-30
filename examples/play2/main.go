package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/GaryBrownEEngr/scratch"
	"github.com/GaryBrownEEngr/scratch/models"
	"github.com/GaryBrownEEngr/scratch/sprite"
)

func main() {
	params := scratch.ScratchParams{Width: 1000, Height: 1000, ShowFPS: true}
	scratch.Start(params, simStartFunc)
}

func simStartFunc(sim models.Scratch) {
	sim.AddCostume(sprite.DecodeCodedSprite(sprite.TurtleImage), "t1")

	a := 0.0
	s := sim.AddSprite("mainTurtle")
	s.Costume("t1")
	s.Scale(10)
	s.Z(0)
	s.Visible(true)

	justPressedChan := sim.SubscribeToJustPressedUserInput()

	go t2(sim)

MainSpriteLoop:
	for {
		inputPressed := s.PressedUserInput()
		s.Pos(float64(inputPressed.Mouse.MouseX), float64(inputPressed.Mouse.MouseY))

		justPressed := scratch.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil {
			// use i
			if justPressed.Mouse.Left {
				a += 10
				s.Angle(a)
			}
			if justPressed.Mouse.Right {
				break MainSpriteLoop
			}
		}
		time.Sleep(time.Millisecond * 10)
	}

	sim.UnSubscribeToJustPressedUserInput(justPressedChan)
	s.DeleteSprite()
	sim.Exit()
}

func t2(sim models.Scratch) {
	s := sim.AddSprite("t2")
	s.Costume("t1")
	s.Scale(1)
	s.Z(0)
	s.Visible(true)
	cb := s.GetClickBody()
	cb.AddCircleBody(0, 0, 32)

	justPressedChan := sim.SubscribeToJustPressedUserInput()

	// MainSpriteLoop:
	for {
		t1Info := sim.GetSpriteInfo("mainTurtle")
		if !t1Info.Deleted {
			s.Pos(t1Info.X-500, t1Info.Y)
		}

		justPressed := scratch.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil && justPressed.Keys.C {
			sClone := s.Clone(fmt.Sprintf("%d%d", rand.Uint64(), rand.Uint64()))
			go tClone(sim, sClone)
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func tClone(sim models.Scratch, s models.Sprite) {
	justPressedChan := sim.SubscribeToJustPressedUserInput()
	cb := s.GetClickBody()
	// MainSpriteLoop:
	for {
		justPressed := scratch.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil && justPressed.Mouse.Left {
			if cb.IsMouseClickInBody(float64(justPressed.Mouse.MouseX), float64(justPressed.Mouse.MouseY)) {
				s.DeleteSprite()
				sim.UnSubscribeToJustPressedUserInput(justPressedChan)
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}

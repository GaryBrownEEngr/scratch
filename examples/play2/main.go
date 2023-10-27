package main

import (
	"time"

	"github.com/GaryBrownEEngr/scratch/ebitensim"
	"github.com/GaryBrownEEngr/scratch/models"
)

func main() {
	params := ebitensim.SimParams{
		Width:   1000,
		Height:  1000,
		ShowFPS: true,
	}
	ebitensim.StartSim(params, simStartFunc)
}

func simStartFunc(sim models.Sim) {
	sim.AddCostume(ebitensim.DecodeCodedSprite(ebitensim.TurtleImage), "t1")

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

		select {
		case i := <-justPressedChan:
			// use i
			if i.Mouse.Left {
				a += 10
				s.Angle(a)
			}
			if i.Mouse.Right {
				break MainSpriteLoop
			}
		default:
			// receiving from chan would block
		}
		time.Sleep(time.Millisecond * 10)
	}

	sim.UnSubscribeToJustPressedUserInput(justPressedChan)
	s.DeleteSprite()
	sim.Exit()
}

func t2(sim models.Sim) {
	s := sim.AddSprite("t2")
	s.Costume("t1")
	s.Scale(1)
	s.Z(0)
	s.Visible(true)

	// MainSpriteLoop:
	for {
		t1Info := sim.GetSpriteInfo("mainTurtle")
		s.Pos(t1Info.X-500, t1Info.Y)

		time.Sleep(time.Millisecond * 10)
	}
}

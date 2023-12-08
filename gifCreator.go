package scratch

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/GaryBrownEEngr/easygif"
	"github.com/GaryBrownEEngr/scratch/models"
)

func TakeScreenshot(sim models.Scratch, outputPNGPath string) error {
	screenshot := sim.GetScreenshot()
	return easygif.SaveImageToPNG(screenshot, outputPNGPath)
}

func TakeScreenshotVideo(
	sim models.Scratch,
	delayBetweenScreenshots time.Duration,
	frameCount int,
) []image.Image {
	// Collect the images
	frames := make([]image.Image, 0, frameCount)
	nextTime := time.Now()
	for frameIndex := 0; frameIndex < frameCount; frameIndex++ {
		screenShot := sim.GetScreenshot()
		frames = append(frames, screenShot)

		nextTime = nextTime.Add(delayBetweenScreenshots)
		time.Sleep(time.Until(nextTime))
	}

	return frames
}

// Start this as a go routine to create a GIF of your creation.
func CreateGif(
	sim models.Scratch,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	frames := TakeScreenshotVideo(sim, delayBetweenScreenshots, frameCount)

	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	err := easygif.EasyGifWrite(frames, delayBetweenGifFrames, outputGifFilePath)
	if err != nil {
		log.Printf("Error while running easygif.EasyGifWrite(): %v\n", err)
	}

	fmt.Printf("GIF: %s: Done\n", outputGifFilePath)
}

// Start this as a go routine to create a GIF of your creation.
func CreateGifDithered(
	sim models.Scratch,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	frames := TakeScreenshotVideo(sim, delayBetweenScreenshots, frameCount)

	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	err := easygif.EasyDitheredGifWrite(frames, delayBetweenGifFrames, outputGifFilePath)
	if err != nil {
		log.Printf("Error while running easygif.EasyGifWrite(): %v\n", err)
	}

	fmt.Printf("GIF: %s: Done\n", outputGifFilePath)
}

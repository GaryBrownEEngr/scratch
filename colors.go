package scratch

import (
	"image/color"
)

// https://www.w3schools.com/colors/colors_names.asp
var (
	Black color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	White color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Red   color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	Lime  color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	Blue  color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}

	Yellow  color.RGBA = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
	Aqua    color.RGBA = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
	Magenta color.RGBA = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}

	Orange color.RGBA = color.RGBA{0xFF, 0xA5, 0x00, 0xFF}
	Green  color.RGBA = color.RGBA{0x00, 0x80, 0x00, 0xFF}
	Purple color.RGBA = color.RGBA{0x80, 0x00, 0x80, 0xFF}
	Indigo color.RGBA = color.RGBA{0x4B, 0x00, 0x82, 0xFF}
	Violet color.RGBA = color.RGBA{0xEE, 0x82, 0xEE, 0xFF}

	// Other random ones
	SkyBlue color.RGBA = color.RGBA{0x87, 0xCE, 0xEB, 0xFF}
)

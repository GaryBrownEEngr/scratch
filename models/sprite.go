package models

///////////////////////

type SpriteState struct {
	SpriteID       int
	SpriteType     int
	UniqueName     string
	CostumeName    string
	X, Y           float64
	Z              int // Effectively the layer index. 0 through 9 with 9 being the top.
	AngleDegrees   float64
	Visible        bool
	ScaleX, ScaleY float64
	Opacity        float64
	Deleted        bool
}

type Sprite interface {
	GetSpriteID() int
	GetUniqueName() string
	Clone(UniqueName string) Sprite

	// Updates
	Costume(name string)
	SetType(newType int)
	Angle(angleDegrees float64)
	Pos(cartX, cartY float64) // Cartesian (x,y). Center in the middle of the window
	Z(int)                    //
	Visible(visible bool)
	Scale(scale float64) // Sets xScale and yScale together
	XYScale(xScale, yScale float64)
	Opacity(opacityPercent float64) // 0 is completely transparent and 100 is completely opaque
	All(in SpriteState)

	// Info
	GetState() SpriteState

	// Click Body
	GetClickBody() ClickOnBody
	ReplaceClickBody(in ClickOnBody)

	// User Input
	PressedUserInput() *UserInput
	JustPressedUserInput() *UserInput

	// Interact With other sprites
	WhoIsNearMe(distance float64) []NearMeInfo
	SendMsg(toSpriteID int, msg any)
	GetMsgs() []any
	AddMsg(msg any)

	// exit
	DeleteSprite()
}

// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	models "github.com/GaryBrownEEngr/scratch/models"
	mock "github.com/stretchr/testify/mock"
)

// Sprite is an autogenerated mock type for the Sprite type
type Sprite struct {
	mock.Mock
}

// AddMsg provides a mock function with given fields: msg
func (_m *Sprite) AddMsg(msg interface{}) {
	_m.Called(msg)
}

// All provides a mock function with given fields: in
func (_m *Sprite) All(in models.SpriteState) {
	_m.Called(in)
}

// Angle provides a mock function with given fields: angleDegrees
func (_m *Sprite) Angle(angleDegrees float64) {
	_m.Called(angleDegrees)
}

// Clone provides a mock function with given fields: UniqueName
func (_m *Sprite) Clone(UniqueName string) models.Sprite {
	ret := _m.Called(UniqueName)

	var r0 models.Sprite
	if rf, ok := ret.Get(0).(func(string) models.Sprite); ok {
		r0 = rf(UniqueName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Sprite)
		}
	}

	return r0
}

// Costume provides a mock function with given fields: name
func (_m *Sprite) Costume(name string) {
	_m.Called(name)
}

// DeleteSprite provides a mock function with given fields:
func (_m *Sprite) DeleteSprite() {
	_m.Called()
}

// GetClickBody provides a mock function with given fields:
func (_m *Sprite) GetClickBody() models.ClickOnBody {
	ret := _m.Called()

	var r0 models.ClickOnBody
	if rf, ok := ret.Get(0).(func() models.ClickOnBody); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ClickOnBody)
		}
	}

	return r0
}

// GetMsgs provides a mock function with given fields:
func (_m *Sprite) GetMsgs() []interface{} {
	ret := _m.Called()

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func() []interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// GetSpriteID provides a mock function with given fields:
func (_m *Sprite) GetSpriteID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetState provides a mock function with given fields:
func (_m *Sprite) GetState() models.SpriteState {
	ret := _m.Called()

	var r0 models.SpriteState
	if rf, ok := ret.Get(0).(func() models.SpriteState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.SpriteState)
	}

	return r0
}

// GetUniqueName provides a mock function with given fields:
func (_m *Sprite) GetUniqueName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// JustPressedUserInput provides a mock function with given fields:
func (_m *Sprite) JustPressedUserInput() *models.UserInput {
	ret := _m.Called()

	var r0 *models.UserInput
	if rf, ok := ret.Get(0).(func() *models.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInput)
		}
	}

	return r0
}

// Opacity provides a mock function with given fields: opacityPercent
func (_m *Sprite) Opacity(opacityPercent float64) {
	_m.Called(opacityPercent)
}

// Pos provides a mock function with given fields: cartX, cartY
func (_m *Sprite) Pos(cartX float64, cartY float64) {
	_m.Called(cartX, cartY)
}

// PressedUserInput provides a mock function with given fields:
func (_m *Sprite) PressedUserInput() *models.UserInput {
	ret := _m.Called()

	var r0 *models.UserInput
	if rf, ok := ret.Get(0).(func() *models.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInput)
		}
	}

	return r0
}

// ReplaceClickBody provides a mock function with given fields: in
func (_m *Sprite) ReplaceClickBody(in models.ClickOnBody) {
	_m.Called(in)
}

// Scale provides a mock function with given fields: scale
func (_m *Sprite) Scale(scale float64) {
	_m.Called(scale)
}

// SendMsg provides a mock function with given fields: toSpriteID, msg
func (_m *Sprite) SendMsg(toSpriteID int, msg interface{}) {
	_m.Called(toSpriteID, msg)
}

// SetType provides a mock function with given fields: newType
func (_m *Sprite) SetType(newType int) {
	_m.Called(newType)
}

// Visible provides a mock function with given fields: visible
func (_m *Sprite) Visible(visible bool) {
	_m.Called(visible)
}

// WhoIsNearMe provides a mock function with given fields: distance
func (_m *Sprite) WhoIsNearMe(distance float64) []models.NearMeInfo {
	ret := _m.Called(distance)

	var r0 []models.NearMeInfo
	if rf, ok := ret.Get(0).(func(float64) []models.NearMeInfo); ok {
		r0 = rf(distance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.NearMeInfo)
		}
	}

	return r0
}

// XYScale provides a mock function with given fields: xScale, yScale
func (_m *Sprite) XYScale(xScale float64, yScale float64) {
	_m.Called(xScale, yScale)
}

// Z provides a mock function with given fields: _a0
func (_m *Sprite) Z(_a0 int) {
	_m.Called(_a0)
}

// NewSprite creates a new instance of Sprite. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSprite(t interface {
	mock.TestingT
	Cleanup(func())
}) *Sprite {
	mock := &Sprite{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

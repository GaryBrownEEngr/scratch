// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	models "github.com/GaryBrownEEngr/scratch/models"
	mock "github.com/stretchr/testify/mock"
)

// ClickOnBody is an autogenerated mock type for the ClickOnBody type
type ClickOnBody struct {
	mock.Mock
}

// AddCircleBody provides a mock function with given fields: x, y, radius
func (_m *ClickOnBody) AddCircleBody(x float64, y float64, radius float64) {
	_m.Called(x, y, radius)
}

// AddRectangleBody provides a mock function with given fields: x1, x2, y1, y2
func (_m *ClickOnBody) AddRectangleBody(x1 float64, x2 float64, y1 float64, y2 float64) {
	_m.Called(x1, x2, y1, y2)
}

// Angle provides a mock function with given fields: RadAngle
func (_m *ClickOnBody) Angle(RadAngle float64) {
	_m.Called(RadAngle)
}

// Clone provides a mock function with given fields:
func (_m *ClickOnBody) Clone() models.ClickOnBody {
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

// GetMousePosRelativeToOriginalSprite provides a mock function with given fields: x, y
func (_m *ClickOnBody) GetMousePosRelativeToOriginalSprite(x float64, y float64) (float64, float64) {
	ret := _m.Called(x, y)

	var r0 float64
	var r1 float64
	if rf, ok := ret.Get(0).(func(float64, float64) (float64, float64)); ok {
		return rf(x, y)
	}
	if rf, ok := ret.Get(0).(func(float64, float64) float64); ok {
		r0 = rf(x, y)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(float64, float64) float64); ok {
		r1 = rf(x, y)
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// IsMouseClickInBody provides a mock function with given fields: x, y
func (_m *ClickOnBody) IsMouseClickInBody(x float64, y float64) bool {
	ret := _m.Called(x, y)

	var r0 bool
	if rf, ok := ret.Get(0).(func(float64, float64) bool); ok {
		r0 = rf(x, y)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Pos provides a mock function with given fields: x, y
func (_m *ClickOnBody) Pos(x float64, y float64) {
	_m.Called(x, y)
}

// NewClickOnBody creates a new instance of ClickOnBody. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClickOnBody(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClickOnBody {
	mock := &ClickOnBody{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

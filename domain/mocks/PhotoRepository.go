// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mygram-byferdiansyah/domain"

	mock "github.com/stretchr/testify/mock"
)

// ImageRepository is an autogenerated mock type for the ImageRepository type
type ImageRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ImageRepository) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ImageRepository) Get(_a0 context.Context, _a1 *[]domain.Image) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]domain.Image) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *ImageRepository) GetByID(_a0 context.Context, _a1 *domain.Image, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Image, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ImageRepository) Create(_a0 context.Context, _a1 *domain.Image) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Image) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: _a0, _a1, _a2
func (_m *ImageRepository) Edit(_a0 context.Context, _a1 domain.Image, _a2 string) (domain.Image, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 domain.Image
	if rf, ok := ret.Get(0).(func(context.Context, domain.Image, string) domain.Image); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(domain.Image)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Image, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewImageRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageRepository creates a new instance of ImageRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageRepository(t mockConstructorTestingTNewImageRepository) *ImageRepository {
	mock := &ImageRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

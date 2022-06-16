// Code generated by mockery v2.12.3. DO NOT EDIT.

package metadata

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/project/metadata/models"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, projectID, meta
func (_m *Manager) Add(ctx context.Context, projectID int64, meta map[string]string) error {
	ret := _m.Called(ctx, projectID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, map[string]string) error); ok {
		r0 = rf(ctx, projectID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, projectID, meta
func (_m *Manager) Delete(ctx context.Context, projectID int64, meta ...string) error {
	_va := make([]interface{}, len(meta))
	for _i := range meta {
		_va[_i] = meta[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...string) error); ok {
		r0 = rf(ctx, projectID, meta...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, projectID, meta
func (_m *Manager) Get(ctx context.Context, projectID int64, meta ...string) (map[string]string, error) {
	_va := make([]interface{}, len(meta))
	for _i := range meta {
		_va[_i] = meta[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...string) map[string]string); ok {
		r0 = rf(ctx, projectID, meta...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, ...string) error); ok {
		r1 = rf(ctx, projectID, meta...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, name, value
func (_m *Manager) List(ctx context.Context, name string, value string) ([]*models.ProjectMetadata, error) {
	ret := _m.Called(ctx, name, value)

	var r0 []*models.ProjectMetadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*models.ProjectMetadata); ok {
		r0 = rf(ctx, name, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ProjectMetadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, projectID, meta
func (_m *Manager) Update(ctx context.Context, projectID int64, meta map[string]string) error {
	ret := _m.Called(ctx, projectID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, map[string]string) error); ok {
		r0 = rf(ctx, projectID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewManagerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t NewManagerT) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

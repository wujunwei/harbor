// Code generated by mockery v2.12.3. DO NOT EDIT.

package scan

import (
	context "context"

	artifact "github.com/goharbor/harbor/src/controller/artifact"

	daoscan "github.com/goharbor/harbor/src/pkg/scan/dao/scan"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/allowlist/models"

	scan "github.com/goharbor/harbor/src/controller/scan"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeleteReports provides a mock function with given fields: ctx, digests
func (_m *Controller) DeleteReports(ctx context.Context, digests ...string) error {
	_va := make([]interface{}, len(digests))
	for _i := range digests {
		_va[_i] = digests[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, digests...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReport provides a mock function with given fields: ctx, _a1, mimeTypes
func (_m *Controller) GetReport(ctx context.Context, _a1 *artifact.Artifact, mimeTypes []string) ([]*daoscan.Report, error) {
	ret := _m.Called(ctx, _a1, mimeTypes)

	var r0 []*daoscan.Report
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, []string) []*daoscan.Report); ok {
		r0 = rf(ctx, _a1, mimeTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*daoscan.Report)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, []string) error); ok {
		r1 = rf(ctx, _a1, mimeTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanLog provides a mock function with given fields: ctx, uuid
func (_m *Controller) GetScanLog(ctx context.Context, uuid string) ([]byte, error) {
	ret := _m.Called(ctx, uuid)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSummary provides a mock function with given fields: ctx, _a1, mimeTypes
func (_m *Controller) GetSummary(ctx context.Context, _a1 *artifact.Artifact, mimeTypes []string) (map[string]interface{}, error) {
	ret := _m.Called(ctx, _a1, mimeTypes)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, []string) map[string]interface{}); ok {
		r0 = rf(ctx, _a1, mimeTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, []string) error); ok {
		r1 = rf(ctx, _a1, mimeTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVulnerable provides a mock function with given fields: ctx, _a1, allowlist
func (_m *Controller) GetVulnerable(ctx context.Context, _a1 *artifact.Artifact, allowlist models.CVESet) (*scan.Vulnerable, error) {
	ret := _m.Called(ctx, _a1, allowlist)

	var r0 *scan.Vulnerable
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, models.CVESet) *scan.Vulnerable); ok {
		r0 = rf(ctx, _a1, allowlist)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.Vulnerable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, models.CVESet) error); ok {
		r1 = rf(ctx, _a1, allowlist)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scan provides a mock function with given fields: ctx, _a1, options
func (_m *Controller) Scan(ctx context.Context, _a1 *artifact.Artifact, options ...scan.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, ...scan.Option) error); ok {
		r0 = rf(ctx, _a1, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScanAll provides a mock function with given fields: ctx, trigger, async
func (_m *Controller) ScanAll(ctx context.Context, trigger string, async bool) (int64, error) {
	ret := _m.Called(ctx, trigger, async)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) int64); ok {
		r0 = rf(ctx, trigger, async)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, trigger, async)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: ctx, _a1
func (_m *Controller) Stop(ctx context.Context, _a1 *artifact.Artifact) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewControllerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t NewControllerT) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

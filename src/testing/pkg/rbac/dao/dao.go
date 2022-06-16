// Code generated by mockery v2.12.3. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/goharbor/harbor/src/pkg/rbac/model"

	q "github.com/goharbor/harbor/src/lib/q"
)

// DAO is an autogenerated mock type for the DAO type
type DAO struct {
	mock.Mock
}

// CreatePermission provides a mock function with given fields: ctx, rp
func (_m *DAO) CreatePermission(ctx context.Context, rp *model.RolePermission) (int64, error) {
	ret := _m.Called(ctx, rp)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.RolePermission) int64); ok {
		r0 = rf(ctx, rp)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.RolePermission) error); ok {
		r1 = rf(ctx, rp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRbacPolicy provides a mock function with given fields: ctx, pp
func (_m *DAO) CreateRbacPolicy(ctx context.Context, pp *model.PermissionPolicy) (int64, error) {
	ret := _m.Called(ctx, pp)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.PermissionPolicy) int64); ok {
		r0 = rf(ctx, pp)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.PermissionPolicy) error); ok {
		r1 = rf(ctx, pp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePermission provides a mock function with given fields: ctx, id
func (_m *DAO) DeletePermission(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePermissionsByRole provides a mock function with given fields: ctx, roleType, roleID
func (_m *DAO) DeletePermissionsByRole(ctx context.Context, roleType string, roleID int64) error {
	ret := _m.Called(ctx, roleType, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, roleType, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRbacPolicy provides a mock function with given fields: ctx, id
func (_m *DAO) DeleteRbacPolicy(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPermissionsByRole provides a mock function with given fields: ctx, roleType, roleID
func (_m *DAO) GetPermissionsByRole(ctx context.Context, roleType string, roleID int64) ([]*model.UniversalRolePermission, error) {
	ret := _m.Called(ctx, roleType, roleID)

	var r0 []*model.UniversalRolePermission
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []*model.UniversalRolePermission); ok {
		r0 = rf(ctx, roleType, roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UniversalRolePermission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, roleType, roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPermissions provides a mock function with given fields: ctx, query
func (_m *DAO) ListPermissions(ctx context.Context, query *q.Query) ([]*model.RolePermission, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.RolePermission); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RolePermission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRbacPolicies provides a mock function with given fields: ctx, query
func (_m *DAO) ListRbacPolicies(ctx context.Context, query *q.Query) ([]*model.PermissionPolicy, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.PermissionPolicy
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.PermissionPolicy); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PermissionPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewDAOT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDAO creates a new instance of DAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDAO(t NewDAOT) *DAO {
	mock := &DAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

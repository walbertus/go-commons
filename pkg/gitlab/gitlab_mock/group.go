// Code generated by mockery 2.9.0. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// Group is an autogenerated mock type for the Group type
type Group struct {
	mock.Mock
}

// ListAllProjects provides a mock function with given fields: id
func (_m *Group) ListAllProjects(id gitlab.NameOrId) (<-chan go_gitlab.Project, error) {
	ret := _m.Called(id)

	var r0 <-chan go_gitlab.Project
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId) <-chan go_gitlab.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan go_gitlab.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

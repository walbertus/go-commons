// Code generated by mockery 2.9.0. DO NOT EDIT.

package tmpl_mock

import (
	template "text/template"

	mock "github.com/stretchr/testify/mock"

	tmpl "github.com/gopaytech/go-commons/pkg/tmpl"
)

// ScanResult is an autogenerated mock type for the ScanResult type
type ScanResult struct {
	mock.Mock
}

// DirList provides a mock function with given fields:
func (_m *ScanResult) DirList() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Execute provides a mock function with given fields: data
func (_m *ScanResult) Execute(data interface{}) (map[string]string, error) {
	ret := _m.Called(data)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(interface{}) map[string]string); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteToPath provides a mock function with given fields: data, targetPath
func (_m *ScanResult) ExecuteToPath(data interface{}, targetPath string) error {
	ret := _m.Called(data, targetPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(data, targetPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Option provides a mock function with given fields:
func (_m *ScanResult) Option() *tmpl.ScanOption {
	ret := _m.Called()

	var r0 *tmpl.ScanOption
	if rf, ok := ret.Get(0).(func() *tmpl.ScanOption); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tmpl.ScanOption)
		}
	}

	return r0
}

// RootPath provides a mock function with given fields:
func (_m *ScanResult) RootPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Template provides a mock function with given fields:
func (_m *ScanResult) Template() *template.Template {
	ret := _m.Called()

	var r0 *template.Template
	if rf, ok := ret.Get(0).(func() *template.Template); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*template.Template)
		}
	}

	return r0
}

// TemplateList provides a mock function with given fields:
func (_m *ScanResult) TemplateList() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// TemplateMap provides a mock function with given fields:
func (_m *ScanResult) TemplateMap() tmpl.FileMap {
	ret := _m.Called()

	var r0 tmpl.FileMap
	if rf, ok := ret.Get(0).(func() tmpl.FileMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tmpl.FileMap)
		}
	}

	return r0
}

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	http "net/http"

	domain "github.com/sreeks87/webpageinfo/pageinfo/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Extract provides a mock function with given fields:
func (_m *Service) Extract() (domain.Pageinfo, error) {
	ret := _m.Called()

	var r0 domain.Pageinfo
	if rf, ok := ret.Get(0).(func() domain.Pageinfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Pageinfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scrape provides a mock function with given fields:
func (_m *Service) Scrape() (*http.Response, error) {
	ret := _m.Called()

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func() *http.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields:
func (_m *Service) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

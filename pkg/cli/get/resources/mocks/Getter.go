// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/opencontrol/compliance-masonry/pkg/lib/common"
import constants "github.com/opencontrol/compliance-masonry/tools/constants"
import mock "github.com/stretchr/testify/mock"

// Getter is an autogenerated mock type for the Getter type
type Getter struct {
	mock.Mock
}

// GetLocalResources provides a mock function with given fields: source, _a1, destination, subfolder, recursively, resourceType
func (_m *Getter) GetLocalResources(source string, _a1 []string, destination string, subfolder string, recursively bool, resourceType constants.ResourceType) error {
	ret := _m.Called(source, _a1, destination, subfolder, recursively, resourceType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, string, string, bool, constants.ResourceType) error); ok {
		r0 = rf(source, _a1, destination, subfolder, recursively, resourceType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRemoteResources provides a mock function with given fields: destination, subfolder, entries
func (_m *Getter) GetRemoteResources(destination string, subfolder string, entries []common.RemoteSource) error {
	ret := _m.Called(destination, subfolder, entries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []common.RemoteSource) error); ok {
		r0 = rf(destination, subfolder, entries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

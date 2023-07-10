// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	logger "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	pkgmodels "github.com/gq-tang/device-sdk-go/v2/pkg/models"
)

// ProtocolDriver is an autogenerated mock type for the ProtocolDriver type
type ProtocolDriver struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: deviceName, protocols, adminState
func (_m *ProtocolDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	ret := _m.Called(deviceName, protocols, adminState)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, models.AdminState) error); ok {
		r0 = rf(deviceName, protocols, adminState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleReadCommands provides a mock function with given fields: deviceName, protocols, reqs
func (_m *ProtocolDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest) ([]*pkgmodels.CommandValue, error) {
	ret := _m.Called(deviceName, protocols, reqs)

	var r0 []*pkgmodels.CommandValue
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) []*pkgmodels.CommandValue); ok {
		r0 = rf(deviceName, protocols, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkgmodels.CommandValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) error); ok {
		r1 = rf(deviceName, protocols, reqs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleWriteCommands provides a mock function with given fields: deviceName, protocols, reqs, params
func (_m *ProtocolDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest, params []*pkgmodels.CommandValue) error {
	ret := _m.Called(deviceName, protocols, reqs, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest, []*pkgmodels.CommandValue) error); ok {
		r0 = rf(deviceName, protocols, reqs, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: lc, asyncCh, deviceCh
func (_m *ProtocolDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *pkgmodels.AsyncValues, deviceCh chan<- []pkgmodels.DiscoveredDevice) error {
	ret := _m.Called(lc, asyncCh, deviceCh)

	var r0 error
	if rf, ok := ret.Get(0).(func(logger.LoggingClient, chan<- *pkgmodels.AsyncValues, chan<- []pkgmodels.DiscoveredDevice) error); ok {
		r0 = rf(lc, asyncCh, deviceCh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDevice provides a mock function with given fields: deviceName, protocols
func (_m *ProtocolDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	ret := _m.Called(deviceName, protocols)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties) error); ok {
		r0 = rf(deviceName, protocols)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: force
func (_m *ProtocolDriver) Stop(force bool) error {
	ret := _m.Called(force)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: deviceName, protocols, adminState
func (_m *ProtocolDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	ret := _m.Called(deviceName, protocols, adminState)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, models.AdminState) error); ok {
		r0 = rf(deviceName, protocols, adminState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

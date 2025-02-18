/*
Copyright 2024 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: map.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ebpf "github.com/cilium/ebpf"
	gomock "github.com/golang/mock/gomock"
)

// MockEbpfMap is a mock of EbpfMap interface.
type MockEbpfMap struct {
	ctrl     *gomock.Controller
	recorder *MockEbpfMapMockRecorder
}

// MockEbpfMapMockRecorder is the mock recorder for MockEbpfMap.
type MockEbpfMapMockRecorder struct {
	mock *MockEbpfMap
}

// NewMockEbpfMap creates a new mock instance.
func NewMockEbpfMap(ctrl *gomock.Controller) *MockEbpfMap {
	mock := &MockEbpfMap{ctrl: ctrl}
	mock.recorder = &MockEbpfMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEbpfMap) EXPECT() *MockEbpfMapMockRecorder {
	return m.recorder
}

// CreateAndPinArrayMap mocks base method.
func (m *MockEbpfMap) CreateAndPinArrayMap(dir, name string, defaultKV []ebpf.MapKV) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndPinArrayMap", dir, name, defaultKV)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAndPinArrayMap indicates an expected call of CreateAndPinArrayMap.
func (mr *MockEbpfMapMockRecorder) CreateAndPinArrayMap(dir, name, defaultKV interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndPinArrayMap", reflect.TypeOf((*MockEbpfMap)(nil).CreateAndPinArrayMap), dir, name, defaultKV)
}

// LookUpMapValue mocks base method.
func (m *MockEbpfMap) LookUpMapValue(pinPath string, key, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookUpMapValue", pinPath, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// LookUpMapValue indicates an expected call of LookUpMapValue.
func (mr *MockEbpfMapMockRecorder) LookUpMapValue(pinPath, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookUpMapValue", reflect.TypeOf((*MockEbpfMap)(nil).LookUpMapValue), pinPath, key, value)
}

// UnpinArrayMap mocks base method.
func (m *MockEbpfMap) UnpinArrayMap(pinPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpinArrayMap", pinPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpinArrayMap indicates an expected call of UnpinArrayMap.
func (mr *MockEbpfMapMockRecorder) UnpinArrayMap(pinPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinArrayMap", reflect.TypeOf((*MockEbpfMap)(nil).UnpinArrayMap), pinPath)
}

// UpdateMapValue mocks base method.
func (m *MockEbpfMap) UpdateMapValue(pinPath string, key, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMapValue", pinPath, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMapValue indicates an expected call of UpdateMapValue.
func (mr *MockEbpfMapMockRecorder) UpdateMapValue(pinPath, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMapValue", reflect.TypeOf((*MockEbpfMap)(nil).UpdateMapValue), pinPath, key, value)
}

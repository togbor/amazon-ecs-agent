// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/stats/resolver (interfaces: ContainerMetadataResolver)

package mock_resolver

import (
	gomock "code.google.com/p/gomock/gomock"
	api "github.com/aws/amazon-ecs-agent/agent/api"
)

// Mock of ContainerMetadataResolver interface
type MockContainerMetadataResolver struct {
	ctrl     *gomock.Controller
	recorder *_MockContainerMetadataResolverRecorder
}

// Recorder for MockContainerMetadataResolver (not exported)
type _MockContainerMetadataResolverRecorder struct {
	mock *MockContainerMetadataResolver
}

func NewMockContainerMetadataResolver(ctrl *gomock.Controller) *MockContainerMetadataResolver {
	mock := &MockContainerMetadataResolver{ctrl: ctrl}
	mock.recorder = &_MockContainerMetadataResolverRecorder{mock}
	return mock
}

func (_m *MockContainerMetadataResolver) EXPECT() *_MockContainerMetadataResolverRecorder {
	return _m.recorder
}

func (_m *MockContainerMetadataResolver) ResolveName(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "ResolveName", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerMetadataResolverRecorder) ResolveName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResolveName", arg0)
}

func (_m *MockContainerMetadataResolver) ResolveTask(_param0 string) (*api.Task, error) {
	ret := _m.ctrl.Call(_m, "ResolveTask", _param0)
	ret0, _ := ret[0].(*api.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerMetadataResolverRecorder) ResolveTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResolveTask", arg0)
}

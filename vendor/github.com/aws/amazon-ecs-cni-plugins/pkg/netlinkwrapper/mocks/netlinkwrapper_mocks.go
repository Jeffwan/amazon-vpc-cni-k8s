// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: github.com/aws/amazon-ecs-cni-plugins/pkg/netlinkwrapper (interfaces: NetLink)

package mock_netlinkwrapper

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/vishvananda/netlink"
)

// Mock of NetLink interface
type MockNetLink struct {
	ctrl     *gomock.Controller
	recorder *_MockNetLinkRecorder
}

// Recorder for MockNetLink (not exported)
type _MockNetLinkRecorder struct {
	mock *MockNetLink
}

func NewMockNetLink(ctrl *gomock.Controller) *MockNetLink {
	mock := &MockNetLink{ctrl: ctrl}
	mock.recorder = &_MockNetLinkRecorder{mock}
	return mock
}

func (_m *MockNetLink) EXPECT() *_MockNetLinkRecorder {
	return _m.recorder
}

func (_m *MockNetLink) AddrAdd(_param0 netlink.Link, _param1 *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrAdd", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrAdd", arg0, arg1)
}

func (_m *MockNetLink) AddrList(_param0 netlink.Link, _param1 int) ([]netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "AddrList", _param0, _param1)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetLinkRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrList", arg0, arg1)
}

func (_m *MockNetLink) LinkAdd(_param0 netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkAdd", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkAdd", arg0)
}

func (_m *MockNetLink) LinkByName(_param0 string) (netlink.Link, error) {
	ret := _m.ctrl.Call(_m, "LinkByName", _param0)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetLinkRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkByName", arg0)
}

func (_m *MockNetLink) LinkList() ([]netlink.Link, error) {
	ret := _m.ctrl.Call(_m, "LinkList")
	ret0, _ := ret[0].([]netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetLinkRecorder) LinkList() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkList")
}

func (_m *MockNetLink) LinkSetDown(_param0 netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetDown", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) LinkSetDown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetDown", arg0)
}

func (_m *MockNetLink) LinkSetMaster(_param0 netlink.Link, _param1 *netlink.Bridge) error {
	ret := _m.ctrl.Call(_m, "LinkSetMaster", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) LinkSetMaster(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetMaster", arg0, arg1)
}

func (_m *MockNetLink) LinkSetNsFd(_param0 netlink.Link, _param1 int) error {
	ret := _m.ctrl.Call(_m, "LinkSetNsFd", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) LinkSetNsFd(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetNsFd", arg0, arg1)
}

func (_m *MockNetLink) LinkSetUp(_param0 netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetUp", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetUp", arg0)
}

func (_m *MockNetLink) ParseAddr(_param0 string) (*netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "ParseAddr", _param0)
	ret0, _ := ret[0].(*netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetLinkRecorder) ParseAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseAddr", arg0)
}

func (_m *MockNetLink) RouteAdd(_param0 *netlink.Route) error {
	ret := _m.ctrl.Call(_m, "RouteAdd", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) RouteAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteAdd", arg0)
}

func (_m *MockNetLink) RouteDel(_param0 *netlink.Route) error {
	ret := _m.ctrl.Call(_m, "RouteDel", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetLinkRecorder) RouteDel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteDel", arg0)
}

func (_m *MockNetLink) RouteList(_param0 netlink.Link, _param1 int) ([]netlink.Route, error) {
	ret := _m.ctrl.Call(_m, "RouteList", _param0, _param1)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetLinkRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteList", arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/proxy/cobra.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/proxy/cobra.go -destination=./pkg/proxy/cobra_mock.go -package=proxy
//

// Package proxy is a generated GoMock package.
package proxy

import (
	io "io"
	reflect "reflect"

	cobra "github.com/spf13/cobra"
	gomock "go.uber.org/mock/gomock"
)

// MockCobra is a mock of Cobra interface.
type MockCobra struct {
	ctrl     *gomock.Controller
	recorder *MockCobraMockRecorder
}

// MockCobraMockRecorder is the mock recorder for MockCobra.
type MockCobraMockRecorder struct {
	mock *MockCobra
}

// NewMockCobra creates a new mock instance.
func NewMockCobra(ctrl *gomock.Controller) *MockCobra {
	mock := &MockCobra{ctrl: ctrl}
	mock.recorder = &MockCobraMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCobra) EXPECT() *MockCobraMockRecorder {
	return m.recorder
}

// ExactArgs mocks base method.
func (m *MockCobra) ExactArgs(arg0 int) PositionalArgs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExactArgs", arg0)
	ret0, _ := ret[0].(PositionalArgs)
	return ret0
}

// ExactArgs indicates an expected call of ExactArgs.
func (mr *MockCobraMockRecorder) ExactArgs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExactArgs", reflect.TypeOf((*MockCobra)(nil).ExactArgs), arg0)
}

// MaximumNArgs mocks base method.
func (m *MockCobra) MaximumNArgs(arg0 int) PositionalArgs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaximumNArgs", arg0)
	ret0, _ := ret[0].(PositionalArgs)
	return ret0
}

// MaximumNArgs indicates an expected call of MaximumNArgs.
func (mr *MockCobraMockRecorder) MaximumNArgs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaximumNArgs", reflect.TypeOf((*MockCobra)(nil).MaximumNArgs), arg0)
}

// NewCommand mocks base method.
func (m *MockCobra) NewCommand() Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCommand")
	ret0, _ := ret[0].(Command)
	return ret0
}

// NewCommand indicates an expected call of NewCommand.
func (mr *MockCobraMockRecorder) NewCommand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCommand", reflect.TypeOf((*MockCobra)(nil).NewCommand))
}

// MockPositionalArgs is a mock of PositionalArgs interface.
type MockPositionalArgs struct {
	ctrl     *gomock.Controller
	recorder *MockPositionalArgsMockRecorder
}

// MockPositionalArgsMockRecorder is the mock recorder for MockPositionalArgs.
type MockPositionalArgsMockRecorder struct {
	mock *MockPositionalArgs
}

// NewMockPositionalArgs creates a new mock instance.
func NewMockPositionalArgs(ctrl *gomock.Controller) *MockPositionalArgs {
	mock := &MockPositionalArgs{ctrl: ctrl}
	mock.recorder = &MockPositionalArgsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPositionalArgs) EXPECT() *MockPositionalArgsMockRecorder {
	return m.recorder
}

// GetPositionalArgs mocks base method.
func (m *MockPositionalArgs) GetPositionalArgs() cobra.PositionalArgs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionalArgs")
	ret0, _ := ret[0].(cobra.PositionalArgs)
	return ret0
}

// GetPositionalArgs indicates an expected call of GetPositionalArgs.
func (mr *MockPositionalArgsMockRecorder) GetPositionalArgs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionalArgs", reflect.TypeOf((*MockPositionalArgs)(nil).GetPositionalArgs))
}

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// AddCommand mocks base method.
func (m *MockCommand) AddCommand(cmds ...Command) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range cmds {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddCommand", varargs...)
}

// AddCommand indicates an expected call of AddCommand.
func (mr *MockCommandMockRecorder) AddCommand(cmds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCommand", reflect.TypeOf((*MockCommand)(nil).AddCommand), cmds...)
}

// Execute mocks base method.
func (m *MockCommand) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockCommandMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute))
}

// GetCommand mocks base method.
func (m *MockCommand) GetCommand() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommand")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetCommand indicates an expected call of GetCommand.
func (mr *MockCommandMockRecorder) GetCommand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommand", reflect.TypeOf((*MockCommand)(nil).GetCommand))
}

// PersistentFlags mocks base method.
func (m *MockCommand) PersistentFlags() FlagSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistentFlags")
	ret0, _ := ret[0].(FlagSet)
	return ret0
}

// PersistentFlags indicates an expected call of PersistentFlags.
func (mr *MockCommandMockRecorder) PersistentFlags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistentFlags", reflect.TypeOf((*MockCommand)(nil).PersistentFlags))
}

// RunE mocks base method.
func (m *MockCommand) RunE(cmd *cobra.Command, args []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunE", cmd, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunE indicates an expected call of RunE.
func (mr *MockCommandMockRecorder) RunE(cmd, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunE", reflect.TypeOf((*MockCommand)(nil).RunE), cmd, args)
}

// SetAliases mocks base method.
func (m *MockCommand) SetAliases(aliases []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAliases", aliases)
}

// SetAliases indicates an expected call of SetAliases.
func (mr *MockCommandMockRecorder) SetAliases(aliases any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAliases", reflect.TypeOf((*MockCommand)(nil).SetAliases), aliases)
}

// SetArgs mocks base method.
func (m *MockCommand) SetArgs(positionalArgs PositionalArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetArgs", positionalArgs)
}

// SetArgs indicates an expected call of SetArgs.
func (mr *MockCommandMockRecorder) SetArgs(positionalArgs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArgs", reflect.TypeOf((*MockCommand)(nil).SetArgs), positionalArgs)
}

// SetErr mocks base method.
func (m *MockCommand) SetErr(io io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetErr", io)
}

// SetErr indicates an expected call of SetErr.
func (mr *MockCommandMockRecorder) SetErr(io any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetErr", reflect.TypeOf((*MockCommand)(nil).SetErr), io)
}

// SetHelpTemplate mocks base method.
func (m *MockCommand) SetHelpTemplate(s string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHelpTemplate", s)
}

// SetHelpTemplate indicates an expected call of SetHelpTemplate.
func (mr *MockCommandMockRecorder) SetHelpTemplate(s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHelpTemplate", reflect.TypeOf((*MockCommand)(nil).SetHelpTemplate), s)
}

// SetMaximumNArgs mocks base method.
func (m *MockCommand) SetMaximumNArgs(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaximumNArgs", n)
}

// SetMaximumNArgs indicates an expected call of SetMaximumNArgs.
func (mr *MockCommandMockRecorder) SetMaximumNArgs(n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaximumNArgs", reflect.TypeOf((*MockCommand)(nil).SetMaximumNArgs), n)
}

// SetOut mocks base method.
func (m *MockCommand) SetOut(io io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOut", io)
}

// SetOut indicates an expected call of SetOut.
func (mr *MockCommandMockRecorder) SetOut(io any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOut", reflect.TypeOf((*MockCommand)(nil).SetOut), io)
}

// SetRunE mocks base method.
func (m *MockCommand) SetRunE(runE func(*cobra.Command, []string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRunE", runE)
}

// SetRunE indicates an expected call of SetRunE.
func (mr *MockCommandMockRecorder) SetRunE(runE any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRunE", reflect.TypeOf((*MockCommand)(nil).SetRunE), runE)
}

// SetSilenceErrors mocks base method.
func (m *MockCommand) SetSilenceErrors(silenceErrors bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSilenceErrors", silenceErrors)
}

// SetSilenceErrors indicates an expected call of SetSilenceErrors.
func (mr *MockCommandMockRecorder) SetSilenceErrors(silenceErrors any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSilenceErrors", reflect.TypeOf((*MockCommand)(nil).SetSilenceErrors), silenceErrors)
}

// SetUsageTemplate mocks base method.
func (m *MockCommand) SetUsageTemplate(s string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUsageTemplate", s)
}

// SetUsageTemplate indicates an expected call of SetUsageTemplate.
func (mr *MockCommandMockRecorder) SetUsageTemplate(s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsageTemplate", reflect.TypeOf((*MockCommand)(nil).SetUsageTemplate), s)
}

// SetUse mocks base method.
func (m *MockCommand) SetUse(use string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUse", use)
}

// SetUse indicates an expected call of SetUse.
func (mr *MockCommandMockRecorder) SetUse(use any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUse", reflect.TypeOf((*MockCommand)(nil).SetUse), use)
}

// SetVersion mocks base method.
func (m *MockCommand) SetVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVersion", version)
}

// SetVersion indicates an expected call of SetVersion.
func (mr *MockCommandMockRecorder) SetVersion(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersion", reflect.TypeOf((*MockCommand)(nil).SetVersion), version)
}

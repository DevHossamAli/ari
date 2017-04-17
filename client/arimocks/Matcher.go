package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Matcher is an autogenerated mock type for the Matcher type
type Matcher struct {
	mock.Mock
}

// Match provides a mock function with given fields: evt
func (_m *Matcher) Match(evt ari.Event) bool {
	ret := _m.Called(evt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ari.Event) bool); ok {
		r0 = rf(evt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
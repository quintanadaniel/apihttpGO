package mock

import (
	beercli "github.com/CodelyTV/test_project/apihttp/cli"
	"github.com/stretchr/testify/mock"
)

//MyMock structure definited mock beers
type MyMock struct {
	mock.Mock
}

//GetBeersAll function mock
func (m *MyMock) GetBeersAll() []beercli.Beer {
	args := m.Called()
	return args.Get(0).([]beercli.Beer)
}

//GetBeersID function mock
func (m *MyMock) GetBeersID(id string) []beercli.Beer {
	args := m.Called(id)
	return args.Get(0).([]beercli.Beer)
}

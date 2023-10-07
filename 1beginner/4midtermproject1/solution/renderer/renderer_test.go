package renderer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockNavigator struct {
	mock.Mock
}

func (mn *MockNavigator) show() error {
	args := mn.Called()
	return args.Error(0)
}

func (mn *MockNavigator) navigate() error {
	args := mn.Called()
	return args.Error(0)
}

func TestRenderer_Render(t *testing.T) {
	mockNavigator := new(MockNavigator)
	R.n = mockNavigator

	mockNavigator.On("show", mock.Anything).Return(errors.New("mock error on shown()"))
	mockNavigator.On("navigate", mock.Anything).Return(nil)
	err := R.Render()
	assert.NotNil(t, err)

	mockNavigator.On("show", mock.Anything).Return(nil)
	mockNavigator.On("navigate", mock.Anything).Return(errors.New("mock error on navigate()"))
	err = R.Render()
	assert.NotNil(t, err)
}
package gorpc_demo

import (
	"github.com/stretchr/testify/assert"
	"gorpc-demo/testdata"
	"testing"
)

func TestRegisterService(t *testing.T) {
	s := NewServer()
	err := s.RegisterService("helloworld", new(testdata.Service))
	assert.Nil(t, err)
}
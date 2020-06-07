package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseServicePath(t *testing.T) {
	_, _, err := ParseServicePath("Greeter.Hello")
	assert.NotNil(t, err)

	_, _, err = ParseServicePath("Greeter/hello")
	assert.NotNil(t, err)

	serviceName, method, err := ParseServicePath("/Greeter/Hello")
	assert.Equal(t, serviceName, "Greeter")
	assert.Equal(t, method, "Hello")
	assert.Equal(t, err, nil)
}
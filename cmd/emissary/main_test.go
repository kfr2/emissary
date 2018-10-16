package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	assert.Equal(t, 3, actual)

	actual = Add(-1, -2)
	assert.Equal(t, -3, actual)

	actual = Add(0, 0)
	assert.Equal(t, 0, actual)
}

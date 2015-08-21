package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessLine(t *testing.T) {
	expected := ProcessLine("test", "prefix")
	assert.Equal(t, expected, "prefix.test")
}

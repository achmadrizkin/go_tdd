package test

import (
	"bytes"
	logic "go_tdd/logic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDownTrue(t *testing.T) {
	buffer := &bytes.Buffer{}

	logic.Countdown(buffer)

	got := buffer.String()
	want := "3"

	assert.Equal(t, want, got)
}

func TestCountdownFalse(t *testing.T) {
	buffer := &bytes.Buffer{}

	logic.Countdown(buffer)

	got := buffer.String()
	want := "5"

	assert.Equal(t, want, got)
}

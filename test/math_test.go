package test

import (
	"testing"
	"time"

	logic "go_tdd/logic"

	"github.com/stretchr/testify/assert"
)

func TestOddOrEvenFAIL(t *testing.T) {
	assert.NotEqual(t, "42 is an odd number", logic.OddOrEven(45))
	assert.NotEqual(t, "46 is an odd number", logic.OddOrEven(47))
	assert.NotEqual(t, "3 is an odd number", logic.OddOrEven(4))
	assert.NotEqual(t, "4 is an odd number", logic.OddOrEven(4))
}

func TestOddOrEvenSUCCESS(t *testing.T) {
	assert.Equal(t, "45 is an odd number", logic.OddOrEven(45))
	assert.Equal(t, "47 is an odd number", logic.OddOrEven(47))
	assert.Equal(t, "123 is an odd number", logic.OddOrEven(123))
	assert.Equal(t, "3 is an odd number", logic.OddOrEven(3))
}

func TestSecondHandAtMidnightSUCCESS(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := logic.Point{X: 150, Y: 150 - 90}
	got := logic.SecondHand(tm)

	assert.Equal(t, got, want)
}

func TestSecondHandAtMidnightFAIL(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := logic.Point{X: 150, Y: 150 - 40}
	got := logic.SecondHand(tm)

	assert.NotEmpty(t, got, want)
}

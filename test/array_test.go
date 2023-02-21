package test

import (
	"go_tdd/logic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSUCCESS(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := logic.Sum(numbers)
	want := 15

	assert.Equal(t, got, want)
}

func TestSumFail(t *testing.T) {
	numbers := []int{1, 56, 5, 1, 1}

	got := logic.Sum(numbers)
	want := 61 // 62 true

	assert.NotEqual(t, got, want)
}

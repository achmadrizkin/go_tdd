package test

import (
	"go_tdd/logic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeterSUCCESS(t *testing.T) {
	got := logic.Perimeter(10.0, 10.0)
	want := 40.0

	assert.Equal(t, want, got)
}

func TestPerimeterFAIL(t *testing.T) {
	got := logic.Perimeter(20.0, 10.0)
	want := 40.0

	assert.NotEqual(t, want, got)
}

func TestAreaSUCCESS(t *testing.T) {
	t.Run("circles", func(t *testing.T) {
		circle := logic.Circle{Radius: 10}
		got := circle.Area()
		want := 314.1592653589793

		assert.Equal(t, want, got)
	})
}

func TestAreaFAIL(t *testing.T) {
	t.Run("circles", func(t *testing.T) {
		circle := logic.Circle{Radius: 10}
		got := circle.Area()
		want := 334.1592653589793

		assert.NotEqual(t, want, got)
	})
}

func TestAreaAndRectangleSUCCESS(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := logic.Rectangle{Width: 12, Height: 6}
		got := rectangle.Area()
		want := 72.0

		assert.Equal(t, want, got)
	})

	t.Run("circles", func(t *testing.T) {
		circle := logic.Circle{Radius: 10}
		got := circle.Area()
		want := 314.1592653589793

		assert.Equal(t, want, got)
	})
}

func TestAreaAndRectangleByInterfaceSUCCESS(t *testing.T) {
	checkArea := func(t testing.TB, shape logic.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.Equal(t, want, got)
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := logic.Rectangle{Width: 12, Height: 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := logic.Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
}

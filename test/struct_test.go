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

func TestAreaAndRectangleByInterfaceFAIL(t *testing.T) {
	checkArea := func(t testing.TB, shape logic.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.NotEqual(t, want, got)
	}

	t.Run("circles", func(t *testing.T) {
		circle := logic.Circle{Radius: 15}
		checkArea(t, circle, 123.1592653589793)
	})

	t.Run("rectangle", func(t *testing.T) {
		rectangle := logic.Rectangle{Width: 12, Height: 6}
		checkArea(t, rectangle, 22.0)
	})
}

func TestAreaAndRectangleBatchSUCCESS(t *testing.T) {
	areaTest := []struct {
		shape logic.Shape
		want  float64
	}{
		{logic.Rectangle{Width: 12, Height: 6}, 72.0},
		{logic.Circle{Radius: 10}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()

		assert.NotEmpty(t, areaTest)
		assert.Equal(t, tt.want, got)
	}
}

func TestAreaAndRectangleBatchFAIL(t *testing.T) {
	areaTest := []struct {
		shape logic.Shape
		want  float64
	}{
		{logic.Rectangle{Width: 12, Height: 6}, 12.0},
		{logic.Circle{Radius: 1}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		assert.NotEqual(t, tt.want, got)
	}
}

func TestAreaRectangleAndTriangleBatchSUCCESS(t *testing.T) {
	areaTest := []struct {
		name  string
		shape logic.Shape
		want  float64
	}{
		{name: "Rectangle", shape: logic.Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: logic.Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: logic.Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAreaRectangleAndTriangleBatchFAIL(t *testing.T) {
	areaTest := []struct {
		name  string
		shape logic.Shape
		want  float64
	}{
		{name: "Rectangle", shape: logic.Rectangle{Width: 35, Height: 6}, want: 72.0},
		{name: "Circle", shape: logic.Circle{Radius: 750}, want: 314.1592653589793},
		{name: "Triangle", shape: logic.Triangle{Base: 132, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Logf("%#v got %g want %g", tt.shape, got, tt.want)
			}

			assert.NotEqual(t, tt.want, got)
		})
	}
}

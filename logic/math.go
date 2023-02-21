package logic

import (
	"fmt"
	"math"
	"time"
)

func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

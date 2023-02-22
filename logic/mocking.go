package logic

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func CountdownLoop(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

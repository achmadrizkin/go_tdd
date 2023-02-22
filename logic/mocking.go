package logic

import (
	"fmt"
	"io"
	"os"
	"time"
)

// sleeper to put delay
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

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

// this countdown will delay in certain of time
func CountdownSleeper(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep() // this delay 3 seconds
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	CountdownSleeper(os.Stdout, sleeper)
}

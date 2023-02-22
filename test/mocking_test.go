package test

import (
	"bytes"
	logic "go_tdd/logic"
	"reflect"
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

	assert.NotEqual(t, want, got)
}

func TestCountdownLoopSUCCESS(t *testing.T) {
	buffer := &bytes.Buffer{} // bit val
	logic.CountdownLoop(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	assert.Equal(t, want, got)
}

func TestCountdownLoopFAILBecauseSPACE(t *testing.T) {
	buffer := &bytes.Buffer{} // bit val
	logic.CountdownLoop(buffer)

	// this program will error, becase there have a space (space have value in bit)
	got := buffer.String()
	want := `3
 2
 1
 Go!`

	assert.NotEqual(t, want, got)
}

func TestCountdownLoopFAILBecauseWrongValue(t *testing.T) {
	buffer := &bytes.Buffer{} // bit val
	logic.CountdownLoop(buffer)

	// this program will error, becase there have a diff value
	got := buffer.String()
	want := `2
 1
 3
 gO!`

	assert.NotEqual(t, want, got)
}

func TestCountdownSleepSUCCESS(t *testing.T) {
	buffer := &bytes.Buffer{} // bit val
	spySleeper := &SpySleeper{}

	logic.CountdownSleeper(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	// this delay come from mocking.go
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, 3, spySleeper.Calls)
}

// test sleep between print (success)
func TestCountdownBetweenPrintSUCCESS(t *testing.T) {
	t.Run("sleep in-between the prints", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		logic.CountdownSleeper(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

		assert.Equal(t, want, spySleepPrinter.Calls)
	})
}

func TestCountdownBetweenPrintFAIL(t *testing.T) {
	t.Run("sleep in-between the prints", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		logic.CountdownSleeper(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.NotEqual(t, want, spySleepPrinter.FailedCalls) // got: sleep, write, sleep, etc
	})
}

func TestCountdownSleepFAIL(t *testing.T) {
	buffer := &bytes.Buffer{} // bit val
	spySleeper := &SpySleeper{}

	logic.CountdownSleeper(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	assert.Equal(t, want, got)
	assert.NotEqual(t, 3, spySleeper.FailedCalls) // expected: 3, got: 6
}

type SpyCountdownOperations struct {
	Calls       []string
	FailedCalls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
	s.FailedCalls = append(s.FailedCalls, write)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	s.FailedCalls = append(s.FailedCalls, sleep)
	return
}

type SpySleeper struct {
	Calls       int
	FailedCalls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpySleeper) SleepFAIL() {
	s.FailedCalls += 2
}

const write = "write"
const sleep = "sleep"

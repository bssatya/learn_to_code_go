package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("Print 3 to go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
	2
	1
	Go!`
		if got != want {
			t.Errorf("Got (%s) want (%s)", got, want)
		}
	})

	t.Run("sleep after every pring", func(t *testing.T) {
		spySleepPrinter := CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Got (%s), Want (%s)", spySleepPrinter.Calls, want)
		}
	})
}

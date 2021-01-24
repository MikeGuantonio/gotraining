package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(bytes []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("Should print output for each step", func(t *testing.T) {
		spy := &CountdownOperationsSpy{}
		buffer := &bytes.Buffer{}

		Countdown(buffer, spy)

		want := `3
2
1
Go
`
		got := buffer.String()

		if want != got {
			t.Errorf("Wanted %s, Got %s", want, got)
		}
	})

	t.Run("Should Sleep Before Every Print", func(t *testing.T) {
		spy := &CountdownOperationsSpy{}

		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("Wanted %v, Got %v", want, spy.Calls)
		}
	})
}

func TestConfiguratableSleeper(t *testing.T) {
	t.Run("Should sleep for a specified duration", func(t *testing.T){
		spy := &SpyTime{}
		sleepTime := 5 * time.Second

		sleeper := ConfigurableSleeper {
			duration: sleepTime,
			sleep: spy.Sleep,
		}

		sleeper.Sleep()

		if (sleepTime != spy.durationSlept) {
			t.Errorf("Wanted to sleep for %v but slept for %v", sleepTime, spy.durationSlept)
		}
	})
}

package main

import (
	"reflect"
	"testing"
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

func TestCountdown(t *testing.T) {
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

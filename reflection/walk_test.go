package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Nick"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string){
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("Wrong number of function calls. Wanted %d, Got %d", 1, len(got))
	}

	if got[0] != expected {
		t.Errorf("Wanted %s, Got %s", expected, got[0])
	}
}
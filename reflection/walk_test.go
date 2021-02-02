package main

import (
	"testing"
	"reflect"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one name field", 
			struct {
				Name string
			}{"Nick"}, 
			[]string{"Nick"},
		},
		{
			"Struct with two string fields", 
			struct {
				Name string
				City string
			}{"Nick", "Monaca"},
			[]string{"Nick", "Monaca"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T){
			var got []string 

			walk(test.Input, func(input string){
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
		})
	}
}
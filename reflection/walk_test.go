package main

import (
	"testing"
	"reflect"
)

type Person struct{
	Name string
	Profile Profile
}

type Profile struct{
	Age int
	City string
}

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
		{
			"Struct with number field", 
			struct {
				Name string
				Age int
			}{"Nick", 33},
			[]string{"Nick"},
		},
		{
			"Struct with nested structs", 
			Person{
				"Nick",
				Profile{33, "Monaca"},
			},
			[]string{"Nick", "Monaca"},
		},
		{
			"Struct that is a pointer", 
			&Person {
				"Nick",
				Profile{33, "Monaca"}, 
			},
			[]string{"Nick", "Monaca"},
		},
		{
			"Struct contains slices", 
			[]Profile{
				{33, "Monaca"}, 
				{34, "Center"},
			}, 
			[]string{"Monaca", "Center"},
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
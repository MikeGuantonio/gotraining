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

func assertContains(t *testing.T, got []string, input string){
	t.Helper()
	contains := false
	for _, data := range got {
		if data == input {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Could not find all keys in return")
	}	
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

	t.Run("Testing Maps", func(t *testing.T){
		input := map[string]string {
			"Foo": "bar",
			"Bar": "foo",
		}
		var got []string

		walk(input, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "foo")
	})

	t.Run("With Channels", func(t *testing.T){
		var aChan = make(chan(Profile))

		go func() {
			aChan <- Profile{33, "Monaca"}
			aChan <- Profile{33, "Moon"}
			close(aChan)
		}()

		var got []string
		want := []string{"Monaca", "Moon"}

		walk(aChan, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, Want %v", got, want)
		}
	})
}
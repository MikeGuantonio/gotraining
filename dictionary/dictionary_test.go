package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Should find a word in the dictionary" ,func(t *testing.T){
		dictionary := Dictionary{"test": "this is only a test"}

		got, _ := dictionary.Search("test")
		want := "this is only a test"

		assertStrings(t, got, want)
	})

	t.Run("Should return an message to the user if no term is found", func(t *testing.T){
		dictionary := Dictionary{"test": "this is only a test"}

		_, err := dictionary.Search("west")
		
		assertError(t, err)
	})
	
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("foo", "Oftentimes a bar")

	got, err := dictionary.Search("foo")
	want := "Oftentimes a bar"

	if err != nil {
		t.Errorf("Wanted to find a word but no word found")
	}

	assertStrings(t, got, want)
}

func assertError(t *testing.T, got error) {
	t.Helper()

	if got != ErrNotFound {
		t.Errorf("Want %q, Got %q", ErrNotFound, got)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

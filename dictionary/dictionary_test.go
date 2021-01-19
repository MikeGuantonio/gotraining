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
		

		if err == nil {
			t.Fatal("Should recieve an error")
		}

		want := "could not find the word you were looking for"
		assertStrings(t, err.Error(), want)
	})
	
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Should find a word in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is only a test"}

		got, _ := dictionary.Search("test")
		want := "this is only a test"

		assertStrings(t, got, want)
	})

	t.Run("Should return an message to the user if no term is found", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is only a test"}

		_, err := dictionary.Search("west")

		assertError(t, ErrNotFound, err)
	})

}

func TestAdd(t *testing.T) {

	t.Run("Should add a term to a dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := "Oftentimes a bar"
		word := "foo"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Should not overwrite a term already in dictionary", func(t *testing.T) {
		word := "my word"
		definition := "Not your word"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, ErrWordExists, err)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Should update a word", func(t *testing.T) {
		word := "Nick"
		newDefinition := "A sexy man"
		dictionary := Dictionary{word: "A bro person"}

		err := dictionary.Update(word, newDefinition)

		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Should not update a word that is not present", func(t *testing.T){
		word := "Nick"
		newDefinition := "A sexy man"
		dictionary := Dictionary{}

		err := dictionary.Update(word, newDefinition)

		assertError(t, ErrDoesNotExist, err)
	})
	
}

func assertError(t *testing.T, want error, got error) {
	t.Helper()

	if got != want {
		t.Errorf("Want %q, Got %q", want, got)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word string, definition string) {
	t.Helper()
	got, err := dict.Search(word)

	if err != nil {
		t.Errorf("Wanted to find a word but no word found")
	}

	assertStrings(t, got, definition)
}

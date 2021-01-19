package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is only a test"}

	got := Search(dictionary, "test")
	want := "this is only a test"

	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

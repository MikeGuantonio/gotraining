package integers

import (
	"testing";
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	want := 4

	if sum != want {
		t.Errorf("Got %d Want %d", sum, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output: 6
}
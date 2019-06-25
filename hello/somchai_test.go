package hello

import (
	"testing"
)

func TestGoodAfternoonToSomchai(t *testing.T) {
	expected := "Very Good Afternoon. Somchai."
	actual := Greeting()
	if expected != actual {
		t.Errorf("Expected %q but %q", expected, actual)
	}
}

package message

import (
	"testing"
)

func TestGoodAfternoon(t *testing.T) {
	expected := "Good Afternoon."
	actual := GoodAfternoon()
	if expected != actual {
		t.Errorf("Expected %q but %q", expected, actual)
	}
}

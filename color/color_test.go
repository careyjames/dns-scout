package color

import (
	"testing"
)

func TestYellow(t *testing.T) {
	text := "Hello, Yellow!"
	expected := "\033[38;5;3mHello, Yellow!\033[0m"
	result := Yellow(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestGreen(t *testing.T) {
	text := "Hello, Grey!"
	expected := "\033[38;5;2mHello, Grey!\033[0m"
	result := Grey(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestBlue(t *testing.T) {
	text := "Hello, Blue!"
	expected := "\033[38;5;4mHello, Blue!\033[0m"
	result := Blue(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestRed(t *testing.T) {
	text := "Hello, Red!"
	expected := "\033[38;5;1mHello, Red!\033[0m"
	result := Red(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

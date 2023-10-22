package color

import (
	"testing"
)

func TestYellow(t *testing.T) {
	text := "Hello, Yellow!"
	expected := "\033[38;5;222mHello, Yellow!\033[0m"
	result := Yellow(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestGreen(t *testing.T) {
	text := "Hello, Green!"
	expected := "\033[38;5;78mHello, Green!\033[0m"
	result := Green(text)

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
	expected := "\033[38;5;88mHello, Red!\033[0m"
	result := Red(text)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

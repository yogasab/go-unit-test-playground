package helpers

import (
	"testing"
)

func TestHelloWorldFirstName(t *testing.T) {
	result := HelloWorld("Yoga")
	if result != "Hello Yoga" {
		panic("The result must be 'Hello Yoga'")
	}
}

func TestHelloWorldLastName(t *testing.T) {
	result := HelloWorld("Baskoro")
	if result != "Hello Baskoro" {
		panic("The result must be 'Hello Baskoro'")
	}
}

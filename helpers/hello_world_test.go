package helpers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
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

func TestHelloWorldFirstNameAssertion(t *testing.T) {
	result := HelloWorld("Yoga")
	assert.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	log.Println("DONE")
}

func TestHelloWorldLastNameRequire(t *testing.T) {
	result := HelloWorld("Baskoro")
	require.Equal(t, "Hello Baskoro", result, "The result must be 'Hello Baskoro'")
	log.Println("DONE")
}

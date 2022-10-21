package helpers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"runtime"
	"testing"
)

// Unit Test Main run only on current package not in entire package
func TestMain(m *testing.M) {
	log.Println("BEFORE UNIT TEST ...")
	log.Println("CONNECTING TO DATABASE CLUSTER ...")

	m.Run()

	log.Println("AFTER UNIT TEST ...")
	log.Println("CLEAN UP THE LOG ...")
}

// Unit Test
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

// Unit Test Assertion
func TestHelloWorldFirstNameAssertion(t *testing.T) {
	result := HelloWorld("Yoga")
	assert.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	log.Println("DONE")
}

// Unit Test Assertion
func TestHelloWorldLastNameRequire(t *testing.T) {
	result := HelloWorld("Baskoro")
	require.Equal(t, "Hello Baskoro", result, "The result must be 'Hello Baskoro'")
	log.Println("DONE")
}

// Unit Test Skip
func TestSkipOnMac(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("This test cannot run on Mac")
	}
	result := HelloWorld("Yoga")
	assert.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	log.Println("DONE")
}

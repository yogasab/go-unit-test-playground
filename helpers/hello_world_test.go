package helpers

import (
	"fmt"
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
func TestHelloToFirstName(t *testing.T) {
	result := HelloTo("Yoga")
	if result != "Hello Yoga" {
		panic("The result must be 'Hello Yoga'")
	}
}

func TestHelloToLastName(t *testing.T) {
	result := HelloTo("Baskoro")
	if result != "Hello Baskoro" {
		panic("The result must be 'Hello Baskoro'")
	}
}

// Unit Test Assertion
func TestHelloToFirstNameAssertion(t *testing.T) {
	result := HelloTo("Yoga")
	assert.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	log.Println("DONE")
}

// Unit Test Assertion
func TestHelloToLastNameRequire(t *testing.T) {
	result := HelloTo("Baskoro")
	require.Equal(t, "Hello Baskoro", result, "The result must be 'Hello Baskoro'")
	log.Println("DONE")
}

// Unit Test Skip
func TestSkipOnMac(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("This test cannot run on Mac")
	}
	result := HelloTo("Yoga")
	assert.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	log.Println("DONE")
}

// Unit Test Sub Test
func TestSubTestHelloTo(t *testing.T) {
	t.Run("It should return 'Hello Yoga'", func(t *testing.T) {
		result := HelloTo("Yoga")
		require.Equal(t, "Hello Yoga", result, "The result must be 'Hello Yoga'")
	})
	t.Run("It should return 'Hello Baskoro'", func(t *testing.T) {
		result := HelloTo("Baskoro")
		require.Equal(t, "Hello Baskoro", result, "The result must be 'Hello Baskoro'")
	})
	t.Run("It should return 'Hello Yoga Baskoro'", func(t *testing.T) {
		result := HelloTo("Yoga Baskoro")
		require.Equal(t, "Hello Yoga Baskoro", result, "The result must be 'Hello Yoga Baskoro'")
	})
}

// Unit Test Table Test to reduce repetitive of tests
func TestSubTableTest(t *testing.T) {
	tests := []struct {
		name      string
		requested string
		expected  string
	}{
		{
			name:      "It should return 'Hello Yoga'",
			requested: "Yoga",
			expected:  "Hello Yoga",
		}, {
			name:      "It should return 'Hello Baskoro'",
			requested: "Baskoro",
			expected:  "Hello Baskoro",
		}, {
			name:      "It should return 'Hello Yoga Baskoro'",
			requested: "Yoga Baskoro",
			expected:  "Hello Yoga Baskoro",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloTo(test.requested)
			require.Equal(t, test.expected, result, fmt.Sprintf("The result must be 'Hello %s'", test.requested))
		})
	}
}

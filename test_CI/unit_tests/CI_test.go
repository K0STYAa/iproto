package unittests_test

import "testing"

func TestHello(t *testing.T) {
	t.Parallel()
	// Arrange
	testTable := []struct {
		name     string
		expected string
	}{
		{
			name:     "Golang",
			expected: "Hello, Golang!",
		},
		{
			name:     "Kostya",
			expected: "Hello, Kostya!",
		},
		{
			name:     "",
			expected: "Hello, !",
		},
	}

	// Act
	for _, testCase := range testTable {
		result := hello(testCase.name)

		t.Logf("Calling hello(%s), result %s\n",
			testCase.name, result)

		// Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expected %s, got %s",
				testCase.expected, result)
		}
	}
}

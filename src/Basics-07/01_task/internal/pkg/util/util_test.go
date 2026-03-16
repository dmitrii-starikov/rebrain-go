package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseInt(t *testing.T) {
	_, expectedErr := ReverseInt("abc")
	if expectedErr == nil {
		t.Fatal("fail")
	}

	answer, expectedErr := ReverseInt(123)
	if answer != 321 || expectedErr != nil {
		t.Fatal("fail")
	}

	answer, expectedErr = ReverseInt(-649)
	if answer != -946 || expectedErr != nil {
		t.Fatal("fail")
	}

	answer, expectedErr = ReverseInt(0)
	if answer != 0 || expectedErr != nil {
		t.Fatal("fail")
	}
}

func TestReverseInt_WithTestify(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    int
		expectedErr error
	}{
		{
			name:        "positive number",
			input:       123,
			expected:    321,
			expectedErr: nil,
		},
		{
			name:        "negative number",
			input:       -456,
			expected:    -654,
			expectedErr: nil,
		},
		{
			name:        "number ending with zero",
			input:       120,
			expected:    21,
			expectedErr: nil,
		},
		{
			name:        "single digit",
			input:       5,
			expected:    5,
			expectedErr: nil,
		},
		{
			name:        "zero",
			input:       0,
			expected:    0,
			expectedErr: nil,
		},
		{
			name:        "not int type (string)",
			input:       "123",
			expected:    0,
			expectedErr: errors.New("not int"),
		},
		{
			name:        "not int type (float)",
			input:       123.45,
			expected:    0,
			expectedErr: errors.New("not int"),
		},
		{
			name:        "overflow positive",
			input:       2147483647, // max int32
			expected:    0,          // reverse is 7463847412 > int32
			expectedErr: nil,
		},
		{
			name:        "overflow negative",
			input:       -2147483648, // min int32
			expected:    0,           // reverse is -8463847412 < int32
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, expectedErr := ReverseInt(tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, expectedErr)
				assert.EqualError(t, expectedErr, tt.expectedErr.Error())
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, expectedErr)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestReverseInt_EdgeCases(t *testing.T) {
	t.Run("palindrome number", func(t *testing.T) {
		result, expectedErr := ReverseInt(121)
		assert.NoError(t, expectedErr)
		assert.Equal(t, 121, result)
	})

	t.Run("large number within bounds", func(t *testing.T) {
		result, expectedErr := ReverseInt(123456789)
		assert.NoError(t, expectedErr)
		assert.Equal(t, 987654321, result)
	})

	t.Run("negative number with zero", func(t *testing.T) {
		result, expectedErr := ReverseInt(-120)
		assert.NoError(t, expectedErr)
		assert.Equal(t, -21, result)
	})

	t.Run("nil", func(t *testing.T) {
		result, expectedErr := ReverseInt(nil)

		assert.Error(t, expectedErr)
		assert.EqualError(t, expectedErr, errors.New("not int").Error())
		assert.Equal(t, 0, result)
	})
}

func TestReverseInt_TypeAssertions(t *testing.T) {
	t.Run("type assertion failure", func(t *testing.T) {
		_, expectedErr := ReverseInt(nil)
		assert.Error(t, expectedErr)
		assert.EqualError(t, expectedErr, "not int")
	})

	t.Run("type assertion with custom type", func(t *testing.T) {
		type MyInt int
		var myInt MyInt = 123

		_, expectedErr := ReverseInt(myInt)
		assert.Error(t, expectedErr)
		assert.EqualError(t, expectedErr, "not int")
	})
}

//////////////////////

// TestContainsDuplicate - Table Driven Tests
func TestContainsDuplicate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "empty slice",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "has duplicates at beginning",
			nums:     []int{1, 1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "has duplicates at end",
			nums:     []int{1, 2, 3, 4, 4},
			expected: true,
		},
		{
			name:     "has duplicates in middle",
			nums:     []int{1, 2, 3, 3, 4, 5},
			expected: true,
		},
		{
			name:     "multiple duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: true,
		},
		{
			name:     "all same elements",
			nums:     []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "negative numbers without duplicates",
			nums:     []int{-1, -2, -3, -4},
			expected: false,
		},
		{
			name:     "negative numbers with duplicates",
			nums:     []int{-1, -2, -2, -3},
			expected: true,
		},
		{
			name:     "zero with duplicates",
			nums:     []int{0, 1, 2, 0, 3},
			expected: true,
		},
		{
			name:     "zero without duplicates",
			nums:     []int{0, 1, 2, 3, 4},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := ContainsDuplicate(tc.nums)
			require.Equal(t, tc.expected, result)
		})
	}
}

// TestIsPalindrome - Closure Driven Tests
func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	// Closure-constructor
	palindromeCase := func(x int, expected bool) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			result := IsPalindrome(x)
			require.Equal(t, expected, result, "IsPalindrome(%d) should be %v", x, expected)
		}
	}

	t.Run("positive palindrome single digit", palindromeCase(5, true))
	t.Run("positive palindrome two same digits", palindromeCase(11, true))
	t.Run("positive palindrome three digits", palindromeCase(121, true))
	t.Run("positive palindrome four digits", palindromeCase(1221, true))
	t.Run("positive palindrome five digits", palindromeCase(12321, true))
	t.Run("positive palindrome even digits", palindromeCase(123321, true))
	t.Run("positive palindrome with zero", palindromeCase(101, true))
	t.Run("positive palindrome 1001", palindromeCase(1001, true))
	t.Run("positive non-palindrome", palindromeCase(123, false))
	t.Run("positive non-palindrome two digits", palindromeCase(12, false))
	t.Run("positive non-palindrome four digits", palindromeCase(1234, false))
	t.Run("positive non-palindrome with zero", palindromeCase(120, false))
	t.Run("zero is palindrome", palindromeCase(0, true))

	// IsPalindrome: negative numbers are not palindromes
	t.Run("negative number always false", palindromeCase(-121, false))
	t.Run("negative single digit", palindromeCase(-5, false))
	t.Run("negative palindrome-like", palindromeCase(-101, false))

	// Edge cases - ending with 0
	t.Run("number ending with zero but not zero", palindromeCase(10, false))
	t.Run("number ending with zero 100", palindromeCase(100, false))
	t.Run("number ending with zero 1010", palindromeCase(1010, false))

	// Big numbers
	t.Run("large palindrome", palindromeCase(123454321, true))
	t.Run("large non-palindrome", palindromeCase(123456789, false))
	t.Run("max int palindrome", palindromeCase(2147447412, true))
}

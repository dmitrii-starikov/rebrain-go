package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	_, err := ReverseInt("abc")
	if err == nil {
		t.Fatal("fail")
	}

	answer, err := ReverseInt(123)
	if answer != 321 || err != nil {
		t.Fatal("fail")
	}

	answer, err = ReverseInt(-649)
	if answer != -946 || err != nil {
		t.Fatal("fail")
	}

	answer, err = ReverseInt(0)
	if answer != 0 || err != nil {
		t.Fatal("fail")
	}
}

func TestReverseInt_WithTestify(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
		err      error
	}{
		{
			name:     "positive number",
			input:    123,
			expected: 321,
			err:      nil,
		},
		{
			name:     "negative number",
			input:    -456,
			expected: -654,
			err:      nil,
		},
		{
			name:     "number ending with zero",
			input:    120,
			expected: 21,
			err:      nil,
		},
		{
			name:     "single digit",
			input:    5,
			expected: 5,
			err:      nil,
		},
		{
			name:     "zero",
			input:    0,
			expected: 0,
			err:      nil,
		},
		{
			name:     "not int type (string)",
			input:    "123",
			expected: 0,
			err:      errors.New("not int"),
		},
		{
			name:     "not int type (float)",
			input:    123.45,
			expected: 0,
			err:      errors.New("not int"),
		},
		{
			name:     "overflow positive",
			input:    2147483647, // max int32
			expected: 0,          // reverse is 7463847412 > int32
			err:      nil,
		},
		{
			name:     "overflow negative",
			input:    -2147483648, // min int32
			expected: 0,           // reverse is -8463847412 < int32
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReverseInt(tt.input)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestReverseInt_EdgeCases(t *testing.T) {
	t.Run("palindrome number", func(t *testing.T) {
		result, err := ReverseInt(121)
		assert.NoError(t, err)
		assert.Equal(t, 121, result)
	})

	t.Run("large number within bounds", func(t *testing.T) {
		result, err := ReverseInt(123456789)
		assert.NoError(t, err)
		assert.Equal(t, 987654321, result)
	})

	t.Run("negative number with zero", func(t *testing.T) {
		result, err := ReverseInt(-120)
		assert.NoError(t, err)
		assert.Equal(t, -21, result)
	})

	t.Run("nil", func(t *testing.T) {
		result, err := ReverseInt(nil)

		assert.Error(t, err)
		assert.EqualError(t, err, errors.New("not int").Error())
		assert.Equal(t, 0, result)
	})
}

func TestReverseInt_TypeAssertions(t *testing.T) {
	t.Run("type assertion failure", func(t *testing.T) {
		_, err := ReverseInt(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "not int")
	})

	t.Run("type assertion with custom type", func(t *testing.T) {
		type MyInt int
		var myInt MyInt = 123

		_, err := ReverseInt(myInt)
		assert.Error(t, err)
		assert.EqualError(t, err, "not int")
	})
}

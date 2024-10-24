package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIndexService(t *testing.T) {
	tt := []struct {
		name          string
		value         int
		expected      Number
		expectedError bool
	}{
		{
			name:          "Failure for value -100",
			value:         -100,
			expected:      Number{Index: -1, Value: -100},
			expectedError: true,
		},
		{
			name:     "Success for value 1",
			value:    1,
			expected: Number{Index: 0, Value: 1},
		},
		{
			name:          "Failure for value 5",
			value:         5,
			expected:      Number{Index: -1, Value: 5},
			expectedError: true,
		},
		{
			name:     "Success for value 65",
			value:    65,
			expected: Number{Index: 5, Value: 60},
		},
		{
			name:     "Success for value 66",
			value:    66,
			expected: Number{Index: 5, Value: 60},
		},
		{
			name:          "Failure for value 67",
			value:         67,
			expected:      Number{Index: -1, Value: 67},
			expectedError: true,
		},
		{
			name:          "Failure for value 636",
			value:         636,
			expected:      Number{Index: -1, Value: 636},
			expectedError: true,
		},
		{
			name:     "Success for value 637",
			value:    637,
			expected: Number{Index: 6, Value: 700},
		},
		{
			name:     "Success for value 840",
			value:    840,
			expected: Number{Index: 7, Value: 800},
		},
		{
			name:     "Success for value 900",
			value:    900,
			expected: Number{Index: 8, Value: 900},
		},
		{
			name:     "Success for value 910",
			value:    910,
			expected: Number{Index: 8, Value: 900},
		},
		{
			name:     "Success for value 1000",
			value:    1000,
			expected: Number{Index: 9, Value: 1000},
		},
		{
			name:     "Success for value 1090",
			value:    1090,
			expected: Number{Index: 9, Value: 1000},
		},
		{
			name:     "Success for value 1100",
			value:    1100,
			expected: Number{Index: 9, Value: 1000},
		},
		{
			name:          "Failure for value 2000",
			value:         2000,
			expected:      Number{Index: -1, Value: 2000},
			expectedError: true,
		},
	}
	testData := []int{
		1,
		2,
		3,
		40,
		50,
		60,
		700,
		800,
		900,
		1000,
	}
	testConformationLevel := 10

	nopLogger := zap.NewNop()
	repo := NewNumbersSliceRepository(testData, testConformationLevel, nopLogger)
	service := NewIndexService(repo, nopLogger)
	for idx, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gotNumber, err := service.GetIndex(tc.value)
			if tc.expectedError {
				assert.ErrorIs(t, err, ErrNotFound, "Test case %d", idx+1)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, gotNumber, "Test case %d", idx+1)
		})
	}
}

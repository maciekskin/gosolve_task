package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexService(t *testing.T) {
	tt := []struct {
		name          string
		value         int
		expectedIndex int
		expectedError bool
	}{
		{
			name:          "Failure for value -100",
			value:         -100,
			expectedIndex: -1,
			expectedError: true,
		},
		{
			name:          "Success for value 1",
			value:         1,
			expectedIndex: 0,
		},
		{
			name:          "Failure for value 5",
			value:         5,
			expectedIndex: -1,
			expectedError: true,
		},
		{
			name:          "Success for value 65",
			value:         65,
			expectedIndex: 5,
		},
		{
			name:          "Success for value 840",
			value:         840,
			expectedIndex: 7,
		},

		{
			name:          "Success for value 900",
			value:         900,
			expectedIndex: 8,
		},
		{
			name:          "Success for value 910",
			value:         910,
			expectedIndex: 8,
		},
		{
			name:          "Success for value 1000",
			value:         1000,
			expectedIndex: 9,
		},
		{
			name:          "Success for value 1090",
			value:         1090,
			expectedIndex: 9,
		},
		{
			name:          "Failure for value 1100",
			value:         1100,
			expectedIndex: -1,
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

	repo := NewNumbersSliceRepository(testData, testConformationLevel)
	service := NewIndexService(repo)
	for idx, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gotIndex, err := service.GetIndex(tc.value)
			if tc.expectedError {
				assert.ErrorIs(t, err, ErrNotFound, "Test case %d", idx+1)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedIndex, gotIndex, "Test case %d", idx+1)
		})
	}
}

package merge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"primo-merge-func/internal/merge"
)

type MergeCollectionTestCase struct {
	name        string
	collection1 []int
	collection2 []int
	collection3 []int
	expected    []int
}

func TestMergeCollection(t *testing.T) {
	table := []MergeCollectionTestCase{
		{
			name:        "Should return empty slices",
			collection1: []int{},
			collection2: []int{},
			collection3: []int{},
			expected:    []int{},
		},
		{
			name:        "Should merge ascending collection and return ascending slices",
			collection1: []int{},
			collection2: []int{1, 2, 3, 4},
			collection3: []int{8, 9, 10},
			expected:    []int{1, 2, 3, 4, 8, 9, 10},
		},
		{
			name:        "Should merge all collection and return ascending slices",
			collection1: []int{7, 6, 5},
			collection2: []int{1, 2, 3, 4},
			collection3: []int{8, 9, 10},
			expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:        "Should reordered the 1st collection",
			collection1: []int{4, 3, 2, 1},
			collection2: []int{},
			collection3: []int{},
			expected:    []int{1, 2, 3, 4},
		},
		{
			name:        "Should merge collection collectly when negative interger in 1st collection",
			collection1: []int{-1, -2, -3, -4},
			collection2: []int{1, 2, 3, 4},
			collection3: []int{0},
			expected:    []int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
		},
		{
			name:        "Should merge in between",
			collection1: []int{99, 66, 4, -99},
			collection2: []int{1, 2, 3, 10},
			collection3: []int{22, 77, 88},
			expected:    []int{-99, 1, 2, 3, 4, 10, 22, 66, 77, 88, 99},
		},
		{
			name: "Should merge same int coll",
			collection1: []int{7, 7, 7},
			collection2: []int{7, 7},
			collection3: []int{7},
			expected:    []int{7, 7, 7, 7, 7, 7},
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			actual := merge.MergeCollection(
				tc.collection1,
				tc.collection2,
				tc.collection3,
			)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

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

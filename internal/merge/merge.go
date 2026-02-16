package merge


func MergeCollection(
	collection1,
	collection2,
	collection3 []int,
) []int {
	size := len(collection1) + len(collection2) + len(collection3)
	mergedCollection := make([]int, size)

	return mergedCollection
}
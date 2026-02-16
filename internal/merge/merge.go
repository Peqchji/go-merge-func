package merge


func MergeCollection(
	collection1,
	collection2,
	collection3 []int,
) []int {
	mergedCollection := mergeHelper(collection2, collection3)

	return mergedCollection
}

func mergeHelper(coll1, coll2 []int) []int {
	lenColl1 := len(coll1) 
	lenColl2 := len(coll2)
	mergedColl := make([]int, lenColl1 + lenColl2)

	i :=  0
	ptr1 := 0
	ptr2 := 0

	for ptr1 != lenColl1 && ptr2 != lenColl2 {
		if coll1[ptr1] < coll2[ptr2] {
			mergedColl[i] = coll1[ptr1]
			ptr1 += 1
		} else {
			mergedColl[i] = coll2[ptr2]
			ptr2 += 1
		}

		i += 1
	}

	for ptr1 < lenColl1 {
		mergedColl[i] = coll1[ptr1]
		ptr1 += 1
		i += 1
	}

	for ptr2 < lenColl2 {
		mergedColl[i] = coll2[ptr2]
		ptr2 += 1
		i += 1
	}

	return mergedColl
}
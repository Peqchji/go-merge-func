package merge


func MergeCollection(
	coll1,
	coll2,
	coll3 []int,
) []int {
	lenColl1 := len(coll1) 
	lenColl2 := len(coll2)
	lenColl3 := len(coll3)
	mergedColl := make([]int, lenColl1 + lenColl2 + lenColl3)

	i :=  0
	ptr1 := lenColl1 - 1
	ptr2 := 0
	ptr3 := 0

	for ptr1 != -1 && ptr2 != lenColl2 && ptr3 != lenColl3 {
		num1, num2, num3 := coll1[ptr1], coll2[ptr2], coll3[ptr3]

		if num1 <= num2 && num1 <= num3 {
			mergedColl[i] = num1
			ptr1 -= 1
		} else if num2 <= num1 && num2 <= num3{
			mergedColl[i] = num2
			ptr2 += 1
		} else if num3 <= num1 && num3 <= num2{
			mergedColl[i] = num3
			ptr3 += 1
		}

		i += 1
	}

	for ptr1 != -1 && ptr2 != lenColl2 {
		num1, num2 := coll1[ptr1], coll2[ptr2]

		if num1 <= num2{
			mergedColl[i] = num1
			ptr1 -= 1
		} else {
			mergedColl[i] = num2
			ptr2 += 1
		} 

		i += 1
	}

	for ptr1 != -1 && ptr3 != lenColl3 {
		num1, num3 := coll1[ptr1], coll3[ptr3]

		if num1 <= num3 {
			mergedColl[i] = num1
			ptr1 -= 1
		} else {
			mergedColl[i] = num3
			ptr3 += 1
		} 

		i += 1
	}

	for ptr2 != lenColl2 && ptr3 != lenColl3 {
		num2, num3 := coll2[ptr2], coll3[ptr3]

		if num2 <= num3 {
			mergedColl[i] = num2
			ptr2 += 1
		} else {
			mergedColl[i] = num3
			ptr3 += 1
		} 

		i += 1
	}

	for ptr1 >= 0 {
		mergedColl[i] = coll1[ptr1]
		ptr1 -= 1
		i += 1
	}

	for ptr2 < lenColl2 {
		mergedColl[i] = coll2[ptr2]
		ptr2 += 1
		i += 1
	}

	for ptr3 < lenColl3 {
		mergedColl[i] = coll3[ptr3]
		ptr3 += 1
		i += 1
	}

	return mergedColl
}
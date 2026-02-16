package merge

import (
	"math"
)


func MergeCollection(
	coll1,
	coll2,
	coll3 []int,
) []int {
	lenColl1 := len(coll1) 
	lenColl2 := len(coll2)
	lenColl3 := len(coll3)
	size := lenColl1 + lenColl2 + lenColl3
	mergedColl := make([]int, lenColl1 + lenColl2 + lenColl3)

	ptr1 := lenColl1 - 1
	ptr2 := 0
	ptr3 := 0

	for i := range size {
		num1, num2, num3 := math.MaxInt, math.MaxInt, math.MaxInt
 
		if ptr1 >= 0 {
            num1 = coll1[ptr1]
        }
        if ptr2 < len(coll2) {
            num2 = coll2[ptr2]
        }
        if ptr3 < len(coll3) {
            num3 = coll3[ptr3]
        }

		if num1 <= num2 && num1 <= num3 {
            mergedColl[i] = num1
            ptr1 -= 1
        } else if num2 <= num1 && num2 <= num3 {
            mergedColl[i] = num2
            ptr2 +=1
        } else {
            mergedColl[i] = num3
            ptr3 += 1
        }
	}

	return mergedColl
}
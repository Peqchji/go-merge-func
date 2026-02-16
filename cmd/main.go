package main

import (
	"fmt"
	"primo-merge-func/internal/merge"
)

func main() {
	//  given
    //  collection_1 already sorted from max to min
    //  collection_2, collection_3 already sorted from min(0) to max
	coll1 := []int{10, 9}
	coll2 := []int{6, 7}
	coll3 := []int{0, 8}

	merged := merge.MergeCollection(coll1, coll2, coll3)
	fmt.Printf("merged collection: %v", merged)
}
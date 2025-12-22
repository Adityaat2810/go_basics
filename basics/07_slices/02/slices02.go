package main

import (
	"fmt"
	"slices"
)


func main() {
	nums := make([]int, 0, 5)
	nums = append(nums, 2)

	nums2 := make([]int, len(nums)) // here len(nums) is 1


	/********************** copy function **********************/
	copy(nums2, nums)   // copy(source, destination)


	fmt.Println("nums is :-",nums, "nums2 is :-",nums2)

	/************************ slice operator **********************/
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl[0:2])   // [i, j] --> start from index i and end on j-1
	// [i, j)

	fmt.Println(sl[:2]) // if we don't give from index it will by default take it 0
	// if we won't give end itdex it will pick size-1 --> print till end

	/************************ slices package **************************/
	sl2 := []int {1, 2}
	sl3 := []int {1, 2}
	fmt.Println(slices.Equal(sl2, sl3)) // return true


	// 2d slice
	var ds2 = [][]int{{1, 2, 3, 4, 4}, {1, 3, 3, 7}}
	fmt.Println(ds2)



}

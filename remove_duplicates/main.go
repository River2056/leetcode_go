package remove_duplicates

import "fmt"

func RemoveDuplicates(nums []int) int {
	placeholder := -999
	duplicates := 0
	if len(nums) == 0 {
		return 0
	}
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == temp {
			nums[i] = placeholder
			duplicates++
		} else {
			temp = nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == placeholder {
			nums = append(nums[:i], nums[i + 1:]...)
		}
	}
	fmt.Println(nums)
	return len(nums)
}
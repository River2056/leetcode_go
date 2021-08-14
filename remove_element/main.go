package remove_element

import "fmt"

func RemoveElement(nums []int, val int) int {
	placeholder := -999
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = placeholder
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

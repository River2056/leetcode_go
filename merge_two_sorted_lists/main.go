package merge_two_sorted_lists

import (
	"fmt"
	l "leetcode_go/add_two_numbers"
)

func GetDummyData() (*l.ListNode, *l.ListNode) {
	l1 := l.ListNode{ Val: 1 }
	l2 := l.ListNode{ Val: 2 }
	l3 := l.ListNode{ Val: 4 }
	l1.Next = &l2
	l2.Next = &l3

	l11 := l.ListNode{ Val: 1 }
	l12 := l.ListNode{ Val: 3 }
	l13 := l.ListNode{ Val: 4 }
	l11.Next = &l12
	l12.Next = &l13
	return &l1, &l11
}

func getLinkListValuesIntoSlice(l1 *l.ListNode, arr []int) []int {
	current := l1
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

func MergeTwoLists(l1 *l.ListNode, l2 *l.ListNode) *l.ListNode {
	var nums []int = make([]int, 0)
	nums = getLinkListValuesIntoSlice(l1, nums)
	nums = getLinkListValuesIntoSlice(l2, nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {

			}
		}
	}
	fmt.Println(nums)
	return nil
}

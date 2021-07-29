package merge_two_sorted_lists

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func GetDummyData() (*ListNode, *ListNode) {
	l1 := ListNode{ Val: 1 }
	l2 := ListNode{ Val: 2 }
	l3 := ListNode{ Val: 4 }
	l1.Next = &l2
	l2.Next = &l3

	l11 := ListNode{ Val: 1 }
	l12 := ListNode{ Val: 3 }
	l13 := ListNode{ Val: 4 }
	l11.Next = &l12
	l12.Next = &l13

	//l1 := ListNode{}
	//l11 := ListNode{}

	return &l1, &l11
}

func getLinkListValuesIntoSlice(l1 *ListNode, arr []int) []int {
	current := l1
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var nums []int = make([]int, 0)
	nums = getLinkListValuesIntoSlice(l1, nums)
	nums = getLinkListValuesIntoSlice(l2, nums)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j + 1] {
				tmp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = tmp
			}
		}
	}

	var firstNode *ListNode
	var lastNode *ListNode = nil
	for i := len(nums) - 1; i >= 0; i-- {
		node := ListNode{ Val: nums[i] }
		if i == 0 {
			firstNode = &node
		}
		node.Next = lastNode
		lastNode = &node
	}
	fmt.Println(nums)
	return firstNode
}

func PrintLinkList(node *ListNode) {
	var current *ListNode
	if node != nil { current = node }
	for current != nil {
		fmt.Println(current.Val)
		fmt.Println(current.Next)
		current = current.Next
	}
}

//func main() {
	//	l1, l11 := GetDummyData()
	//	result := MergeTwoLists(l1, l11)
	//	fmt.Println(result)
	//	PrintLinkList(result)
	//}
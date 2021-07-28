package main

import (
	"fmt"
	"leetcode_go/merge_two_sorted_lists"
)

func main() {
	l1, l11 := merge_two_sorted_lists.GetDummyData()
	result := merge_two_sorted_lists.MergeTwoLists(l1, l11)
	fmt.Println(result)
	//l.PrintLinkList(result)
}
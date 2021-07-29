package merge_two_sorted_lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1, l11 := GetDummyData()
	result := MergeTwoLists(l1, l11)
	PrintLinkList(result)
}

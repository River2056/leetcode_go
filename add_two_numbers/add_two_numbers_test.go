package add_two_numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1, l11 := GetDummyData()
	result := AddTwoNumbers(&l1, &l11)
	PrintLinkList(result)
}
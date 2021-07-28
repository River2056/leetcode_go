package add_two_numbers

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func generateLinkList(arr []int) *ListNode {
	var result string = ""
	for _, n := range arr {
		result += strconv.Itoa(n)
	}
	resultList := getLinkList(result)
	return resultList
}

func GetDummyData() (ListNode, ListNode) {
	// [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	//	[5,6,4]
	l1 := generateLinkList([]int { 1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1 })
	l11 := generateLinkList([]int { 5,6,4 })

	return *l1, *l11
}

func getStringList(node *ListNode) []string {
	list := make([]string, 0)
	list = append(list, strconv.Itoa(node.Val))
	nextNode := node.Next
	for nextNode != nil {
		nextVal := nextNode.Val
		list = append(list, strconv.Itoa(nextVal))
		nextNode = nextNode.Next
	}
	return list
}

func getLinkList(sumStr string) *ListNode {
	var firstNode *ListNode
	var lastNode *ListNode
	for i := len(sumStr) - 1; i >= 0; i-- {
		node := ListNode{}
		val, _ := strconv.Atoi(string(sumStr[i]))
		node.Val = val
		if lastNode != nil {
			node.Next = lastNode
		}
		if i == 0 {
			firstNode = &node
		}
		lastNode = &node
	}
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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := getStringList(l1)
	list2 := getStringList(l2)
	var str1 string
	var str2 string

	for i := len(list1) - 1; i >= 0; i-- {
		str1 += list1[i]
	}

	for i := len(list2) - 1; i >= 0; i-- {
		str2 += list2[i]
	}

	num1 := big.NewInt(0)
	num1.SetString(str1, 10)
	num2 := new(big.Int)
	num2.SetString(str2, 10)

	sum := num1.Add(num1, num2)

	sumStr := sum.String()
	var reverseStr string = ""
	for i := len(sumStr) - 1; i >= 0; i-- {
		reverseStr += string(sumStr[i])
	}
	linkList := getLinkList(reverseStr)

	fmt.Printf("sumStr: %v\n", sumStr)
	PrintLinkList(linkList)

	return linkList
}

// Add two numbers
//func main() {
//	l1, l11 := add_two_numbers.GetDummyData()
//	add_two_numbers.AddTwoNumbers(&l1, &l11)
//}
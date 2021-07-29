package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	data, target := TestData()
	result := TwoSum(data, target)
	fmt.Println(result)
}

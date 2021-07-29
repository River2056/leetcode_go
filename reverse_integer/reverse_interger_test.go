package reverse_integer

import (
	"fmt"
	"testing"
)

func testCases(t *testing.T, input, output int) {
	result := Reverse(input)
	fmt.Printf("input: %v\n", input)
	fmt.Printf("got: %v\n", result)
	if result != output {
		t.Errorf("want: %v, got %v", output, result)
	}
}

func TestReverse(t *testing.T) {
	testCases(t, -123, -321)
	testCases(t, 123, 321)
	testCases(t, 120, 21)
	testCases(t, 0, 0)
	testCases(t, 1534236469, 0)
	testCases(t, 900000, 9)
}

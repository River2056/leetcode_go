package remove_element

import "testing"

func testCase(t *testing.T, input []int, val, output int) {
	result := RemoveElement(input, val)
	if result != output {
		t.Errorf("expected: %v, got: %v", output, result)
	}
}

func TestRemoveElement(t *testing.T) {
	testCase(t, []int { 3, 2, 2, 3 }, 3, 2)
	testCase(t, []int { 0, 1, 2, 2, 3, 0, 4, 2 }, 2, 5)
}
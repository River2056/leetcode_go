package remove_duplicates

import "testing"

func testCase(t *testing.T, input []int, output int) {
	result := RemoveDuplicates(input)
	if result != output {
		t.Errorf("expected: %v, got: %v", output, result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCase(t, []int { 1, 1, 2 }, 2)
	testCase(t, []int { 0, 0, 1, 1, 1, 2, 2, 3, 3, 4 }, 5)
	testCase(t, []int { 0, 0, 0, 0, 3 }, 2)
	testCase(t, []int { 1, 1 }, 1 )
}

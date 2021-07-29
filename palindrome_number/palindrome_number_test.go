package palindrome_number

import "testing"

func testCase(t *testing.T, input int, output bool) {
	result := IsPalindrome(input)
	if result != output {
		t.Errorf("want: %v, got: %v\n", output, result)
	}
}

func TestIsPalindrome(t *testing.T) {
	testCase(t, 121, true)
	testCase(t, -121, false)
	testCase(t, 10, false)
	testCase(t, -101, false)
}

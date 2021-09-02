package implement_strStr

import (
	"testing"
)

func testCase(t *testing.T, haystack, needle string, output int) {
	result := StrStr(haystack, needle)
	if result != output {
		t.Errorf("expected: %v, got: %v\n", output, result)
	}
}

func TestStrStr(t *testing.T) {
	testCase(t, "hello", "ll", 2)
	testCase(t, "aaaaa", "bba", -1)
	testCase(t, "", "", 0)
	testCase(t, "", "a", -1)
}
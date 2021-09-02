package implement_strStr

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		subLength := i + len(needle)
		if subLength > len(haystack) {
			subLength = len(haystack)
		}
		arr := haystack[i:subLength]
		if arr == needle {
			return i
		}
	}

	return -1
}
package palindrome_number

func IsPalindrome(x int) bool {
	if x < 0 { return false }
	input := x
	reverse := 0
	for input != 0 {
		lastDigit := input % 10
		input /= 10
		reverse = (reverse * 10) + lastDigit
	}
	return reverse == x
}

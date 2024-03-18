package mystrings

// unless you capitalize the function it won't get exported
func Reverse(s string) string {
	result := ""
	for _, el := range s {
		result = string(el) + result
	}
	return result
}

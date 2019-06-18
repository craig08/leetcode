func isValid(s string) bool {
	var stack []byte
	for idx := 0; idx < len(s); idx++ {
		switch s[idx] {
		case '(', '[', '{':
			stack = append(stack, s[idx])
		case ')', ']', '}':
			if len(stack) == 0 ||
				(s[idx] == ')' && stack[len(stack)-1] != '(') ||
				(s[idx] == ']' && stack[len(stack)-1] != '[') ||
				(s[idx] == '}' && stack[len(stack)-1] != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}

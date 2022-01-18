package main

func isValid(s string) bool {
	mapping := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	if len(s) == 0 {
		return true
	}
	stack := make([]string, 0)

	for _, v := range s {
		c := string(v)

		if c == "{" || c == "[" || c == "(" {
			stack = push(stack, c)
		} else if c == "}" || c == "]" || c == ")" {
			var temp string
			temp, stack = pop(stack)
			if temp != mapping[c] {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

func push(stack []string, elements ...string) []string {
	return append(stack, elements...)
}

func pop(stack []string) (string, []string) {
	l := len(stack)
	if l > 0 {
		return stack[l-1], stack[:l-1]
	}
	return "", nil
}

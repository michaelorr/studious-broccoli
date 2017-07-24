package brackets

const testVersion = 5
const Opening = "opening"

var BraceMap = map[string]string{
	"(": Opening,
	"{": Opening,
	"[": Opening,
	")": "(",
	"}": "{",
	"]": "[",
}

func Bracket(input string) (bool, error) {
	var stack []string

	for _, r := range input {
		if a, ok := BraceMap[string(r)]; !ok {
			continue
		} else if a == Opening {
			stack = append(stack, string(r))
		} else {
			if len(stack) == 0 {
				return false, nil
			}
			match := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if match != a {
				return false, nil
			}
		}
	}
	if len(stack) != 0 {
		return false, nil
	}
	return true, nil
}

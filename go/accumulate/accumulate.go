package accumulate

const testVersion = 1

func Accumulate(data []string, oper func(string) string) []string {
	res := make([]string, len(data))
	for i, item := range data {
		res[i] = oper(item)
	}
	return res
}

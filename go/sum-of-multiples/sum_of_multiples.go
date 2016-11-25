package summultiples

func MultipleSummer(nums ...int) func(int) int {
	return func(limit int) int {
		result := 0
		for i := 1; i < limit; i++ {
			for _, n := range nums {
				if (i % n) == 0 {
					result += i
					break
				}
			}
		}
		return result
	}
}

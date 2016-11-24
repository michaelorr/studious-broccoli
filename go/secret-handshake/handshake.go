package secret

const (
	Wink int = 1 << iota
	DblBlink
	CloseEyes
	Jump
	Reverse
)

func Handshake(code int) []string {
	res := *new([]string)

	if code < 0 {
		code = 0
	}
	if codePassed(code, Wink) {
		res = append(res, "wink")
	}
	if codePassed(code, DblBlink) {
		res = append(res, "double blink")
	}
	if codePassed(code, CloseEyes) {
		res = append(res, "close your eyes")
	}
	if codePassed(code, Jump) {
		res = append(res, "jump")
	}

	if codePassed(code, Reverse) {
		res2 := *new([]string)
		for i := len(res) - 1; i >= 0; i-- {
			res2 = append(res2, res[i])
		}
		res = res2
	}
	return res
}

func codePassed(code int, flag int) bool {
	return code&flag != 0
}

package luhn

import (
        "strings"
        "unicode"
)

const testVersion = 1

func Valid(input string) bool {
        input = strings.Replace(input, " ", "", -1)

        for _, r := range input {
                if !unicode.IsDigit(r) {
                        return false
                }
        }
        if len(input) < 2 {
                return false
        }

        return validChecksum(input)
}

func validChecksum(input string) bool {
        var sum int
        for i, r := range input {
                d := int(r - '0')
                if i%2 == 1 {
                        d *= 2
                        if d > 9 {
                                d -= 9
                        }
                }
                sum += d
        }

        return sum%10 == 0
}

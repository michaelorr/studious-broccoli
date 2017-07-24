package palindrome

import "errors"
import "strconv"

const testVersion = 1

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	all := allPalindromes(fmin, fmax)
	if len(all) == 0 {
		return pmin, pmax, errors.New("no palindromes...")
	}

	// define pmin sentinal as a Product struct with intmax, could also use arbitrarily large value
	pmin = Product{int(^uint(0) >> 1), nil}
	for _, palindrome := range all {
		if palindrome.Product <= pmin.Product {
			pmin = palindrome
		}
		if palindrome.Product > pmax.Product {
			pmax = palindrome
		}
	}
	return pmin, pmax, nil
}

func allPalindromes(fmin, fmax int) []Product {
	palindromeMap := allPalindromesAsMap(fmin, fmax)

	var products []Product
	for k, v := range palindromeMap {
		products = append(products, Product{k, v})
	}
	return products
}

func isPalindrome(x int) bool {
	str1 := strconv.Itoa(x)
	return str1 == reverse(str1)
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func allPalindromesAsMap(fmin, fmax int) map[int][][2]int {
	productMap := make(map[int][][2]int)
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if isPalindrome(product) {
				productMap[product] = append(productMap[product], [2]int{i, j})
			}
		}
	}
	return productMap
}

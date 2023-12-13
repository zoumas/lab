package solution

import (
	"sort"
	"unicode"
)

func Solution1(strArr []string) []string {
	if len(strArr) < 1 {
		return strArr
	}

	sort.Slice(strArr, func(i, j int) bool {
		a := strArr[i]
		b := strArr[j]

		la := len(a) % 3
		lb := len(b) % 3

		switch {
		case la < lb:
			return true
		case la > lb:
			return false
		}

		l := len(a)
		if len(b) > l {
			l = len(b)
		}

		for i := 0; i < l; i++ {
			ai := unicode.ToLower(rune(a[i]))
			bi := unicode.ToLower(rune(b[i]))
			switch {
			case ai < bi:
				return true
			case ai > bi:
				return false
			}
		}
		return false
	})
	return strArr
}

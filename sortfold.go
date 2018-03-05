/*
Package sortfold enables sorting string slices in increasing order using
a case-insensitive comparison.
*/
package sortfold

import (
	"sort"

	"unicode"
	"unicode/utf8"
)

// StringSlice attaches the methods of Interface to []string,
// sorting in increasing order using a case insensitive comparison.
type StringSlice []string

func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int) bool {
	return CompareFold(p[i], p[j]) < 0
}
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { sort.Sort(p) }

// Strings sorts a slice of strings in increasing order using a case
// insensitive comparison.
func Strings(a []string) { sort.Sort(StringSlice(a)) }

// StringsAreSorted ¶ tests whether a slice of strings is sorted in increasing
// order using a case insensitive comparison.
func StringsAreSorted(a []string) bool {
	return sort.IsSorted(StringSlice(a))
}

// CompareFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding. A return value of 0 means s==t,
// <0 means s<t, and >0 means s>t.
//
// This function is a modification of the golang strings.EqualFold function
// from strings/strings.go at https://goo.gl/yiMur6.
func CompareFold(s, t string) int {
	for s != "" && t != "" {

		// Extract first rune from each string.
		var sr, tr rune
		if s[0] < utf8.RuneSelf {
			sr, s = rune(s[0]), s[1:]
		} else {
			r, size := utf8.DecodeRuneInString(s)
			sr, s = r, s[size:]
		}
		if t[0] < utf8.RuneSelf {
			tr, t = rune(t[0]), t[1:]
		} else {
			r, size := utf8.DecodeRuneInString(t)
			tr, t = r, t[size:]
		}

		// If they match, keep going; if not, return false.
		//log.Printf("sr=%[1]v,%[1]c tr=%[2]v,%[2]c", sr, tr)

		// Easy case.
		if sr == tr {
			continue
		}

		// Make sr < tr to simplify what follows.
		result := 1
		if tr < sr {
			tr, sr = sr, tr
			result = -result
		}

		// Fast check for ASCII.
		if sr < utf8.RuneSelf && tr < utf8.RuneSelf {
			sr := sr
			tr := tr
			if 'A' <= sr && sr <= 'Z' {
				sr = sr + 'a' - 'A'
			}
			if 'A' <= tr && tr <= 'Z' {
				tr = tr + 'a' - 'A'
			}
			//log.Printf("fastcheck: sr=%[1]v,%[1]c tr=%[2]v,%[2]c", sr, tr)
			if sr == tr {
				continue
			}
			if sr < tr {
				//log.Printf("sr < tr = %d", -1)
				return -result
			}
			if sr > tr {
				//log.Printf("sr > tr = %d", 1)
				return result
			}
		}

		// General case. SimpleFold(x) returns the next equivalent rune > x
		// or wraps around to smaller values.
		r := unicode.SimpleFold(sr)
		for r != sr && r < tr {
			r = unicode.SimpleFold(r)
			//log.Printf("+simplefold: r=%[1]v,%[1]c", r)
		}
		//log.Printf("simplefold: r=%[1]v,%[1]c tr=%[2]v,%[2]c", r, tr)
		if r == tr {
			continue
		}
		if r < tr {
			return -result
		}
		if r > tr {
			return result
		}
	}

	// One string is empty. Are both?
	if s == "" && t != "" {
		return -1
	}
	if s != "" && t == "" {
		return 1
	}
	return 0
}

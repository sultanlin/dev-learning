package iteration

import "strings"

func Repeat(c string, n int) string {
	// r := ""
	// for i := 0; i < n; i++ {
	// 	r += c
	// }
	// return r
	return strings.Repeat(c, n)
}

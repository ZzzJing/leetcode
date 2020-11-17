package stack

import "bytes"

func isValid(s string) bool {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		st = append(st, s[i])
		if len(st) >= 3 && bytes.Equal(st[len(st)-3:], []byte("abc")) {
			st = st[:len(st)-3]
		}
	}
	return len(st) == 0
}

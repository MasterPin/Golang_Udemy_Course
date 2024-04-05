package tests

import "testing"

func TestROT13__success(t *testing.T) {
	result := rot13([]byte("abc"))
	println(result)
}

func rot13(bs []byte) string {
	var r13 = make([]byte, len(bs))
	for i, b := range bs {
		if b <= 109 {
			r13[i] = b + 13
		} else {
			r13[i] = b - 13
		}
	}
	return string(r13)
}

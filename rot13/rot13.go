package rot13

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

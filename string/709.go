package string

func toLowerCase(str string) string {
	bl := []byte(str)
	for k, v := range bl {
		if v >='A' && v <= 'Z' {
			bl[k] = v + 32
		}
	}

	return string(bl)
}

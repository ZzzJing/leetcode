package string

func reverseString(s []byte)  {
	l := len(s) - 1
	if l != -1 {
		mid := l /2
		for i:=0;i<= mid;i++ {
			tmp := s[i]
			s[i] = s[l-i]
			s[l-i] = tmp
		}
	}
}

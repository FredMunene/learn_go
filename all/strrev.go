package piscine

func StrRev(s string) string { // convert string to a slice of bytes
	Astring := []byte(s)
	length := len(Astring)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 { // iterate over slice and swap
		Astring[i], Astring[j] = Astring[j], Astring[i]
	}
	return string(Astring)
}

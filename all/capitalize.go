package piscine

func IsAlphanumeric(a rune) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

func Capitalize(s string) string {
	arr := []rune(s)
	toCapitalize := true
	for i := 0; i < len(arr); i++ {
		if IsAlphanumeric(arr[i]) && toCapitalize {
			if arr[i] >= 'a' && arr[i] <= 'z' {
				arr[i] = arr[i] - 'a' + 'A'
			}
			toCapitalize = false
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] - 'A' + 'a'
		} else if !IsAlphanumeric(arr[i]) {
			toCapitalize = true
		}
	}
	return string(arr)
}

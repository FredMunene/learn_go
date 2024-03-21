package piscine

func IsEven(nb int) bool { // checks if nb is even or odd
	return nb%2 == 0
}

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	if start == 1 { // to catch infinite loop
		return 0
	}
	step := 1 // initialise step
	if IsEven(start) {
		start = start / 2
	} else {
		start = start*3 + 1
	}
	return step + CollatzCountdown(start) // >> step + step + reccursive
}

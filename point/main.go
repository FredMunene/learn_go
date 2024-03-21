package main

import "github.com/01-edu/z01"

var a string = "x = 42, y = 21"

/*type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)

	// for _,value := range points{
	// }
	z01.PrintRune('x')
*/
// fmt.Printf("x = %d, y = %d\n",points.x, points.y)
func main() {
	for _, value := range a {
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')
}

package piscine

func AppendRange(min, max int) []int {
	var grid []int
	if min > max {
		return nil
	}
	for i := 0; i < max-min; i++ {
		grid = append(grid, min+i)
	}
	return grid
}

package main

func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, move := range moves {
		switch move {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return false
}

package main

import (
	"fmt"
)

func IsPathCrossed(steps []string) bool {
	// Starting position
	x, y := 0, 0

	// Use a map to track visited coordinates as "x,y"
	visited := map[string]bool{
		"0,0": true, // Starting point is visited
	}

	// Direction map to update coordinates
	dirs := map[string][2]int{
		"north": {0, 1},
		"south": {0, -1},
		"east":  {1, 0},
		"west":  {-1, 0},
	}

	for _, step := range steps {
		if delta, ok := dirs[step]; ok {
			x += delta[0]
			y += delta[1]

			pos := fmt.Sprintf("%d,%d", x, y)

			if visited[pos] {
				// Already been here before — path crossed
				return true
			}
			visited[pos] = true
		}
	}

	return false
}

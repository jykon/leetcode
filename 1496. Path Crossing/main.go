package main

import "fmt"

func isPathCrossing(path string) bool {
	coordinateSet := make(map[string]struct{})
	x, y := 0, 0
	coordinateSet["0,0"] = struct{}{}

	for _, ch := range path {
		switch ch {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		coordinate := fmt.Sprintf("%d,%d", x, y)

		if _, exists := coordinateSet[coordinate]; exists {
			return true
		}

		coordinateSet[coordinate] = struct{}{}
	}

	return false
}
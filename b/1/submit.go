package main

import (
	"errors"
)

// わからなかったのでChatGPTによるsolutionコードを参考にしました
func CalculateOptimalPath(grid [][]int, startX, startY, endX, endY int) (totalCost int, err error) {
	if err = Validate(grid, startX, startY, endX, endY); err != nil {
		return 0, err
	}

	gridLen := len(grid)
	distancesList := make([][]int, gridLen)
	for i := range distancesList {
		distancesList[i] = make([]int, gridLen)
		for j := range distancesList[i] {
			distancesList[i][j] = int(^uint(0) >> 1)
		}
	}
	distancesList[startX][startY] = 0

	visitedList := make([][]bool, gridLen)
	for i := range visitedList {
		visitedList[i] = make([]bool, gridLen)
	}
	dxList := []int{0, 1, 0, -1}
	dyList := []int{-1, 0, 1, 0}
	for {
		isFound := false
		var sx, sy int
		for i := range grid {
			for j := range grid[i] {
				if !visitedList[i][j] && (!isFound || distancesList[i][j] < distancesList[sx][sy]) {
					sx, sy = i, j
					isFound = true
				}
			}
		}
		if !isFound {
			break
		}
		visitedList[sx][sy] = true

		for i := 0; i < 4; i++ {
			nx, ny := sx+dxList[i], sy+dyList[i]
			if IsInRange(nx, 0, gridLen-1) &&
				IsInRange(ny, 0, gridLen-1) &&
				!visitedList[nx][ny] &&
				distancesList[sx][sy]+grid[nx][ny] < distancesList[nx][ny] {
				distancesList[nx][ny] = distancesList[sx][sy] + grid[nx][ny]
			}
		}
	}

	return distancesList[endX][endY], nil
}

func Validate(grid [][]int, startX, startY, endX, endY int) error {
	gridLen := len(grid)
	for _, row := range grid {
		if len(row) != gridLen {
			return errors.New("invalid grid length")
		}
		for _, cell := range row {
			if cell <= 0 {
				return errors.New("invalid grid value")
			}
		}
	}

	if !IsInRange(startX, 0, gridLen-1) ||
		!IsInRange(startY, 0, gridLen-1) ||
		!IsInRange(endX, 0, gridLen-1) ||
		!IsInRange(endY, 0, gridLen-1) {
		return errors.New("invalid start or end coordinates")
	}

	return nil
}

func IsInRange(num, start, end int) bool {
	return num >= start && num <= end
}

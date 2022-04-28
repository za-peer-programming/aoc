package day4

import (
	"strconv"
	"strings"
)

type Board [][]Point

type Point struct {
	i      int
	marked bool
}

func TextToBoard(s string) (Board, error) {
	x := strings.Split(s, "\n")
	newBoard := make([][]Point, 5)
	for i := 0; i < len(x); i++ {
		splitLine := strings.Split(x[i], " ")
		newBoard[i] = make([]Point, 5)
		col := 0
		for _, point := range splitLine {
			if point == "" {
				continue
			}
			s, err := strconv.Atoi(point)
			if err != nil {
				return nil, err
			}
			newBoard[i][col] = Point{i: s}
			col++
		}
	}
	return newBoard, nil
}

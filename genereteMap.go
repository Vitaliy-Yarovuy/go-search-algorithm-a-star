package main

import (
	"math/rand"
	"time"
)

func generateMap(size Point) (matrix []Point, noize []Point, start Point, finish Point) {
	var length int = size.X * size.Y
	var noizeLength = length / 10
	matrix = make([]Point, 0, length)
	noize = make([]Point, noizeLength, noizeLength)

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			matrix = append(matrix, Point{i, j})
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	var perms = rand.Perm(length - 1)

	for i, value := range perms[:noizeLength] {
		noize[i] = matrix[value]
	}

	start = matrix[perms[noizeLength+1]]
	finish = matrix[perms[noizeLength+2]]

	return
}

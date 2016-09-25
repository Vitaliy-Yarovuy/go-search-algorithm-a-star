package main

// input
var matrix []Point
var noize []Point

var start Point
var finish Point

// output
var result []Point

func main() {
	var size Point = Point{40, 20}
	matrix, noize, start, finish = generateMap(size)
	path := AStar(matrix, noize, start, finish)
	renderMap(size, noize, start, finish, path)
}

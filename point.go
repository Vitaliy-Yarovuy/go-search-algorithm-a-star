package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type PointData struct {
	Point
	LengthFromStart int
	LengthToFinish  int
	Prev            *PointData
}

type Points []Point
type PointDatas []PointData

func (a PointDatas) Len() int           { return len(a) }
func (a PointDatas) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PointDatas) Less(i, j int) bool { return a[i].WayLength() < a[j].WayLength() }

func PointDataSearch(ps []PointData, p PointData) func(int) bool {
	return func(i int) bool {
		return p.WayLength() < ps[i].WayLength()
	}
}

func (that Point) String() string {
	return fmt.Sprintf("%d|%d", that.X, that.Y)
}

func (that PointData) WayLength() int {
	return that.LengthFromStart + that.LengthToFinish
}

func (that PointData) String() string {
	return fmt.Sprintf("%d|%d#%d", that.X, that.Y, that.LengthToFinish)
}

func (that Points) String() string {
	var result = "-"
	for _, value := range that {
		result += fmt.Sprintf("[%d|%d].", value.X, value.Y)
	}
	return result
}

func (that PointDatas) String() string {
	var result = "-"
	for _, value := range that {
		result += fmt.Sprintf("[%d|%d#%d],", value.X, value.Y, value.LengthToFinish)
	}
	return result
}

func (that Point) Move(diff Point) Point {
	return Point{that.X + diff.X, that.Y + diff.Y}
}

func (that Point) AttachData(lengthFStart int, finish Point, prev *PointData) PointData {
	return PointData{that, lengthFStart, that.expectLength(finish), prev}
}

func (that *PointData) GetPath() []Point {
	var result = make([]Point, 0)
	for that != nil {
		result = append([]Point{that.Point}, result...)
		that = that.Prev
	}
	return result
}

func (that Points) ToMap() (result map[string]int) {
	result = make(map[string]int, 0)
	for _, value := range that {
		result[value.String()]++
	}
	return
}

func (that Points) Contains(point Point) bool {
	for _, value := range that {
		if value == point {
			return true
		}
	}
	return false
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}

func (from Point) expectLength(to Point) int {
	return abs(from.X-to.X) + abs(from.Y-to.Y)
}

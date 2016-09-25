package main

import (
	"sort"
)

type EngueuePoint struct {
	ps []PointData
}

func (that *EngueuePoint) Add(p PointData) {
	var i int = sort.Search(len(that.ps), PointDataSearch(that.ps, p))
	that.ps = append(that.ps[:i], append([]PointData{p}, that.ps[i:]...)...)
}

func (that *EngueuePoint) Get() (PointData, bool) {
	if len(that.ps) == 0 {
		return PointData{}, false
	}
	var p = that.ps[0]
	that.ps = that.ps[1:]
	return p, true
}

func (that *EngueuePoint) Sort() {
	sort.Sort(PointDatas(that.ps))
}

func (that *EngueuePoint) ContainsPoint(point Point) (bool, PointData) {
	for _, value := range that.ps {
		if value.Point == point {
			return true, value
		}
	}
	return false, PointData{}
}

func (that *EngueuePoint) IsEmpty() bool {
	return len(that.ps) == 0
}

func (that *EngueuePoint) Len() int {
	return len(that.ps)
}

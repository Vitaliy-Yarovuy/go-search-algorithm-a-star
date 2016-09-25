package main

var moveList = Points{Point{0,-1},Point{-1,0},Point{1,0},Point{0,1}}
// var moveList = Points{Point{0, -1}, Point{-1, 0}, Point{1, 0}, Point{0, 1}, Point{1, -1}, Point{-1, 1}, Point{1, 1}, Point{-1, -1}}

func AStar(matrix []Point, noize []Point, start Point, finish Point) []Point {
	var openlist EngueuePoint = EngueuePoint{}
	var closedlist Points = Points{}
	var pData PointData

	openlist.Add(start.AttachData(0, finish, nil))

	for {
		pData, _ = openlist.Get()
		if pData.Point == finish {
			return pData.GetPath()
		}
		openlist = expandNode(pData, matrix, noize, openlist, closedlist)
		closedlist = append(closedlist, pData.Point)

		if openlist.IsEmpty() {
			break
		}
	}
	return []Point{}
}

func expandNode(pData PointData, matrix Points, noize Points, openlist EngueuePoint, closedlist Points) EngueuePoint {
	for _, moveDiff := range moveList {
		expPoint := pData.Point.Move(moveDiff)
		if !matrix.Contains(expPoint) || noize.Contains(expPoint) {
			continue
		}

		if closedlist.Contains(expPoint) {
			continue
		}
		lenghtFromStart := pData.LengthFromStart + 1
		inOpenList, expPointData := openlist.ContainsPoint(expPoint)

		if inOpenList && lenghtFromStart >= expPointData.LengthFromStart {
			continue
		}

		if inOpenList {
			expPointData.Prev = &pData
			expPointData.LengthFromStart = lenghtFromStart
			openlist.Sort()
		} else {
			openlist.Add(expPoint.AttachData(lenghtFromStart, finish, &pData))
		}
	}

	return openlist
}

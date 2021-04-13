package models

// Point2D : a struct that stores X and Y coordinate values
type Point2D struct {
	xValue float64
	yValue float64
}

// Graph2D : list of all 2D points
type Graph2D struct {
	points []Point2D
	//numPoints int
}
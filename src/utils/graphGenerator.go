package utils

import (
	"github.com/golang/geo/r3"
	"math/rand"
	"models"
)

func New(numPoints int) models.Graph2D {
	return models.Graph2D{Points: make([]models.Point2D, numPoints)}
}

func New3(numPoints int) models.Graph3D {
	return models.Graph3D{Points: make([]r3.Vector, numPoints)}
}

func GenerateTriangularGraph3D(numPoints int) models.Graph3D {
	var g = New3(numPoints)

	for i := 0; i < numPoints; i++ {
		xVal := rand.Float64() * 1000
		yVal := rand.Float64() * 1000
		zVal := rand.Float64() * 1000
		if xVal < yVal {
			g.Points[i].X = xVal
			g.Points[i].Y = yVal
			g.Points[i].Z = zVal
		} else {
			numPoints--
		}
	}

	return g
}

func GenerateTriangularGraph2D(numPoints int) models.Graph2D {
	var g = New(numPoints)

	for i := 0; i < numPoints; i++ {
		xVal := rand.Float64() * 1000
		yVal := rand.Float64() * 1000
		if xVal < yVal {
			g.Points[i].XValue = xVal
			g.Points[i].YValue = yVal
		} else {
			numPoints--
		}
	}

	return g
}

// GenerateGraph3D : creates a graph with 'numPoints' points that
// are randomly assigned values from [0, 1000.0) float 64
func GenerateGraph3D(numPoints int) models.Graph3D {
	var g = New3(numPoints)

	for i := 0; i < numPoints; i++ {
		g.Points[i].X = rand.Float64() * 1000
		g.Points[i].Y = rand.Float64() * 1000
		g.Points[i].Z = rand.Float64() * 1000
	}

	return g
}


// GenerateGraph2D : creates a graph with 'numPoints' points that
// are randomly assigned values from [0, 1000.0) float 64
func GenerateGraph2D(numPoints int) models.Graph2D {
	var g = New(numPoints)

	for i := 0; i < numPoints; i++ {
		g.Points[i].XValue = rand.Float64() * 1000
		g.Points[i].YValue = rand.Float64() * 1000
	}

	return g
}

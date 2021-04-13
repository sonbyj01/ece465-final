package utils

import (
	"math/rand"
	"models"
)

func New(numPoints int) models.Graph2D {
	return models.Graph2D{Points: make([]models.Point2D, numPoints)}
}

// GenerateGraph2D : creates a graph with 'numPoints' points that
// are randomly assigned values from [0, 10.0) float 64
func GenerateGraph2D(numPoints int) models.Graph2D {
	var g = New(numPoints)

	for i := 0; i < numPoints; i++ {
		g.Points[i].XValue = rand.Float64() * 10
		g.Points[i].YValue = rand.Float64() * 10
	}

	return g
}
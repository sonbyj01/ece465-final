package utils

import(
	"math/rand"
	"models"
)

func New(numVertices int) models.Graph2D {
	return models.Graph2D{make([]models.Point2D, numVertices)}
}

func GenerateGraph2D(numVertices int) models.Graph2D {
	var g = New(numVertices)
	return g
}
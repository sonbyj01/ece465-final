package algorithm

import (
	"bufio"
	"fmt"
	"models"
	"os"
	"path/filepath"
)

func Select(S *models.Graph2D) (*models.Point2D, *models.Point2D) {
	var p,q models.Point2D
	for index, point := range S.Points {
		if index == 0 {
			p = point
			continue
		}

		if index == 1 {
			q = point
			continue
		}

		// maximum y-value, then maximum x-value if y values are same
		if p.YValue < point.YValue {
			p = point
		} else if p.YValue == point.YValue {
			if p.XValue < point.XValue {
				p = point
			}
		}
		// minimum y-value, then minimum x-value if y values are same
		if q.YValue > point.YValue {
			q = point
		} else if q.YValue == point.YValue {
			if q.XValue > point.XValue {
				q = point
			}
		}
	}
	return &p, &q
}

func Split(p *models.Point2D, q *models.Point2D, S *models.Graph2D) (*models.Graph2D, *models.Graph2D) {
	var R, L models.Graph2D

	for _, point := range S.Points {
		d := models.RightSide(p, q, &point)
		if d == 0 {
			R.Points = append(R.Points, point)
		} else if d == 1 {
			L.Points = append(L.Points, point)
		}
	}
	return &R, &L
}

func QuickHull(p *models.Point2D, q *models.Point2D, S *models.Graph2D, convex *models.Graph2D) {
	if !S.IsEmpty() {
		//fmt.Println("S: ", *S)
		r := FarthestPoint(p, q, S)
		U, L := PruneAndSplit(p, q, r, S)
		//fmt.Println("p: ", *p)
		//fmt.Println("q: ", *q)
		QuickHull(p, r, U, convex)
		convex.Points = append(convex.Points, *r)
		fmt.Println("r: ", *r)
		QuickHull(r, q, L, convex)
	}
}

func FarthestPoint(p *models.Point2D, q *models.Point2D, S *models.Graph2D) *models.Point2D {
	var line models.Line
	line.Create(p, q)

	var farPoint models.Point2D
	var farDistance float64

	for _, point := range S.Points {
		distance := line.PerpendicularDistance(&point)
		if farDistance < distance {
			farPoint = point
			farDistance = distance
		}
	}
	return &farPoint
}

func PruneAndSplit(p *models.Point2D, q *models.Point2D, r *models.Point2D, S *models.Graph2D) (*models.Graph2D, *models.Graph2D) {
	U, _ := Split(p, r, S)
	L, _ := Split(r, q, S)
	return U, L
}

func RunQuickHull(input string, output string) (*models.Graph2D, *models.Graph2D) {
	inputAbs, _ := filepath.Abs("./tmp/" + input)
	inputFile, _ := os.Open(inputAbs)
	g, _ := models.Import(inputFile)
	inputFile.Close()

	var convex *models.Graph2D
	convex = new(models.Graph2D)
	p, q := Select(g)
	convex.Points = append(convex.Points, *p)
	convex.Points = append(convex.Points, *q)
	R, L := Split(p, q, g)
	QuickHull(p, q, R, convex)
	QuickHull(q, p, L, convex)

	outputAbs, _ := filepath.Abs("./tmp/" + output)
	outputFile, _ := os.Create(outputAbs)
	writer := bufio.NewWriter(outputFile)
	convex.Export(writer)
	writer.Flush()
	outputFile.Close()

	return g, convex
}
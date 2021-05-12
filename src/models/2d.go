package models

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// Point2D : a struct that stores X and Y coordinate values
type Point2D struct {
	XValue float64
	YValue float64
}

// Graph2D : list of all 2D points
type Graph2D struct {
	Points []Point2D
	//numPoints int
}

// Line : linear equation Ax + By + C = 0
type Line struct {
	A float64
	B float64
	C float64
}

// FloatToString : converts float value to string
// https://stackoverflow.com/questions/19101419/formatfloat-convert-float-number-to-string
func FloatToString(floatVal float64) string {
	return strconv.FormatFloat(floatVal, 'f', 6, 64)
}

func (g *Graph2D) Print() {
	fmt.Println(g.Points[0])
}

func (g *Graph2D) IsEmpty() bool {
	if g.Points == nil {
		return true
	}
	return false
}

func (p *Point2D) Print() {
	fmt.Println(p)
}

// Create : from points p and q, solves for A, B, C of standard linear form
// mx - y - mx_1 + y_1 = 0
func (l *Line) Create(p *Point2D, q *Point2D) {
	m := (p.YValue - q.YValue) / (p.XValue - q.XValue)
	l.A = m
	l.B = -1
	l.C = -m * p.XValue + p.YValue
}

// PerpendicularDistance : calculates the perpendicular distance between a point and line
// https://www.intmath.com/plane-analytic-geometry/perpendicular-distance-point-line.php
func (l *Line) PerpendicularDistance(p *Point2D) float64 {
	return math.Abs(l.A * p.XValue + l.B * p.YValue + l.C) / math.Sqrt(l.A * l.A + l.B * l.B)
}

// RightSide : determines whether a point r is on the left or right side of a line between points p and q
// Return 1 = on the right side
//		  0 = on the left side
//		 -1 = on the line
// https://math.stackexchange.com/questions/274712/calculate-on-which-side-of-a-straight-line-is-a-given-point-located
// https://stackoverflow.com/questions/1560492/how-to-tell-whether-a-point-is-to-the-right-or-left-side-of-a-line#3461533
func RightSide(p *Point2D, q *Point2D, r *Point2D) int {
	d := (r.XValue - p.XValue) * (q.YValue - p.YValue) - (r.YValue - p.YValue) * (q.XValue - p.XValue)
	if 0 > d {
		return 0
	} else if d == 0 {
		return -1
	}
	return 1
}

func Import(reader io.Reader) (*Graph2D, error) {
	scanner := bufio.NewScanner(reader)

	var numPoints int
	scanner.Scan()
	numPoints, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, err
	}

	g := Graph2D{make([]Point2D, numPoints)}

	for i := 0; scanner.Scan(); i++ {
		v := &g.Points[i]

		coordinates := strings.Split(scanner.Text(), ",")

		valueX, err := strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return nil, err
		}
		v.XValue = valueX

		valueY, err := strconv.ParseFloat(coordinates[1], 64)
		if err != nil {
			return nil, err
		}
		v.YValue = valueY
	}
	return &g, nil
}

func (g *Graph2D) Export(writer io.Writer) error {
	// write number of points
	_, err := io.WriteString(writer, strconv.Itoa(len(g.Points)) + "\n")
	if err != nil {
		return err
	}

	// write each point
	for _, point := range g.Points {
		s := FloatToString(point.XValue)
		s += ","
		s += FloatToString(point.YValue)
		s += "\n"

		_, err = io.WriteString(writer, s)
		if err != nil {
			return err
		}
	}
	return nil
}
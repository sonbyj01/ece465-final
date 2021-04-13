package models

import (
	"io"
	"strconv"
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

// FloatToString : converts float value to string
// https://stackoverflow.com/questions/19101419/formatfloat-convert-float-number-to-string
func FloatToString(floatVal float64) string {
	return strconv.FormatFloat(floatVal, 'f', 6, 64)
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
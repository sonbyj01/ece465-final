package models

import (
	"bufio"
	"fmt"
	"io"
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

// FloatToString : converts float value to string
// https://stackoverflow.com/questions/19101419/formatfloat-convert-float-number-to-string
func FloatToString(floatVal float64) string {
	return strconv.FormatFloat(floatVal, 'f', 6, 64)
}

func (g *Graph2D) Print() {
	fmt.Println(g.Points[0])
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
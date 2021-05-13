package models

import (
	"bufio"
	"github.com/golang/geo/r3"
	"io"
	"strconv"
	"strings"
)

// Graph3D : list of all 3D points
type Graph3D struct {
	Points []r3.Vector
	//numPoints int
}

func Import3(reader io.Reader) (*Graph3D, error) {
	scanner := bufio.NewScanner(reader)

	var numPoints int
	scanner.Scan()
	numPoints, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, err
	}

	g := Graph3D{make([]r3.Vector, numPoints)}

	for i := 0; scanner.Scan(); i++ {
		v := &g.Points[i]

		coordinates := strings.Split(scanner.Text(), ",")

		valueX, err := strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return nil, err
		}
		v.X = valueX

		valueY, err := strconv.ParseFloat(coordinates[1], 64)
		if err != nil {
			return nil, err
		}
		v.Y = valueY

		valueZ, err := strconv.ParseFloat(coordinates[2], 64)
		if err != nil {
			return nil, err
		}
		v.Z = valueZ
	}
	return &g, nil
}

func (g *Graph3D) Export3(writer io.Writer) error {
	// write number of points
	_, err := io.WriteString(writer, strconv.Itoa(len(g.Points)) + "\n")
	if err != nil {
		return err
	}

	// write each point
	for _, point := range g.Points {
		s := FloatToString(point.X)
		s += ","
		s += FloatToString(point.Y)
		s += ","
		s += FloatToString(point.Z)
		s += "\n"

		_, err = io.WriteString(writer, s)
		if err != nil {
			return err
		}
	}
	return nil
}
package algorithm

import (
	"bufio"
	"fmt"
	"github.com/markus-wa/quickhull-go"
	"models"
	"os"
	"path/filepath"
)

func RunQuickHull3(input string, output string) *models.Graph3D {
	inputAbs, _ := filepath.Abs("./tmp/" + input)
	inputFile, _ := os.Open(inputAbs)

	g, err1 := models.Import3(inputFile)
	if err1 != nil {
		fmt.Errorf("ha")
	}
	inputFile.Close()

	convex := new(models.Graph3D)
	convex.Points = new(quickhull.QuickHull).ConvexHull(g.Points, true, false, 0).Vertices

	//fmt.Println(hull.Vertices) // does not contain (4,4,1)
	//fmt.Println(hull.Triangles()) // triangles that make up the convex hull - [][3]r3.Vector, where each vector is a corner of the triangle

	outputAbs, _ := filepath.Abs("./tmp/" + output)
	outputFile, _ := os.Create(outputAbs)
	writer := bufio.NewWriter(outputFile)
	convex.Export3(writer)
	writer.Flush()
	outputFile.Close()

	return convex
}

package main

import (
	"bufio"
	"fmt"
	quickhull "github.com/markus-wa/quickhull-go"
	"models"
	"os"
	"path/filepath"
)

func main() {
	abs, _ := filepath.Abs("./sample/graph3d-1000.txt")
	file, err := os.Open(abs)

	if err != nil {
		fmt.Print(err)
	}

	g, err1 := models.Import3(file)
	if err1 != nil {
		fmt.Errorf("ha")
	}
	file.Close()

	convex := new(models.Graph3D)
	convex.Points = new(quickhull.QuickHull).ConvexHull(g.Points, true, false, 0).Vertices

	//fmt.Println(hull.Vertices) // does not contain (4,4,1)
	//fmt.Println(hull.Triangles()) // triangles that make up the convex hull - [][3]r3.Vector, where each vector is a corner of the triangle

	file1, _ := os.Create("./sample/graph3d-1000-convex.txt")
	writer := bufio.NewWriter(file1)
	convex.Export3(writer)
	writer.Flush()
	file1.Close()

	//abs, _ := filepath.Abs("./sample/graph2d-1000.txt")
	//file, err := os.Open(abs)
	//
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//g, err1 := models.Import(file)
	//if err1 != nil {
	//	fmt.Errorf("ha")
	//}
	//file.Close()
	//
	//var convex *models.Graph2D
	//convex = new(models.Graph2D)
	//
	//p, q := algorithm.Select(g)
	//convex.Points = append(convex.Points, *p)
	//convex.Points = append(convex.Points, *q)
	//
	//R, L := algorithm.Split(p, q, g)
	//
	//algorithm.QuickHull(p, q, R, convex)
	//
	//algorithm.QuickHull(q, p, L, convex)
	//
	//file1, _ := os.Create("./sample/graph2d-1000-convex.txt")
	//writer := bufio.NewWriter(file1)
	//convex.Export(writer)
	//writer.Flush()
	//file1.Close()
}
package main

import (
	"algorithm"
	"bufio"
	"fmt"
	"models"
	"os"
	"path/filepath"
)

func main() {
	abs, _ := filepath.Abs("./sample/graph2d-1000.txt")
	file, err := os.Open(abs)

	if err != nil {
		fmt.Print(err)
	}

	g, err1 := models.Import(file)
	if err1 != nil {
		fmt.Errorf("ha")
	}
	file.Close()

	var convex *models.Graph2D
	convex = new(models.Graph2D)

	p, q := algorithm.Select(g)
	convex.Points = append(convex.Points, *p)
	convex.Points = append(convex.Points, *q)
	fmt.Println("p: ", *p)
	fmt.Println("q: ", *q)

	R, L := algorithm.Split(p, q, g)
	//fmt.Println("R: ", *R)
	//fmt.Println("L: ", *L)

	//fmt.Println("p: ", *p)

	algorithm.QuickHull(p, q, R, convex)

	algorithm.QuickHull(q, p, L, convex)

	//fmt.Println("q: ", *q)
	//test := algorithm.FarthestPoint(p, q, R)
	//test.Print()
	//
	//R1, _ := algorithm.Split(p, test, R)
	//test1 := algorithm.FarthestPoint(p, test, R1)
	//test1.Print()

	//fmt.Println(L)

	//p.Print()
	//
	//algorithm.QuickHull(p, q, R)
	//
	//q.Print()
	//
	//algorithm.QuickHull(q, p, L)

	//g.Print()
	//
	file1, _ := os.Create("./sample/graph2d-1000-convex.txt")
	writer := bufio.NewWriter(file1)
	convex.Export(writer)
	writer.Flush()
	file1.Close()
}
package main

import (
	"bufio"
	"os"
	"utils"
)

func main() {
	g := utils.GenerateGraph3D(1000)
	file, _ := os.Create("./sample/graph3d-1000.txt")
	writer := bufio.NewWriter(file)
	g.Export3(writer)
	writer.Flush()

	g = utils.GenerateTriangularGraph3D(1000)
	file, _ = os.Create("./sample/graph3dTriangle-1000.txt")
	writer = bufio.NewWriter(file)
	g.Export3(writer)
	writer.Flush()
}
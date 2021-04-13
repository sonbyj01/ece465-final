package main

import (
	"bufio"
	"os"
	"utils"
)

func main() {
	g := utils.GenerateGraph2D(100)
	file, _ := os.Create("./sample/graph2d-100.txt")
	writer := bufio.NewWriter(file)
	g.Export(writer)
	writer.Flush()
}
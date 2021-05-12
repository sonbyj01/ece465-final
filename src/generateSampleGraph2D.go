package main

import (
	"bufio"
	"os"
	"utils"
)

func main() {
	g := utils.GenerateGraph2D(1000)
	file, _ := os.Create("./sample/graph2d-1000.txt")
	writer := bufio.NewWriter(file)
	g.Export(writer)
	writer.Flush()
}
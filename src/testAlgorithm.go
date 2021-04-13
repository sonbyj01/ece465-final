package main

import (
	"bufio"
	"fmt"
	"models"
	"os"
	"path/filepath"
)

func main() {
	abs, _ := filepath.Abs("./sample/graph2d-100.txt")
	file, err := os.Open(abs)

	if err != nil {
		fmt.Print(err)
	}

	g, err1 := models.Import(file)
	if err1 != nil {
		fmt.Errorf("ha")
	}
	file.Close()
	g.Print()

	file1, _ := os.Create("./sample/graph2d-100-copy.txt")
	writer := bufio.NewWriter(file1)
	g.Export(writer)
	writer.Flush()
	file1.Close()
}
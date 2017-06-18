package main

import (
	"bufio"
	"os"

	"github.com/buger/goterm"
)

// FileMap monta linea a linea de un mapa de int a strings dado un fichero.
func FileMap(filePath string) map[int]string {
	file, _ := os.Open(filePath)
	defer file.Close()

	words := make(map[int]string)
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		words[i] = scanner.Text()
	}

	return words
}

// ResetPos limpia la consola y situa el cursor al inicio.
func ResetPos() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Flush()
}

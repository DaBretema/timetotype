package back

import (
	"bufio"
	"os"
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

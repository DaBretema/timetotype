// Verano de 2016.
// Daniel Camba Lamas.
// JUEGO PROPIO: TimeToType.

/* TO DO...
- Añadir GUI.
- Elegir dificultad al inicio.
- Particionar 'words' por dificultad.
- Detener programa al ganar o perder.
- Mejorar 'formula' inc/decr del tiempo.
*/

package main

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/dcabalas/TimeToType/models"
	"math/rand"
	"os"
	"time"
)

func main() {

	t := time.Now()
	p := player.NewPlayer(2*time.Minute, 5, 0, 0, 0)
	words := TextMap()

	rand.Seed(t.Unix())

	go p.Time()

	for p.Flag == 0 {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
		p.UI()
		p.Play(words)
	}

	fmt.Println("\n\nTÚ PIERDES D:")
}

//Montaje del map de palabras.
func TextMap() map[int]string {
	file, _ := os.Open("words.txt")
	defer file.Close()

	words := make(map[int]string)
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		words[i] = scanner.Text()
	}

	return words
}

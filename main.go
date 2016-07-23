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
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/dcabalas/TimeToType/back"
	"math/rand"
	"os"
	"time"
)

func main() {

	t := time.Now()
	p := back.NewPlayer(2*time.Minute, 5, 0, 0, 0)
	words := back.FileMap("resources/words.txt")

	rand.Seed(t.Unix())

	go p.Time()

	for p.Flag == 0 {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
		p.UI()
		p.Play(words)
	}

	fmt.Println("\n>> TÚ PIERDES D:")
	os.Exit(1)
}

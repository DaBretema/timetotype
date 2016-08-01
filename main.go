// Verano de 2016.
// Daniel Camba Lamas.
// JUEGO PROPIO: TimeToType.

/* TO DO...
- Añadir GUI.
- Elegir dificultad al inicio.
- Detener programa al ganar o perder.
- Mejorar 'formula' inc/decr del tiempo.
- Particionar map 'words' por dificultad.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dcabalas/TimeToType/back"
)

func main() {
	t := time.Now()
	rand.Seed(t.Unix())

	words := back.FileMap("resources/words.txt")
	p := back.NewPlayer(2*time.Minute, 5, 0, 0, true)

	go p.Time()

	for p.Flag {
		back.ResetPos()
		p.ShowData()
		input, wordNum := p.Play(words)
		if !(input == "" && wordNum == 0) {
			p.Check(input, wordNum, words)
		}
	}

	fmt.Println("\n>> TÚ PIERDES D:")
	os.Exit(1)
}

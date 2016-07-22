package player

import (
	"fmt"
	"math/rand"
	"time"
)

// Estructura de un jugador.
type Player struct {
	Timer                    time.Duration
	Lifes, Level, Okay, Flag byte
}

// Constructor.
func NewPlayer(timer time.Duration, lifes, level, okay, flag byte) Player {
	return Player {
		Timer: timer,
		Lifes: lifes,
		Level: level,
		Okay:  okay,
		Flag:  flag }
}

// Decremento del tiempo.
func (p *Player) Time() {
	for p.Timer != 0 {
		time.Sleep(time.Second)
		p.Timer -= time.Duration(p.Level+1) * 2 * time.Second
	}
}

// Linea de visualizacion de datos.
func (p Player) UI() {
	fmt.Println("T...", p.Timer)
	fmt.Printf("♥.%d √.%d ╙.%d\n", p.Lifes, p.Okay, p.Level)
}

// Muestra palabra, solicita entrada.
func (p *Player) Play(words map[int]string) {
	var input string
	if len(words) > 0 {
		if p.Lifes > 0 {
			cont := rand.Intn(len(words))

			fmt.Print("\nSYS: ",words[cont],"\nYOU: " )
			fmt.Scan(&input)

			p.Check(input, cont, words)
		} else {
			p.Flag = 1
		}
	} else {
		fmt.Println("\n\n TÚ GANAS :D")
	}
}

// Comprueba que la ultima palabra sea correcta
// si lo es incrementa aciertos y borra la palabra
// si el numero de aciertos es *5 el nivel del jugador incrementa el nivel
// si no es correcta, resta una vida al jugador.
func (p *Player) Check(in string, cont int, words map[int]string) {
	if in == words[cont] {
		p.Okay++
		delete(words, cont)
		if p.Okay == (p.Level+1)*5 {
			p.Level++
			p.Timer += time.Duration(p.Level+1) * 3 * time.Second
		} else {
			p.Timer += time.Duration(p.Level+1) * 3 * time.Second
		}
	} else {
		p.Lifes--
	}
}

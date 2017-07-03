package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Player es la clase del jugador.
type Player struct {
	Timer              time.Duration
	Flag               bool
	Lifes, Level, Okay byte
}

// NewPlayer es el constructor de la clase jugador.
func NewPlayer(timer time.Duration, lifes, level, okay byte, flag bool) Player {
	return Player{
		Timer: timer,
		Lifes: lifes,
		Level: level,
		Okay:  okay,
		Flag:  flag}
}

// Time manjea el decremento del tiempo.
func (p *Player) Time() {
	// Per second loop
	for p.Timer > 0 {
		time.Sleep(time.Second)
		p.Timer -= time.Duration(p.Level+1) * 2 * time.Second
		fmt.Print(p.Timer)
	}
}

// ShowData printa los datos del jugador.
func (p Player) ShowData() {
	fmt.Println("T...", p.Timer)
	fmt.Printf("♥.%d √.%d ╙.%d\n", p.Lifes, p.Okay, p.Level)
}

// Play muestra una palabra y solicita entrada.
func (p *Player) Play(words map[int]string) (string, int) {
	var input string
	if len(words) > 0 {
		if p.Lifes > 0 {
			cont := rand.Intn(len(words))
			fmt.Print("\nSYS: ", words[cont], "\nYOU: ")
			log.Fatal(fmt.Scan(&input))
			return input, cont
		}
		p.Flag = false
	} else {
		fmt.Println("\n>> TÚ GANAS :D")
		os.Exit(2)
	}
	return "", 0
}

// Check comprueba que la ultima palabra sea correcta
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

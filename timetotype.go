package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	port   = os.Getenv("PORT")
	words  = FillMap("words.txt")
	player = NewPlayer(2*time.Minute, 5, 0, 0, true)
)

func init() {
	// Port out of production
	if port == "" {
		port = "8080"
	}
	// Randomize
	rand.Seed(time.Now().Unix())
}

func main() {
	// // GAME
	// for player.Flag {
	// 	player.ShowData()
	// 	input, wNum := player.Play(words)
	// 	if !(input == "" && wNum == 0) {
	// 		player.Check(input, wNum, words)
	// 	}
	// }
	// // EXIT
	// fmt.Println("\n>> TÚ PIERDES D:")
	// os.Exit(1)

	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//
// HANDLERS
//

// handleRoot responde a las peticiones en "/"
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Parallel timer
	go player.Time()

	// for {
	// 	fmt.Fprintln(w, player.Timer)
	// }
	// for i := 0; i < len(words); i++ {
	// 	fmt.Fprint(w, words[i]+"\n")
	// }
}

//
// FUNCTIONS
//

// FillMap monta linea a linea de un mapa de int a strings dado un fichero.
func FillMap(filePath string) map[int]string {
	f, err := os.Open(filePath)
	// Open errors
	if err != nil {
		log.Fatal(err)
	}
	// Close errors
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	// Get data from file
	data := bufio.NewScanner(f)
	// Populate the map
	words := make(map[int]string)
	for i := 0; data.Scan(); i++ {
		words[i] = data.Text()
	}
	// Return the map
	return words
}

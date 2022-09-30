package main

import (
	"fmt"
	"math/rand"
	"time"
)

type IPlayer interface {
	guess() int
}

type Game struct {
	NumberOfGuesses     int
	NumberTryingToMatch int
	MostRecentGuess     int
	Player              IPlayer
}

func (g *Game) play() {

	num := g.Player.guess()
	g.MostRecentGuess = num
	if g.NumberOfGuesses < 2 {
		g.NumberOfGuesses = g.NumberOfGuesses + 1
		if num == g.NumberTryingToMatch {
			fmt.Println("You win!")

		} else if num > g.NumberTryingToMatch {
			fmt.Println("Too High")
			g.play()

		} else if num < g.NumberTryingToMatch {
			fmt.Println("Too low")
			g.play()
		}

	} else {
		fmt.Println("You ran out of guesses. Game over")
	}

}

type Autoguess struct {
	MinValue    int
	MaxValue    int
	GamePointer *Game
}
type Human struct {
}

// it implements the IPlayer interface
func (h *Human) guess() int {
	fmt.Println("Enter your next guess")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Wrong input ,please restart the game.")
	}
	return i
}

// it implements the IPlayer interface
func (a *Autoguess) guess() int {

	rand.Seed(time.Now().UnixNano())
	n := a.MinValue + rand.Intn(a.MaxValue-a.MinValue+1)
	fmt.Printf("The computer has chosen %d\n", n)
	return n

}

func generateRandomNumber(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	n := max + rand.Intn(max-min+1)
	return n
}

func main() {

	game := &Game{
		NumberOfGuesses:     0,
		NumberTryingToMatch: generateRandomNumber(1, 10),
	}

	auto := &Autoguess{
		MinValue:    1,
		MaxValue:    10,
		GamePointer: game,
	}

	human := &Human{}

	fmt.Print("Welcome To ,Guess a number Game,Press Enter To Continue")
	fmt.Scanln()
	fmt.Print("You have 3 guesses to guess a number from 1 to 10 ,Enter to start")
	fmt.Scanln()
	fmt.Println("Do you want to make the guesses? (y/n -- if n guesses will be generated for you)")
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		panic(err)
	}

	if s == "y" {
		game.Player = human
	} else {
		// auto guess -->
		game.Player = auto
	}

	game.play()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Question struct {
	Text string
	Options []string
	Answer uint8
}

type GameState struct {
	Name string
	Points string
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Welcome to the Go Quiz")
	fmt.Println("Enter your name:")

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Writing string error")
	}

	g.Name = name

	fmt.Printf("Let's get start %s", g.Name)
}

func (g* GameState) ProcessCSV() {
	f, err := os.Open("quiz-go.csv")
	
	if err != nil {
		panic("Failed to open questions")
	}

	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		panic("Failed to read questions")
	}

	fmt.Println(content)

	/*
		file []byte

		when byte == '\n' new line
		when byte == '\r' continue

		column = []byte
		columns = []column

		question: {
			Text: string
			Options: []string
			Answer: uint8
		}
	*/

}

func main() {
	game := &GameState{}

	// game.Init()
	game.ProcessCSV()

	fmt.Println(game.Questions)
}

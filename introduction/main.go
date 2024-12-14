package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Question struct {
	Text string
	Options []string
	Answer uint8
}

type GameState struct {
	Name string
	Points uint8
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

	fmt.Printf("Let's get started %s", g.Name)
}

func (g* GameState) ProcessCSV() {
	start := time.Now()
	file, err := os.Open("quiz-go.csv")

	if err != nil {
		panic("Failed to open quiz questions file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic("Failed to read quiz questions file")
	}

	isFirstLine := true

	var column []byte
	var columns [][]byte

	for i, b := range content {
		if b == 13 {
			continue
		}

		if b == 10 && isFirstLine {
			isFirstLine = false

			column = column[:0]
			columns = columns[:0]

			continue
		}

		if b == 44 {
			newColumn := make([]byte, len(column))
			copy(newColumn, column)

			columns = append(columns, newColumn)
			column = column[:0]
			continue
		}

		if i == len(content) - 1 {
			column = append(column, b)
		}

		if b == 10 || i == len(content) - 1 {
			columns = append(columns, column)

			question := Question{
				Text: string(columns[0]),
				Options: []string{
					string(columns[1]),
					string(columns[2]),
					string(columns[3]),
					string(columns[4]),
				},
				Answer: columns[5][0] - 48,
			}

			g.Questions = append(g.Questions, question)

			column = column[:0]
			columns = columns[:0]
			continue
		}

		column = append(column, b)
	}

	fmt.Printf("Question files loaded in %d nanoseconds\n", time.Since(start).Nanoseconds())
}

func (g * GameState) Run() {
	for i, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", i + 1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Enter the number of the correct answer")

		var answer uint8
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[0:len(read) - 1])
			
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		fmt.Println("---------------")
		if answer == question.Answer {
			fmt.Println("Congratulations, You got the answer right")
			g.Points += 10
		} else {
			fmt.Println("Ops! You got the answer wrong")
		}
		fmt.Println("---------------")
	}
}

func toInt(s string) (uint8, error) {
	const message = "only numbers from 1 to 4 are allowed, please enter a correct alternative"
	if len(s) > 0 && s[len(s)-1] == '\r' {
		s = s[:len(s)-1]
	}
	
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New(message)
	}

	if i < 1 || i > 4 {
		return 0, errors.New(message)
	}
	return uint8(i), nil
}

func main() {

	game := &GameState{}
	go game.ProcessCSV()
	
	game.Init()
	game.Run()

	fmt.Printf("Game over, you made %d points.\nPress enter to close the game", game.Points)
	
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player interface {
	GetMove(game *TicTacToe) int
}

type HumanPlayer struct {
}

func (p HumanPlayer) GetMove(game *TicTacToe) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(game.Turn() + "'s Turn. Enter move (0-8): ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Invalid Input. Please try again.\n")
			continue
		}
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		move, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Move must be an integer. Please try again.\n")
			continue
		}

		if move < 0 || move > 8 {
			fmt.Print("Move must be between 0 and 8. Please try again.\n")
			continue
		}

		var validMove bool = false
		for _, v := range game.GenerateMoves() {
			if v == move {
				validMove = true
				break
			}
		}

		if !validMove {
			fmt.Print("Invalid move. Please try again.\n")
			continue
		}

		return move
	}
}

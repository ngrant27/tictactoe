package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type TicTacToe struct {
	board   []string
	winner  string
	history []int
	turn    string
}

const (
	empty string = "_"
	x     string = "X"
	o     string = "O"
	none  string = "N"
	draw  string = "D"
)

func NewGame() TicTacToe {
	g := TicTacToe{
		board:   make([]string, 9),
		winner:  none,
		history: make([]int, 9),
		turn:    x,
	}

	for i := range g.board {
		g.board[i] = empty
	}

	return g
}

func (g *TicTacToe) PrintBoard(writer io.Writer) {
	for i, v := range g.board {
		fmt.Fprintf(writer, "| "+v+" ")
		if i%3 == 2 {
			fmt.Fprintf(writer, "|\n")
		}
	}
}

func (g *TicTacToe) PrintBoardNumbers(writer io.Writer) {
	for i := range g.board {
		fmt.Fprintf(writer, "| "+strconv.Itoa(i)+" ")
		if i%3 == 2 {
			fmt.Fprintf(writer, "|\n")
		}
	}
}

func (g *TicTacToe) Turn() string {
	return g.turn
}

func (g *TicTacToe) makeMove(square int, letter string) bool {
	if g.board[square] != empty {
		return false
	}

	g.board[square] = letter
	if g.checkWinner(square, letter) {
		g.winner = letter
	}
	return true
}

func (g *TicTacToe) checkWinner(square int, letter string) bool {
	// check the row
	rowIndex := square / 3
	if g.board[rowIndex*3] == g.board[rowIndex*3+1] && g.board[rowIndex*3] == g.board[rowIndex*3+2] {
		return true
	}

	// check the column
	colIndex := square % 3
	if g.board[colIndex] == g.board[colIndex+3] && g.board[colIndex] == g.board[colIndex+6] {
		return true
	}

	// checking if the square is not on diagonals
	if square%2 != 0 || letter != g.board[4] {
		return false
	}

	// check the diagonals
	if g.board[0] == g.board[8] && g.board[0] == g.board[4] {
		return true
	}

	if g.board[2] == g.board[6] && g.board[2] == g.board[4] {
		return true
	}

	return false
}

func (g *TicTacToe) hasEmptySquares() bool {
	for _, v := range g.board {
		if v == empty {
			return true
		}
	}
	return false
}

/*
func (g *TicTacToe) numEmptySquares() int {
	var count int = 0
	for _, v := range g.board {
		if v == empty {
			count++
		}
	}
	return count
}*/

func (g *TicTacToe) GenerateMoves() []int {
	var moveList []int
	for i, v := range g.board {
		if v == empty {
			moveList = append(moveList, i)
		}
	}
	return moveList
}

func (g *TicTacToe) Play(XPlayer, OPlayer Player) {
	var move int

	for g.hasEmptySquares() {
		// get player move
		if g.turn == x {
			move = XPlayer.GetMove(g)
		} else {
			move = OPlayer.GetMove(g)
		}

		// make move, then print board
		if g.makeMove(move, g.turn) {
			fmt.Println(g.turn + " makes a move to square " + strconv.Itoa(move))
			g.PrintBoard(os.Stdout)
			fmt.Print("\n")
		}

		// check if a player has won
		if g.winner != none {
			fmt.Println(g.turn + " wins!!")
			return
		}

		// change player
		if g.turn == x {
			g.turn = o
		} else {
			g.turn = x
		}
	}

	// board is full and no winner, so we have a draw
	fmt.Println("The game is a draw!")
}

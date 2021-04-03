package main

import (
	"bytes"
	"testing"
)

func TestPrintBoard(t *testing.T) {
	t.Run("test to make sure we print the board", func(t *testing.T) {
		buffer := bytes.Buffer{}

		game := NewGame()
		game.PrintBoard(&buffer)

		got := buffer.String()
		want := "| _ | _ | _ |\n| _ | _ | _ |\n| _ | _ | _ |\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test to make sure we print board numbers correctly", func(t *testing.T) {
		buffer := bytes.Buffer{}

		game := NewGame()
		game.PrintBoardNumbers(&buffer)

		got := buffer.String()
		want := "| 0 | 1 | 2 |\n| 3 | 4 | 5 |\n| 6 | 7 | 8 |\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestMakeMove(t *testing.T) {
	t.Run("makeMove changes the board state correctly and returns the correct value", func(t *testing.T) {
		buffer := bytes.Buffer{}

		game := NewGame()
		game.makeMove(1, "X")
		game.makeMove(8, "O")
		game.makeMove(3, "X")

		game.PrintBoard(&buffer)

		got := buffer.String()
		want := "| _ | X | _ |\n| X | _ | _ |\n| _ | _ | O |\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test to make sure makeMove does nothing and returns false when you enter a square that is already occupied", func(t *testing.T) {
		var isMoveSuccessful bool = true
		buffer := bytes.Buffer{}

		game := NewGame()

		isMoveSuccessful = isMoveSuccessful && game.makeMove(1, "X")
		isMoveSuccessful = isMoveSuccessful && game.makeMove(8, "O")
		isMoveSuccessful = isMoveSuccessful && game.makeMove(8, "X")

		game.PrintBoard(&buffer)

		got := buffer.String()
		want := "| _ | X | _ |\n| _ | _ | _ |\n| _ | _ | O |\n"

		if isMoveSuccessful {
			t.Errorf("Wanted makeMove to fail but it didn't")
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

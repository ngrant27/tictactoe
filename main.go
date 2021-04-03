package main

func main() {
	X := HumanPlayer{}
	Y := HumanPlayer{}

	game := NewGame()
	game.Play(X, Y)
}

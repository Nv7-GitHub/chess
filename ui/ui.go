package ui

import "github.com/Nv7-Github/chess/chess"

const WIDTH = 480
const HEIGHT = 480

type UI struct {
	board *chess.Board
}

func NewUI(board *chess.Board) *UI {
	return &UI{
		board: board,
	}
}

func (u *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

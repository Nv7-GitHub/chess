package ui

import "github.com/Nv7-Github/chess/chess"

const WIDTH = 480
const HEIGHT = 480

type UI struct {
	board *chess.Board

	selected    chess.Pos
	hasSelected bool
	canMove     [][]bool

	hoverPos chess.Pos
	hover    bool
}

func NewUI(board *chess.Board) *UI {
	u := &UI{
		board: board,

		hasSelected: false,
		selected:    chess.Pos{Row: -1, Col: -1},
	}
	u.canMove = make([][]bool, len(board.Pieces))
	for i := range u.canMove {
		u.canMove[i] = make([]bool, len(board.Pieces[i]))
	}
	return u
}

func (u *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

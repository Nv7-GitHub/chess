package ui

import (
	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (u *UI) Update() error {
	// Check if clicking on square
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		posx, posy := ebiten.CursorPosition()
		squareWidth, squareHeight := WIDTH/len(u.board.Pieces[0]), HEIGHT/len(u.board.Pieces)
		squareX, squareY := posx/squareWidth, posy/squareHeight

		if u.board.Pieces[squareY][squareX] == nil { // Clicking on no piece
			u.hasSelected = false
		} else if u.selected.Col == squareX && u.selected.Row == squareY { // Clicking on that piece
			u.hasSelected = !u.hasSelected
		} else { // Clicking on another piece
			u.hasSelected = true
			u.selected.Col = squareX
			u.selected.Row = squareY
			u.RecalcCanMove()
		}
	}
	return nil
}

func (u *UI) RecalcCanMove() {
	p := u.board.Piece(u.selected)

	for r, v := range u.board.Pieces {
		for c := range v {
			u.canMove[r][c] = p.CanMoveTo(u.board, u.selected, chess.Pos{Row: r, Col: c})
		}
	}
}

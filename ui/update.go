package ui

import (
	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (u *UI) Update() error {
	if u.board.Checkmate {
		return nil
	}

	// Hover
	posx, posy := ebiten.CursorPosition()
	posy = HEIGHT - posy // flip

	squareWidth, squareHeight := WIDTH/len(u.board.Pieces[0]), HEIGHT/len(u.board.Pieces)
	squareX, squareY := posx/squareWidth, posy/squareHeight

	p := u.board.Piece(chess.Pos{Row: squareY, Col: squareX})
	if p != nil && p.Side() == u.board.Turn {
		u.hover = true
		u.hoverPos = chess.Pos{Row: squareY, Col: squareX}
	} else {
		u.hover = false
		u.hoverPos = chess.Pos{Row: -1, Col: -1}
	}

	// Check if clicking on square
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Is moving piece?
		if u.hasSelected && u.canMove[squareY][squareX] {
			u.board.Move(u.selected, chess.Pos{Row: squareY, Col: squareX})
			u.board.NextTurn()
			u.hasSelected = false
			return nil
		}

		if u.board.Pieces[squareY][squareX] == nil { // Clicking on no piece
			u.hasSelected = false
		} else if u.selected.Col == squareX && u.selected.Row == squareY { // Clicking on that piece
			u.hasSelected = !u.hasSelected
		} else if u.board.Piece(chess.Pos{Row: squareY, Col: squareX}).Side() == u.board.Turn { // Clicking on another piece, check if its your turn
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
			if r == u.selected.Row && c == u.selected.Col {
				u.canMove[r][c] = false
				continue
			}
			u.canMove[r][c] = p.CanMoveTo(u.board, u.selected, chess.Pos{Row: r, Col: c})
		}
	}
}

package chess

type Pawn struct {
	BasicPiece
}

func (p *Pawn) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	if !p.BasicPiece.CanMoveTo(board, currPos, newPos) {
		return false
	}

	// Check if going 2 squares forward, make sure nothing in that spot
	if iabs(currPos.Row-newPos.Row) == 2 && (currPos.Row == 1 || currPos.Row == 6) && (currPos.Col-newPos.Col) == 0 && board.Piece(newPos) == nil {
		// Make sure nothing in the way
		return board.Piece(Pos{currPos.Row + (newPos.Row-currPos.Row)/2, currPos.Col}) == nil
	}

	// Check if going forwards
	dir := 1
	if p.Side() == BLACK {
		dir = -1
	}
	if newPos.Row-currPos.Row != dir {
		return false
	}

	// Check if to the side
	if newPos.Col-currPos.Col == 0 && board.Piece(newPos) != nil { // Check if capturing non-diagonally
		return false
	}

	// Check if capturing diagonally
	if iabs(newPos.Col-currPos.Col) == 1 && board.Piece(newPos) != nil {
		return true
	} else if iabs(newPos.Col-currPos.Col) != 0 { // Check if going diagonal when not capturing
		return false
	}

	return p.BasicPiece.PostCanMoveTo(board, currPos, newPos)
}

func (p *Pawn) Type() PieceType {
	return PAWN
}

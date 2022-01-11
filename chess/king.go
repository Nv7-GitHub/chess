package chess

type King struct {
	BasicPiece
}

func (k *King) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	if !k.BasicPiece.CanMoveTo(board, currPos, newPos) {
		return false
	}

	diff := newPos.Sub(currPos)
	diff.Row = iabs(diff.Row)
	diff.Col = iabs(diff.Col)

	if !(diff.Row <= 1 && diff.Col <= 1) {
		return false
	}

	// Move temporarily to the new position & check if in check
	p := board.Piece(newPos)
	board.Pieces[newPos.Row][newPos.Col] = k
	board.Pieces[currPos.Row][currPos.Col] = nil

	inCheck := board.IsPosCheck(board.Turn, newPos) // Check if in check

	// Move back
	board.Pieces[currPos.Row][currPos.Col] = k
	board.Pieces[newPos.Row][newPos.Col] = p

	if inCheck {
		return false
	}

	return k.BasicPiece.PostCanMoveTo(board, currPos, newPos)
}

func (k *King) Type() PieceType {
	return KING
}

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

	return diff.Row <= 1 && diff.Col <= 1
}

func (k *King) Type() PieceType {
	return KING
}

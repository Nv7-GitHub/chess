package chess

type Knight struct {
	BasicPiece
}

func (k *Knight) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	if !k.BasicPiece.CanMoveTo(board, currPos, newPos) {
		return false
	}

	diff := newPos.Sub(currPos)
	diff.Row = iabs(diff.Row)
	diff.Col = iabs(diff.Col)

	return (diff.Row == 2 && diff.Col == 1) || (diff.Row == 1 && diff.Col == 2)
}

func (k *Knight) Type() PieceType {
	return KNIGHT
}

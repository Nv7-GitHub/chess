package chess

type Bishop struct {
	BasicPiece
}

func (b *Bishop) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	if !b.BasicPiece.CanMoveTo(board, currPos, newPos) {
		return false
	}

	// Check if going in one dir
	if iabs(newPos.Row-currPos.Row) != iabs(newPos.Col-currPos.Col) {
		return false
	}
	if newPos.Row-currPos.Row == 0 {
		return false
	}

	// Check dir
	changeRow := 0
	changeCol := 0
	if newPos.Row-currPos.Row > 0 {
		changeRow = 1
	} else if newPos.Row-currPos.Row < 0 {
		changeRow = -1
	}
	if newPos.Col-currPos.Col > 0 {
		changeCol = 1
	} else if newPos.Col-currPos.Col < 0 {
		changeCol = -1
	}

	// Get path length
	diff := newPos.Sub(currPos)
	change := iabs(diff.Row)

	// Check path
	rV := currPos.Row + changeRow
	cV := currPos.Col + changeCol
	for i := 0; i < change-1; i++ { // change-1 makes it ignore the last one, which allows things at the very end to be killed
		if board.Piece(Pos{rV, cV}) != nil {
			return false
		}
		rV += changeRow
		cV += changeCol
	}

	return true
}

func (b *Bishop) Type() PieceType {
	return BISHOP
}

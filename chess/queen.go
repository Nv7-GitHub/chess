package chess

type Queen struct {
	BasicPiece
}

func (q *Queen) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	if !q.BasicPiece.CanMoveTo(board, currPos, newPos) {
		return false
	}

	// Get dir
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

	if changeRow == 0 && changeCol == 0 {
		return false
	}

	// Get path length
	diff := newPos.Sub(currPos)
	change := 0
	if iabs(diff.Row) > change {
		change = iabs(diff.Row)
	}
	if iabs(diff.Col) > change {
		change = iabs(diff.Col)
	}

	// Check if going const amount
	if newPos.Row-currPos.Row != changeRow*change {
		return false
	}
	if newPos.Col-currPos.Col != changeCol*change {
		return false
	}

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

	return q.BasicPiece.PostCanMoveTo(board, currPos, newPos)
}

func (q *Queen) Type() PieceType {
	return QUEEN
}

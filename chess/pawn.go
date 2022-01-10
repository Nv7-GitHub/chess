package chess

type Pawn struct {
	BasicPiece
}

func (p *Pawn) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	// Check if going forwards
	dir := 1
	if p.Side() == BLACK {
		dir = -1
	}
	if newPos.Row-currPos.Row != dir {
		return false
	}

	// Check if to the side
	if iabs(newPos.Col-currPos.Col) != 1 {
		return false
	}

	return true
}

func (p *Pawn) Type() PieceType {
	return PAWN
}

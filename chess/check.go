package chess

func (b *Board) IsPosCheck(side Side, kingPos Pos) bool {
	oppositeSide := WHITE
	if side == WHITE {
		oppositeSide = BLACK
	}

	// Go through all black pieces and check if they can go to the king
	for r, v := range b.Pieces {
		for c, p := range v {
			if p != nil && p.Side() == oppositeSide && p.CanMoveTo(b, Pos{r, c}, kingPos) {
				return true
			}
		}
	}

	return false
}

func (b *Board) IsCheckMate() bool {
	var king Piece
	var kingPos Pos
	// Find king
	for r, v := range b.Pieces {
		for c, p := range v {
			if p != nil && p.Side() == b.Turn && p.Type() == KING {
				king = p
				kingPos = Pos{r, c}
			}
		}
	}

	// Check if king in check
	if !b.IsPosCheck(b.Turn, kingPos) {
		return false
	}

	// Check if king can go anywhere
	for r, v := range b.Pieces {
		for c := range v {
			if king.CanMoveTo(b, kingPos, Pos{r, c}) {
				return false
			}
		}
	}
	return true
}

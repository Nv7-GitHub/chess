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

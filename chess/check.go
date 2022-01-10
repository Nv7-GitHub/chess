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

type CheckMateResult int

const (
	CheckMateNone CheckMateResult = iota
	CheckMateCheck
	CheckMateCheckMate
)

func (b *Board) IsCheckMate() CheckMateResult {
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
		return CheckMateNone
	}

	// Check if king can go anywhere
	for r, v := range b.Pieces {
		for c := range v {
			if king.CanMoveTo(b, kingPos, Pos{r, c}) {
				return CheckMateCheck
			}
		}
	}

	// Try all the moves to see if you can stop the check
	for r, v := range b.Pieces {
		for c, p := range v {
			if p != nil && p.Side() == b.Turn {
				// Try all moves
				for r2, v2 := range b.Pieces {
					for c2 := range v2 {
						if p.CanMoveTo(b, Pos{r, c}, Pos{r2, c2}) {
							// Simulate move
							oldP := b.Pieces[r2][c2]
							b.Pieces[r2][c2] = p
							b.Pieces[r][c] = nil

							// Check if king is still in check
							isInCheck := b.IsPosCheck(b.Turn, kingPos)

							// Undo
							b.Pieces[r][c] = p
							b.Pieces[r2][c2] = oldP

							if !isInCheck {
								return CheckMateCheck
							}
						}
					}
				}
			}
		}
	}
	return CheckMateCheckMate
}

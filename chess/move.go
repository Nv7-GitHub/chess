package chess

func (b *Board) Move(start Pos, end Pos) {
	p := b.Piece(start)
	b.Pieces[start.Row][start.Col] = nil
	b.Pieces[end.Row][end.Col] = p
}

func (b *Board) NextTurn() {
	if b.Turn == WHITE {
		b.Turn = BLACK
	} else {
		b.Turn = WHITE
	}
	res := b.IsCheckMate()
	switch res {
	case CheckMateNone:
		b.Check = false
		b.Checkmate = false

	case CheckMateCheck:
		b.Check = true
		b.Checkmate = false

	case CheckMateCheckMate:
		b.Check = true
		b.Checkmate = true
	}
}

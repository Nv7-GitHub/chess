package chess

func iabs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (b *Board) Piece(pos Pos) Piece {
	return b.Pieces[pos.Row][pos.Col]
}

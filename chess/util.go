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

func (p *Pos) Sub(p2 Pos) Pos {
	return Pos{Row: p.Row - p2.Row, Col: p.Col - p2.Col}
}

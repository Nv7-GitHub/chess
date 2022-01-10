package chess

func iabs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (b *Board) Piece(pos Pos) Piece {
	if pos.Row > 7 || pos.Row < 0 || pos.Col > 7 || pos.Col < 0 {
		return nil
	}
	return b.Pieces[pos.Row][pos.Col]
}

func (p *Pos) Sub(p2 Pos) Pos {
	return Pos{Row: p.Row - p2.Row, Col: p.Col - p2.Col}
}

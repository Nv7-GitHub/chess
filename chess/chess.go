package chess

type Pos struct {
	Row int
	Col int
}

type Side int

const (
	WHITE Side = iota // White at 0, black at 8
	BLACK
)

type PieceType int

const (
	PAWN PieceType = iota
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
)

type Piece interface {
	SetSide(side Side)
	Side() Side
	CanMoveTo(board *Board, currPos Pos, newPos Pos) bool
	Type() PieceType
}

type BasicPiece struct {
	side Side
}

func (b *BasicPiece) SetSide(side Side) {
	b.side = side
}

func (b *BasicPiece) Side() Side {
	return b.side
}

func (b *BasicPiece) CanMoveTo(board *Board, currPos Pos, newPos Pos) bool {
	return board.Piece(newPos) == nil || (board.Piece(newPos) != nil && board.Piece(newPos).Side() != b.Side()) // Checks if there is a piece of your side there
}

type Board struct {
	Pieces [8][8]Piece
	Turn   Side
}

func NewBoard() *Board {
	board := new(Board)
	board.Turn = WHITE
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			board.Pieces[row][col] = nil
		}
	}

	// Set up board
	// Pawns
	for i := 0; i < 8; i++ {
		board.Pieces[1][i] = &Pawn{}
		board.Pieces[1][i].SetSide(WHITE)

		board.Pieces[6][i] = &Pawn{}
		board.Pieces[6][i].SetSide(BLACK)
	}

	// Rooks
	board.Pieces[0][0] = &Rook{}
	board.Pieces[0][0].SetSide(WHITE)
	board.Pieces[0][7] = &Rook{}
	board.Pieces[0][7].SetSide(WHITE)
	board.Pieces[7][0] = &Rook{}
	board.Pieces[7][0].SetSide(BLACK)
	board.Pieces[7][7] = &Rook{}
	board.Pieces[7][7].SetSide(BLACK)

	// Bishops
	board.Pieces[0][1] = &Bishop{}
	board.Pieces[0][1].SetSide(WHITE)
	board.Pieces[0][6] = &Bishop{}
	board.Pieces[0][6].SetSide(WHITE)
	board.Pieces[7][1] = &Bishop{}
	board.Pieces[7][1].SetSide(BLACK)
	board.Pieces[7][6] = &Bishop{}
	board.Pieces[7][6].SetSide(BLACK)

	// Knights
	board.Pieces[0][2] = &Knight{}
	board.Pieces[0][2].SetSide(WHITE)
	board.Pieces[0][5] = &Knight{}
	board.Pieces[0][5].SetSide(WHITE)
	board.Pieces[7][2] = &Knight{}
	board.Pieces[7][2].SetSide(BLACK)
	board.Pieces[7][5] = &Knight{}
	board.Pieces[7][5].SetSide(BLACK)

	// Kings
	board.Pieces[0][4] = &King{}
	board.Pieces[0][4].SetSide(WHITE)
	board.Pieces[7][4] = &King{}
	board.Pieces[7][4].SetSide(BLACK)

	// Queens
	board.Pieces[0][3] = &Queen{}
	board.Pieces[0][3].SetSide(WHITE)
	board.Pieces[7][3] = &Queen{}
	board.Pieces[7][3].SetSide(BLACK)

	return board
}

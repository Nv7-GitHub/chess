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

type Board struct {
	Pieces [8][8]Piece
}

func NewBoard() *Board {
	board := new(Board)
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
	return board
}

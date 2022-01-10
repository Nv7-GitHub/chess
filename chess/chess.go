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
	return board
}

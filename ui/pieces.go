package ui

import (
	"embed"
	"image/png"

	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed pieces/*.png
var pieces embed.FS

var blackPieces = map[chess.PieceType]*ebiten.Image{}
var whitePieces = map[chess.PieceType]*ebiten.Image{}

var pieceTypeNames = map[string]chess.PieceType{
	"pawn":   chess.PAWN,
	"knight": chess.KNIGHT,
	"bishop": chess.BISHOP,
	"rook":   chess.ROOK,
	"queen":  chess.QUEEN,
}

func init() {
	for k, v := range pieceTypeNames {
		// Black
		blackf, err := pieces.Open("pieces/black_" + k + ".png")
		if err != nil {
			panic(err)
		}
		img, err := png.Decode(blackf)
		if err != nil {
			panic(err)
		}
		blackPieces[v] = ebiten.NewImageFromImage(img)
		blackf.Close()

		// White
		whitef, err := pieces.Open("pieces/white_" + k + ".png")
		if err != nil {
			panic(err)
		}
		img, err = png.Decode(whitef)
		if err != nil {
			panic(err)
		}
		whitePieces[v] = ebiten.NewImageFromImage(img)
		whitef.Close()
	}
}

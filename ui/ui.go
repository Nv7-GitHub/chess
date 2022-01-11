package ui

import (
	"image/color"

	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
)

const width = 480
const height = 480

var WIDTH int
var HEIGHT int

type UI struct {
	board *chess.Board

	selected    chess.Pos
	hasSelected bool
	canMove     [][]bool

	hoverPos chess.Pos
	hover    bool

	colorImages map[color.Color]*ebiten.Image
}

func NewUI(board *chess.Board) *UI {
	u := &UI{
		board: board,

		hasSelected: false,
		selected:    chess.Pos{Row: -1, Col: -1},

		colorImages: make(map[color.Color]*ebiten.Image),
	}
	u.canMove = make([][]bool, len(board.Pieces))
	for i := range u.canMove {
		u.canMove[i] = make([]bool, len(board.Pieces[i]))
	}
	return u
}

func (u *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	WIDTH = int(width * s)
	HEIGHT = int(height * s)
	return WIDTH, HEIGHT + HEIGHT/len(u.board.Pieces) // Add space on bottom, mult by hidpi
}

func (u *UI) getImage(color color.Color) *ebiten.Image {
	im, exists := u.colorImages[color]
	if !exists {
		im = ebiten.NewImage(1, 1)
		im.Fill(color)
		u.colorImages[color] = im
	}
	return im
}

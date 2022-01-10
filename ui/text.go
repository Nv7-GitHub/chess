package ui

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed OpenSans.ttf
var openSans []byte

const dpi = 72
const size = 48

var fontBottom font.Face

func init() {
	tt, err := opentype.Parse(openSans)
	if err != nil {
		panic(err)
	}

	scale := ebiten.DeviceScaleFactor()
	fontBottom, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size * scale,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}

func (u *UI) DrawTextBottom(txt string, screen *ebiten.Image) {
	bounds := text.BoundString(fontBottom, txt)
	bottomHeight := HEIGHT / len(u.board.Pieces)
	y := HEIGHT + bottomHeight - (bottomHeight-bounds.Dy())/2
	x := (WIDTH - bounds.Max.X) / 2
	text.Draw(screen, txt, fontBottom, x, y, color.White)
}

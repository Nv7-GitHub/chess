package ui

import (
	"fmt"

	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (u *UI) Draw(screen *ebiten.Image) {
	u.DrawSquares(screen)
	u.DrawPieces(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
}

func (u *UI) DrawPieces(screen *ebiten.Image) {
	for r, v := range u.board.Pieces {
		for c, p := range v {
			if p == nil {
				continue
			}

			var im *ebiten.Image
			if p.Side() == chess.WHITE {
				im = whitePieces[p.Type()]
			} else {
				im = blackPieces[p.Type()]
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(c)*float64(WIDTH/len(v)), float64(r)*float64(HEIGHT/len(u.board.Pieces)))
			screen.DrawImage(im, opts)
		}
	}
}

func (u *UI) DrawSquares(screen *ebiten.Image) {
	im := ebiten.NewImage(1, 1)

	for r, v := range u.board.Pieces {
		for c := range v {
			if (r+c)%2 == 1 {
				im.Set(0, 0, whiteBoardSquare)
			} else {
				im.Set(0, 0, blackBoardSquare)
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(c), float64(r))
			opts.GeoM.Scale(float64(WIDTH/len(v)), float64(HEIGHT/len(u.board.Pieces)))
			screen.DrawImage(im, opts)
		}
	}
}

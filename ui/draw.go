package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (u *UI) Draw(screen *ebiten.Image) {
	u.DrawSquares(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
}

func (u *UI) DrawSquares(screen *ebiten.Image) {
	im := ebiten.NewImage(1, 1)
	im.Set(0, 0, color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 0)
	opts.GeoM.Scale(60, 60)
	screen.DrawImage(im, opts)

	for r, v := range u.board.Pieces {
		for c := range v {
			if (r+c)%2 == 1 {
				im.Set(0, 0, color.Black)
			} else {
				im.Set(0, 0, color.White)
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(c), float64(r))
			opts.GeoM.Scale(float64(WIDTH/len(v)), float64(HEIGHT/len(u.board.Pieces)))
			screen.DrawImage(im, opts)
		}
	}
}

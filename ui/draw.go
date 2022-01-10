package ui

import (
	"fmt"

	"github.com/Nv7-Github/chess/chess"
	"github.com/hajimehoshi/ebiten/v2"
)

func (u *UI) Draw(screen *ebiten.Image) {
	// Draw
	u.DrawSquares(screen)
	u.DrawPieces(screen)

	turnName := "White"
	if u.board.Turn == chess.BLACK {
		turnName = "Black"
	}
	u.DrawTextBottom(fmt.Sprintf("%s's Turn", turnName), screen)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
}

func (u *UI) DrawPieces(screen *ebiten.Image) {
	for r, v := range u.board.Pieces {
		for c, p := range v {
			if p == nil {
				continue
			}

			var im *ebiten.Image
			opts := &ebiten.DrawImageOptions{}
			if p.Side() == chess.WHITE {
				im = whitePieces[p.Type()]
			} else {
				im = blackPieces[p.Type()]
			}

			// Scale
			scale := ebiten.DeviceScaleFactor()
			opts.GeoM.Scale(scale, scale)
			opts.Filter = ebiten.FilterLinear

			// Translate
			opts.GeoM.Translate(float64(c)*float64(WIDTH/len(v)), float64(HEIGHT)-(float64(r+1)*float64(HEIGHT/len(u.board.Pieces)))) // Height flipped
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
			opts.GeoM.Translate(float64(c), float64(len(u.board.Pieces)-r-1))
			opts.GeoM.Scale(float64(WIDTH/len(v)), float64(HEIGHT/len(u.board.Pieces)))

			screen.DrawImage(im, opts)

			drawSpecial := false
			if u.hasSelected {
				// Add color filters
				if r == u.selected.Row && c == u.selected.Col {
					im.Set(0, 0, selectedPiece)
					drawSpecial = true
				} else if u.canMove[r][c] {
					drawSpecial = true
					if u.board.Piece(chess.Pos{Row: r, Col: c}) != nil {
						im.Set(0, 0, killColor)
					} else {
						im.Set(0, 0, canMoveSquare)
					}
				}
			} else if u.hover && u.hoverPos.Row == r && u.hoverPos.Col == c { // Hovering
				drawSpecial = true
				im.Set(0, 0, hoverColor)
			}

			// Overlay special squares
			if drawSpecial {
				opts = &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(c), float64(len(u.board.Pieces)-r-1))
				opts.GeoM.Scale(float64(WIDTH/len(v)), float64(HEIGHT/len(u.board.Pieces)))
				screen.DrawImage(im, opts)
			}
		}
	}
}

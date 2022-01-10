package main

import (
	"log"

	"github.com/Nv7-Github/chess/chess"
	"github.com/Nv7-Github/chess/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	board := chess.NewBoard()
	ui := ui.NewUI(board)

	ebiten.SetWindowSize(480, 480)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(ui); err != nil {
		log.Fatal(err)
	}
}

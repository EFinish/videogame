package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ticTacToeBoardimg *ebiten.Image
	ximg              *ebiten.Image
	oimg              *ebiten.Image
	titleimg          *ebiten.Image
)

func init() {
	var err error

	ticTacToeBoardimg, _, err = ebitenutil.NewImageFromFile("img/tic_tac_toe_board.png")

	if err != nil {
		log.Fatal(err)
	}

	ximg, _, err = ebitenutil.NewImageFromFile("img/x.png")

	if err != nil {
		log.Fatal(err)

	}

	oimg, _, err = ebitenutil.NewImageFromFile("img/o.png")

	if err != nil {
		log.Fatal(err)
	}

	titleimg, _, err = ebitenutil.NewImageFromFile("img/amazing_title.png")

	if err != nil {
		log.Fatal(err)
	}

}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(ticTacToeBoardimg, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 680, 380
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

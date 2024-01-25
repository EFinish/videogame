package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ticTacToeBoardImg *ebiten.Image
	ximg              *ebiten.Image
	oimg              *ebiten.Image
	titleImg          *ebiten.Image
	screenWidth       int = 800
	screenHeight      int = 600
)

func init() {
	var err error

	ticTacToeBoardImg, _, err = ebitenutil.NewImageFromFile("img/tic_tac_toe_board.png")

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

	titleImg, _, err = ebitenutil.NewImageFromFile("img/amazing_title.png")

	if err != nil {
		log.Fatal(err)
	}

}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	boardImgWidth := ticTacToeBoardImg.Bounds().Dx()
	boardWidthScaleFactor := float64(screenWidth) / float64(boardImgWidth)
	boardHeightScaleFactor := float64(screenHeight) * 0.75 / float64(ticTacToeBoardImg.Bounds().Dy())
	boardOpts := &ebiten.DrawImageOptions{}

	boardOpts.GeoM.Scale(boardWidthScaleFactor, boardHeightScaleFactor)
	boardOpts.GeoM.Translate(0, float64(screenHeight)*0.25)

	screen.DrawImage(ticTacToeBoardImg, boardOpts)

	titleImgWidth := titleImg.Bounds().Dx()
	titleWidthScaleFactor := float64(screenWidth) / float64(titleImgWidth)
	titleHeightScaleFactor := float64(screenHeight) * 0.25 / float64(titleImg.Bounds().Dy())
	titleOpts := &ebiten.DrawImageOptions{}

	titleOpts.GeoM.Scale(titleWidthScaleFactor, titleHeightScaleFactor)

	screen.DrawImage(titleImg, titleOpts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

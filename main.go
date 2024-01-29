package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ticTacToeBoardImg *ebiten.Image
	xImg              *ebiten.Image
	oImg              *ebiten.Image
	titleImg          *ebiten.Image
	xTurnImg          *ebiten.Image
	oTurnImg          *ebiten.Image
	xWinningImg       *ebiten.Image
	oWinningImg       *ebiten.Image
	winner            string
	isXTurn           bool = true
	screenWidth       int  = 800
	screenHeight      int  = 600
)

func init() {
	var err error

	ticTacToeBoardImg, _, err = ebitenutil.NewImageFromFile("img/tic_tac_toe_board.png")

	if err != nil {
		log.Fatal(err)
	}

	xImg, _, err = ebitenutil.NewImageFromFile("img/x.png")

	if err != nil {
		log.Fatal(err)

	}

	oImg, _, err = ebitenutil.NewImageFromFile("img/o.png")

	if err != nil {
		log.Fatal(err)
	}

	titleImg, _, err = ebitenutil.NewImageFromFile("img/amazing_title.png")

	if err != nil {
		log.Fatal(err)
	}

	xTurnImg, _, err = ebitenutil.NewImageFromFile("img/x_turn.png")

	if err != nil {
		log.Fatal(err)

	}

	oTurnImg, _, err = ebitenutil.NewImageFromFile("img/O_turn.png")

	if err != nil {
		log.Fatal(err)
	}

	xWinningImg, _, err = ebitenutil.NewImageFromFile("img/x_wins.png")

	if err != nil {
		log.Fatal(err)

	}

	oWinningImg, _, err = ebitenutil.NewImageFromFile("img/o_wins.png")

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
	boardHeightScaleFactor := float64(screenHeight) * 0.60 / float64(ticTacToeBoardImg.Bounds().Dy())
	boardOpts := &ebiten.DrawImageOptions{}

	boardOpts.GeoM.Scale(boardWidthScaleFactor, boardHeightScaleFactor)
	boardOpts.GeoM.Translate(0, float64(screenHeight)*0.40)

	screen.DrawImage(ticTacToeBoardImg, boardOpts)

	titleImgWidth := titleImg.Bounds().Dx()
	titleWidthScaleFactor := float64(screenWidth) / float64(titleImgWidth)
	titleHeightScaleFactor := float64(screenHeight) * 0.20 / float64(titleImg.Bounds().Dy())
	titleOpts := &ebiten.DrawImageOptions{}

	titleOpts.GeoM.Scale(titleWidthScaleFactor, titleHeightScaleFactor)

	screen.DrawImage(titleImg, titleOpts)

	drawInfo(screen)
}

func drawInfo(screen *ebiten.Image) {
	infoImgWidth := xTurnImg.Bounds().Dx()
	infoWidthScaleFactor := float64(screenWidth) / float64(infoImgWidth)
	infoHeightScaleFactor := float64(screenHeight) * 0.20 / float64(xTurnImg.Bounds().Dy())

	infoOpts := &ebiten.DrawImageOptions{}

	infoOpts.GeoM.Scale(infoWidthScaleFactor, infoHeightScaleFactor)
	infoOpts.GeoM.Translate(0, float64(screenHeight)*0.20)

	if winner == "X" {
		screen.DrawImage(xWinningImg, infoOpts)

		return
	}
	if winner == "O" {
		screen.DrawImage(oWinningImg, infoOpts)

		return
	}
	if isXTurn {
		screen.DrawImage(xTurnImg, infoOpts)

		return
	}

	screen.DrawImage(oTurnImg, infoOpts)
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

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	boardslotImg    *ebiten.Image
	xImg            *ebiten.Image
	oImg            *ebiten.Image
	titleImg        *ebiten.Image
	xTurnImg        *ebiten.Image
	oTurnImg        *ebiten.Image
	xWinningImg     *ebiten.Image
	oWinningImg     *ebiten.Image
	winner          string
	isXTurn         bool = true
	squareSelection int  = 0
	screenWidth     int  = 800
	screenHeight    int  = 600
)

func init() {
	var err error

	boardslotImg, _, err = ebitenutil.NewImageFromFile("img/boardslot.png")

	if err != nil {
		log.Fatal(err)
	}

	xImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_x.png")

	if err != nil {
		log.Fatal(err)

	}

	oImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_o.png")

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
	boardslotImgWidth := boardslotImg.Bounds().Dx()
	boardslotWidthScaleFactor := float64(screenWidth) / 3 / float64(boardslotImgWidth)
	boardslotHeightScaleFactor := float64(screenHeight) / 3 * 0.60 / float64(boardslotImg.Bounds().Dy())

	boardslot1Opts := &ebiten.DrawImageOptions{}
	boardslot1Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot1Opts.GeoM.Translate(0, float64(screenHeight)*0.40)
	screen.DrawImage(boardslotImg, boardslot1Opts)

	boardslot2Opts := &ebiten.DrawImageOptions{}
	boardslot2Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot2Opts.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.40)
	screen.DrawImage(boardslotImg, boardslot2Opts)

	boardslot3Opts := &ebiten.DrawImageOptions{}
	boardslot3Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot3Opts.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.40)
	screen.DrawImage(boardslotImg, boardslot3Opts)

	boardslot4Opts := &ebiten.DrawImageOptions{}
	boardslot4Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot4Opts.GeoM.Translate(0, float64(screenHeight)*0.6)
	screen.DrawImage(boardslotImg, boardslot4Opts)

	boardslot5Opts := &ebiten.DrawImageOptions{}
	boardslot5Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot5Opts.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.6)
	screen.DrawImage(boardslotImg, boardslot5Opts)

	boardslot6Opts := &ebiten.DrawImageOptions{}
	boardslot6Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot6Opts.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.6)
	screen.DrawImage(boardslotImg, boardslot6Opts)

	boardslot7Opts := &ebiten.DrawImageOptions{}
	boardslot7Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot7Opts.GeoM.Translate(0, float64(screenHeight)*0.8)
	screen.DrawImage(boardslotImg, boardslot7Opts)

	boardslot8Opts := &ebiten.DrawImageOptions{}
	boardslot8Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot8Opts.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.8)
	screen.DrawImage(boardslotImg, boardslot8Opts)

	boardslot9Opts := &ebiten.DrawImageOptions{}
	boardslot9Opts.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslot9Opts.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.8)
	screen.DrawImage(boardslotImg, boardslot9Opts)

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

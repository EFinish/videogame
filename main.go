package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type IsXOrOEnum int

const (
	IsX IsXOrOEnum = iota
	IsO
	Niether
)

type Slot struct {
	slotNumber   int
	isSelected   bool
	displayError bool
	isXOrO       IsXOrOEnum
}

var (
	boardslotImg          *ebiten.Image
	boardSlotSelectionImg *ebiten.Image
	xSlotImg              *ebiten.Image
	xSlotSelectionImg     *ebiten.Image
	xSlotSelectionErrImg  *ebiten.Image
	oSlotImg              *ebiten.Image
	oSlotSelectionImg     *ebiten.Image
	oSlotSelectionErrImg  *ebiten.Image
	titleImg              *ebiten.Image
	xTurnImg              *ebiten.Image
	oTurnImg              *ebiten.Image
	xWinningImg           *ebiten.Image
	oWinningImg           *ebiten.Image
	board                 [9]Slot
	winner                string
	whosTurn              IsXOrOEnum
	screenWidth           int = 800
	screenHeight          int = 600
)

func init() {
	var err error

	boardslotImg, _, err = ebitenutil.NewImageFromFile("img/boardslot.png")
	if err != nil {
		log.Fatal(err)
	}

	boardSlotSelectionImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_selection.png")
	if err != nil {
		log.Fatal(err)
	}

	xSlotImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_x.png")
	if err != nil {
		log.Fatal(err)

	}

	xSlotSelectionImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_selection_x.png")
	if err != nil {
		log.Fatal(err)
	}

	xSlotSelectionErrImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_selection_err_x.png")
	if err != nil {
		log.Fatal(err)
	}

	oSlotImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_o.png")
	if err != nil {
		log.Fatal(err)
	}

	oSlotSelectionImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_selection_o.png")
	if err != nil {
		log.Fatal(err)
	}

	oSlotSelectionErrImg, _, err = ebitenutil.NewImageFromFile("img/boardslot_selection_err_o.png")
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

	board = [9]Slot{
		{slotNumber: 0, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 1, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 2, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 3, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 4, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 5, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 6, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 7, isSelected: false, displayError: false, isXOrO: Niether},
		{slotNumber: 8, isSelected: false, displayError: false, isXOrO: Niether},
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	boardslotImgWidth := boardslotImg.Bounds().Dx()
	boardslotWidthScaleFactor := float64(screenWidth) / 3 / float64(boardslotImgWidth)
	boardslotHeightScaleFactor := float64(screenHeight) / 3 * 0.6 / float64(boardslotImg.Bounds().Dy())

	boardslotOpts0 := &ebiten.DrawImageOptions{}
	boardslotOpts0.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts0.GeoM.Translate(0, float64(screenHeight)*0.4)
	slotImg0 := getImgForSlot(0)
	screen.DrawImage(slotImg0, boardslotOpts0)

	boardslotOpts1 := &ebiten.DrawImageOptions{}
	boardslotOpts1.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts1.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.4)
	slotImg1 := getImgForSlot(1)
	screen.DrawImage(slotImg1, boardslotOpts1)

	boardslotOpts2 := &ebiten.DrawImageOptions{}
	boardslotOpts2.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts2.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.4)
	slotImg2 := getImgForSlot(2)
	screen.DrawImage(slotImg2, boardslotOpts2)

	boardslotOpts3 := &ebiten.DrawImageOptions{}
	boardslotOpts3.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts3.GeoM.Translate(0, float64(screenHeight)*0.6)
	slotImg3 := getImgForSlot(3)
	screen.DrawImage(slotImg3, boardslotOpts3)

	boardslotOpts4 := &ebiten.DrawImageOptions{}
	boardslotOpts4.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts4.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.6)
	slotImg4 := getImgForSlot(4)
	screen.DrawImage(slotImg4, boardslotOpts4)

	boardslotOpts5 := &ebiten.DrawImageOptions{}
	boardslotOpts5.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts5.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.6)
	slotImg5 := getImgForSlot(5)
	screen.DrawImage(slotImg5, boardslotOpts5)

	boardslotOpts6 := &ebiten.DrawImageOptions{}
	boardslotOpts6.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts6.GeoM.Translate(0, float64(screenHeight)*0.8)
	slotImg6 := getImgForSlot(6)
	screen.DrawImage(slotImg6, boardslotOpts6)

	boardslotOpts7 := &ebiten.DrawImageOptions{}
	boardslotOpts7.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts7.GeoM.Translate(float64(screenWidth)*0.33, float64(screenHeight)*0.8)
	slotImg7 := getImgForSlot(7)
	screen.DrawImage(slotImg7, boardslotOpts7)

	boardslotOpts8 := &ebiten.DrawImageOptions{}
	boardslotOpts8.GeoM.Scale(boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	boardslotOpts8.GeoM.Translate(float64(screenWidth)*0.66, float64(screenHeight)*0.8)
	slotImg8 := getImgForSlot(8)
	screen.DrawImage(slotImg8, boardslotOpts8)

	titleImgWidth := titleImg.Bounds().Dx()
	titleWidthScaleFactor := float64(screenWidth) / float64(titleImgWidth)
	titleHeightScaleFactor := float64(screenHeight) * 0.20 / float64(titleImg.Bounds().Dy())
	titleOpts := &ebiten.DrawImageOptions{}

	titleOpts.GeoM.Scale(titleWidthScaleFactor, titleHeightScaleFactor)

	screen.DrawImage(titleImg, titleOpts)

	drawInfo(screen)
}

func getImgForSlot(slotNumber int) *ebiten.Image {
	slot := board[slotNumber]

	if slot.isSelected {
		switch slot.isXOrO {
		case IsX:
			if slot.displayError {
				return xSlotSelectionErrImg
			}
			return xSlotSelectionImg
		case IsO:
			if slot.displayError {
				return oSlotSelectionErrImg
			}
			return oSlotSelectionImg
		case Niether:
			return boardSlotSelectionImg
		}

	}

	switch slot.isXOrO {
	case IsX:
		return xSlotImg
	case IsO:
		return oSlotImg
	}

	return boardslotImg
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
	if whosTurn == IsX {
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

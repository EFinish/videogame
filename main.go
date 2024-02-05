package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type IsXOrOEnum int

const (
	IsX IsXOrOEnum = iota
	IsO
	Niether
)

type PossibleDirections struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Slot struct {
	slotNumber         int
	isSelected         bool
	displayError       bool
	isXOrO             IsXOrOEnum
	possibleDirections PossibleDirections
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
		{
			slotNumber: 0, isSelected: true, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: false, Down: true, Left: false, Right: true},
		},
		{
			slotNumber: 1, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: false, Down: true, Left: true, Right: true},
		},
		{
			slotNumber: 2, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: false, Down: true, Left: true, Right: false},
		},
		{
			slotNumber: 3, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: true, Left: false, Right: true},
		},
		{
			slotNumber: 4, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: true, Left: true, Right: true},
		},
		{
			slotNumber: 5, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: true, Left: true, Right: false},
		},
		{
			slotNumber: 6, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: false, Left: false, Right: true},
		},
		{
			slotNumber: 7, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: false, Left: true, Right: true},
		},
		{
			slotNumber: 8, isSelected: false, displayError: false, isXOrO: Niether,
			possibleDirections: PossibleDirections{Up: true, Down: false, Left: true, Right: false},
		},
	}
}

type Game struct{}

func (g *Game) Update() error {
	if repeatingKeyPressed(ebiten.KeyA) || repeatingKeyPressed(ebiten.KeyLeft) {
		moveIfPossible(Left)
	}

	if repeatingKeyPressed(ebiten.KeyS) || repeatingKeyPressed(ebiten.KeyDown) {
		moveIfPossible(Down)
	}

	if repeatingKeyPressed(ebiten.KeyD) || repeatingKeyPressed(ebiten.KeyRight) {
		moveIfPossible(Right)
	}

	if repeatingKeyPressed(ebiten.KeyW) || repeatingKeyPressed(ebiten.KeyUp) {
		moveIfPossible(Up)
	}

	return nil
}

func repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)

	return d == 1
}

func moveIfPossible(direction Direction) {
	currentSlotNumber := getSelectedSlotNumber()
	currentSlot := &board[currentSlotNumber]

	switch direction {
	case Up:
		if currentSlot.possibleDirections.Up {
			currentSlot.isSelected = false
			board[currentSlotNumber-3].isSelected = true
		}
	case Down:
		if currentSlot.possibleDirections.Down {
			currentSlot.isSelected = false
			board[currentSlotNumber+3].isSelected = true
		}
	case Left:
		if currentSlot.possibleDirections.Left {
			currentSlot.isSelected = false
			board[currentSlotNumber-1].isSelected = true
		}
	case Right:
		if currentSlot.possibleDirections.Right {
			currentSlot.isSelected = false
			board[currentSlotNumber+1].isSelected = true
		}
	}
}

func getSelectedSlotNumber() int {
	for i := 0; i < len(board); i++ {
		if board[i].isSelected {
			return board[i].slotNumber
		}
	}

	log.Fatal("No slot is selected")
	return 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawInfo(screen)
	drawTitle(screen)

	boardslotImgWidth := boardslotImg.Bounds().Dx()
	boardslotWidthScaleFactor := float64(screenWidth) / 3 / float64(boardslotImgWidth)
	boardslotHeightScaleFactor := float64(screenHeight) / 3 * 0.6 / float64(boardslotImg.Bounds().Dy())

	drawImgForSlot(0, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(1, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(2, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(3, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(4, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(5, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(6, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(7, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)
	drawImgForSlot(8, screen, boardslotWidthScaleFactor, boardslotHeightScaleFactor)

}

func drawTitle(screen *ebiten.Image) {
	titleImgWidth := titleImg.Bounds().Dx()
	titleWidthScaleFactor := float64(screenWidth) / float64(titleImgWidth)
	titleHeightScaleFactor := float64(screenHeight) * 0.20 / float64(titleImg.Bounds().Dy())
	titleOpts := &ebiten.DrawImageOptions{}
	titleOpts.GeoM.Scale(titleWidthScaleFactor, titleHeightScaleFactor)

	screen.DrawImage(titleImg, titleOpts)
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

func drawImgForSlot(slotNumber int, screen *ebiten.Image, standardSlotWidth float64, standardSlotHeight float64) {
	x, y := 0.0, 0.0

	switch slotNumber {
	case 0:
		x, y = 0.0, float64(screenHeight)*0.4
	case 1:
		x, y = float64(screenWidth)*0.33, float64(screenHeight)*0.4
	case 2:
		x, y = float64(screenWidth)*0.66, float64(screenHeight)*0.4
	case 3:
		x, y = 0.0, float64(screenHeight)*0.6
	case 4:
		x, y = float64(screenWidth)*0.33, float64(screenHeight)*0.6
	case 5:
		x, y = float64(screenWidth)*0.66, float64(screenHeight)*0.6
	case 6:
		x, y = 0.0, float64(screenHeight)*0.8
	case 7:
		x, y = float64(screenWidth)*0.33, float64(screenHeight)*0.8
	case 8:
		x, y = float64(screenWidth)*0.66, float64(screenHeight)*0.8
	}

	boardslotOpts8 := &ebiten.DrawImageOptions{}
	boardslotOpts8.GeoM.Scale(standardSlotWidth, standardSlotHeight)
	boardslotOpts8.GeoM.Translate(x, y)
	slotImg8 := getImgForSlot(slotNumber)

	screen.DrawImage(slotImg8, boardslotOpts8)
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

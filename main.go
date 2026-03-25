package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/Kwynto/mech-exp-2/internal/defective"
	"github.com/Kwynto/mech-exp-2/internal/intypes"
)

const FILE_NAME = "data.txt"
const TEXT_MSG_1 = "Запись добалена"

var SlStGames []intypes.TStGame

var (
	slMainPremiumNumber intypes.TPremiumNumber
	slMainRiskNumbers   intypes.TRiskNumbers
)

func fileRead(sName string) (string, error) {
	bRead, err := os.ReadFile(sName)
	if err != nil {
		return "", err
	}
	return string(bRead), nil
}

func convInsideFormat(sData string) []intypes.TStGame {
	var SlStGamesTemp []intypes.TStGame

	slList := strings.Split(sData, "\n")

	for _, sLine := range slList {
		slGame := strings.Split(sLine, "|")

		if len(slGame) == 1 {
			break
		}

		slSWins := strings.Split(slGame[1], "-")
		slSWrongs := strings.Split(slGame[2], "-")

		var slWinNums []int
		var slWrongNums []int

		for _, sWin := range slSWins {
			iNum, err := strconv.Atoi(sWin)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWinNums = append(slWinNums, iNum)
		}

		for _, sWrong := range slSWrongs {
			iNum, err := strconv.Atoi(sWrong)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWrongNums = append(slWrongNums, iNum)
		}

		iGame, err := strconv.Atoi(slGame[0])
		if err != nil {
			fmt.Println("Conversion failed:", err)
		}

		stGameTemp := intypes.TStGame{
			Game:  iGame,
			Wins:  slWinNums,
			Wrong: slWrongNums,
		}

		SlStGamesTemp = append(SlStGamesTemp, stGameTemp)
	}

	return SlStGamesTemp
}

func readBase(sName string) {
	sData, err := fileRead(sName)
	if err != nil {
		return
	}

	clear(SlStGames)
	SlStGames = convInsideFormat(sData)
}

func setStartInterface() {
	iconApp, _ = fyne.LoadResourceFromPath("Icon.png")

	mainWindow.Resize(fyne.NewSize(400, 100))
	mainWindow.SetFixedSize(true)
	mainWindow.CenterOnScreen()
	mainWindow.SetIcon(iconApp)
	mainWindow.SetContent(container.NewVBox(
		btnShowDB,
		btnAddGame,
		btnDefAnalize,
	))
	mainWindow.SetOnClosed(func() {
		fyneApp.Quit()
	})

	showDBWindow.Resize(fyne.NewSize(890, 500))
	showDBWindow.CenterOnScreen()
	showDBWindow.SetIcon(iconApp)
	showDBTable.SetColumnWidth(0, 50)
	showDBTable.SetColumnWidth(1, 735)
	showDBTable.SetColumnWidth(2, 85)

	showDBWindow.SetContent(
		showDBTable,
	)

	showDBWindow.SetCloseIntercept(func() {
		showDBWindow.Hide()
		mainWindow.Show()
	})

	addGameWindow.Resize(fyne.NewSize(400, 150))
	addGameWindow.SetFixedSize(true)
	addGameWindow.CenterOnScreen()
	addGameWindow.SetIcon(iconApp)
	entAddGame.SetPlaceHolder("Тираж")
	entAddWins.SetPlaceHolder("Выпали")
	entAddWrongs.SetPlaceHolder("НЕ выпали")

	addGameWindow.SetContent(container.NewVBox(
		entAddGame,
		entAddWins,
		entAddWrongs,
		sepAddGame,
		btnAddGameOK,
	))

	addGameWindow.SetCloseIntercept(func() {
		addGameWindow.Hide()
		mainWindow.Show()
	})

	defAnalizeWindow.Resize(fyne.NewSize(600, 200))
	defAnalizeWindow.SetFixedSize(true)
	defAnalizeWindow.CenterOnScreen()
	defAnalizeWindow.SetIcon(iconApp)

	defAnLb1.Show()
	defAnEnt1.Show()
	defAnLb3.Hide()
	defAnEnt3.Hide()
	defAnEnt4.Hide()

	defAnEnt1.SetText("0")
	defAnEnt2.SetText("2")
	defAnEnt3.SetText("0")
	defAnEnt4.SetText("0")

	defAnalizeWindow.SetContent(
		container.NewCenter(
			container.NewVBox(
				defAnTopMenu,
			),
		),
	)

	defAnalizeWindow.SetCloseIntercept(fDefAnClose)

	showDefAnGoodTable.SetColumnWidth(0, 50)
	showDefAnGoodTable.SetColumnWidth(1, 200)
	showDefAnGoodTable.SetColumnWidth(2, 50)

	showDefAnBedTable.SetColumnWidth(0, 50)
	showDefAnBedTable.SetColumnWidth(1, 200)
	showDefAnBedTable.SetColumnWidth(2, 50)

	defAnGoodWindow.Resize(fyne.NewSize(350, 200))
	defAnBedWindow.Resize(fyne.NewSize(350, 200))

	defAnGoodWindow.SetContent(
		showDefAnGoodTable,
	)

	defAnBedWindow.SetContent(
		showDefAnBedTable,
	)

	defAnBedWindow.SetCloseIntercept(fDefAnClose)
	defAnGoodWindow.SetCloseIntercept(fDefAnClose)

	defAnEnt1.OnChanged = func(s string) {
		iGame, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Ошибка конвертации")
		}
		iBorder := defective.PreAnalize(SlStGames, iGame)
		iSBorer := fmt.Sprint(iBorder)
		defAnEnt2.SetText(iSBorer)
	}
}

func main() {
	readBase(FILE_NAME)

	setStartInterface()

	mainWindow.Show()
	fyneApp.Run()
}

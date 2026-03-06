package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Kwynto/mech-exp-2/internal/intypes"
)

const FILE_NAME = "data.txt"
const TEXT_SEPARATOR_1 = "               "
const TEXT_SEPARATOR_2 = "Запись добалена"

var SlStGames []intypes.TStGame

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

// Приложение
var fyneApp fyne.App = app.New()

// Главное окно
var mainWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор")
var btnShowDB *widget.Button = widget.NewButton("Показать базу тиражей", showGameDB)
var btnAddGame *widget.Button = widget.NewButton("Добавить данные тиража", addGame)
var btnDefAnalize *widget.Button = widget.NewButton("Дефектный анализ", makeDefAnalize)

// Окно показа базы
var showDBWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор - архив тиражей")

func showGameDB() {
	readBase(FILE_NAME)

	mainWindow.Hide()
	showDBWindow.Show()
	fmt.Println("Типа показали базу.")
}

// Окно добавления тиража
var addGameWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - добавить тираж")

var entAddGame *widget.Entry = widget.NewEntry()
var entAddWins *widget.Entry = widget.NewEntry()
var entAddWrongs *widget.Entry = widget.NewEntry()

// var sepAddGame *widget.Separator = widget.NewSeparator()
var lbAddGameOK *widget.Label = widget.NewLabel(TEXT_SEPARATOR_1)
var btnAddGameOK *widget.Button = widget.NewButton("Добаввить запись тиража", fAddGameOK)

func fAddGameOK() {
	defer readBase(FILE_NAME)

	var sNum, sWins, sWrongs string

	sNum = entAddGame.Text
	sWins = entAddWins.Text
	sWrongs = entAddWrongs.Text

	entAddGame.SetText("")
	entAddWins.SetText("")
	entAddWrongs.SetText("")

	sRecord := fmt.Sprintf("%s|%s|%s\n", sNum, sWins, sWrongs)
	fFileName, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer fFileName.Close()

	if _, err := fFileName.WriteString(sRecord); err != nil {
		return
	}

	lbAddGameOK.SetText(TEXT_SEPARATOR_2)
	go func() {
		time.Sleep(time.Second * 2)
		fyne.Do(func() {
			lbAddGameOK.SetText(TEXT_SEPARATOR_1)
		})
	}()
}

func addGame() {
	mainWindow.Hide()
	addGameWindow.Show()
}

// Окно дефектного анализа
var defAnalizeWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор - дефектный анализ")

func makeDefAnalize() {
	mainWindow.Hide()
	defAnalizeWindow.Show()
	fmt.Println("Типа сделали дефектный анализ.")
}

func setStartInterface() {
	mainWindow.Resize(fyne.NewSize(400, 100))
	mainWindow.SetFixedSize(true)
	mainWindow.CenterOnScreen()
	mainWindow.SetContent(container.NewVBox(
		btnShowDB,
		btnAddGame,
		btnDefAnalize,
	))
	mainWindow.SetOnClosed(func() {
		fyneApp.Quit()
	})

	showDBWindow.Resize(fyne.NewSize(400, 300))
	showDBWindow.CenterOnScreen()

	showDBWindow.SetCloseIntercept(func() {
		showDBWindow.Hide()
		mainWindow.Show()
	})

	addGameWindow.Resize(fyne.NewSize(400, 150))
	addGameWindow.SetFixedSize(true)
	addGameWindow.CenterOnScreen()

	entAddGame.SetPlaceHolder("Тираж")
	entAddWins.SetPlaceHolder("Выпали")
	entAddWrongs.SetPlaceHolder("НЕ выпали")

	addGameWindow.SetContent(container.NewVBox(
		entAddGame,
		entAddWins,
		entAddWrongs,
		// sepAddGame,
		lbAddGameOK,
		btnAddGameOK,
	))

	addGameWindow.SetCloseIntercept(func() {
		addGameWindow.Hide()
		mainWindow.Show()
	})

	defAnalizeWindow.Resize(fyne.NewSize(400, 400))
	defAnalizeWindow.CenterOnScreen()
	defAnalizeWindow.SetCloseIntercept(func() {
		defAnalizeWindow.Hide()
		mainWindow.Show()
	})
}

func main() {
	readBase(FILE_NAME)

	setStartInterface()

	mainWindow.Show()
	fyneApp.Run()
}

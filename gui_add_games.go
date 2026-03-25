package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Окно добавления тиража
var addGameWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - добавить тираж")

var entAddGame *widget.Entry = widget.NewEntry()
var entAddWins *widget.Entry = widget.NewEntry()
var entAddWrongs *widget.Entry = widget.NewEntry()
var sepAddGame *widget.Separator = widget.NewSeparator()
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

	dialog.ShowInformation(
		" ",
		TEXT_MSG_1,
		addGameWindow,
	)
}

func addGame() {
	mainWindow.Hide()
	addGameWindow.Show()
}

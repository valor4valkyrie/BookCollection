package gui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartGui() {
	a := app.New()
	w := a.NewWindow("Book Collection")

	scanLabel := widget.NewLabel("Scan UPC")
	scanLabel.TextStyle.Bold = true
	scanLabel.Alignment = fyne.TextAlignCenter

	input := widget.NewEntry()
	input.MultiLine = true
	input.SetPlaceHolder("Enter UPC")

	addBookButton := widget.NewButton("Add Book", func() {
		input.SetText("")
	})

	closeButton := widget.NewButton("Close", func() { os.Exit(0) })

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Print(k.Name)
	})

	w.SetContent(container.NewVBox(
		scanLabel,
		input,
		addBookButton,
		closeButton,
	))

	w.Resize(fyne.Size{Width: 400, Height: 300})

	w.SetOnClosed(func() {
		os.Exit(0)
	})

	w.ShowAndRun()
}

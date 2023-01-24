package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"mariuzzo.com/chronometer/chronometer"
)

var ticker *time.Ticker = time.NewTicker(time.Second / 20)

func main() {
	app := app.New()
	window := buildMainWindow(app)
	window.ShowAndRun()
}

func buildMainWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Chronometer")

	digits := canvas.NewText(chronometer.Ellapsed(), color.White)
	digits.Alignment = fyne.TextAlignCenter
	digitsLayout := container.New(layout.NewMaxLayout(), digits)

	var startStopButton *widget.Button
	var resetButton *widget.Button

	startStopButton = widget.NewButton("Start", func() {
		switch chronometer.Status() {
		case chronometer.Idle:
			chronometer.Start()
			startStopButton.SetText("Pause")
		case chronometer.Stopped:
			chronometer.Resume()
			startStopButton.SetText("Pause")
		case chronometer.Running:
			chronometer.Stop()
			startStopButton.SetText("Resume")
		}
	})

	resetButton = widget.NewButton("Reset", func() {
		chronometer.Reset()
		digits.Text = chronometer.Ellapsed()
		digits.Refresh()
		startStopButton.SetText("Start")
	})

	controlsContainer := container.NewGridWithColumns(2, startStopButton, resetButton)

	mainContainer := container.NewGridWithRows(2, digitsLayout, controlsContainer)
	window.SetContent(mainContainer)

	go resizeDigits(digitsLayout, digits)
	go onResize(mainContainer, func() {
		resizeDigits(digitsLayout, digits)
	})
	go onTick(func() {
		if chronometer.Status() == chronometer.Running {
			digits.Text = chronometer.Ellapsed()
			digits.Refresh()
		}
	})

	return window
}

func onTick(f func()) {
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}

func onResize(container *fyne.Container, f func()) {
	w := container.Size().Width
	h := container.Size().Height

	onTick(func() {
		if w != container.Size().Width || h != container.Size().Height {
			w = container.Size().Width
			h = container.Size().Height
			go f()
		}
	})
}

func resizeDigits(container *fyne.Container, text *canvas.Text) {
	text.TextSize = container.Size().Height / 2
}

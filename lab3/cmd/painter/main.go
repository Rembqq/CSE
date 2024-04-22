package main

import (
	"net/http"

	"github.com/Rembqq/CSE/painter"
	"github.com/Rembqq/CSE/painter/lang"
	"github.com/Rembqq/CSE/ui"
)

func main() {
	var (
		pv ui.Visualizer // Візуалізатор створює вікно та малює у ньому.

		// Потрібні для частини 2.
		opLoop painter.Loop // Цикл обробки команд.
		parser lang.Parser  // Парсер команд.
	)

	//pv.Debug = true
	pv.WindowWidth = 800  // ширина вікна, пікселів
	pv.WindowHeight = 800 // висота вікна, пікселів
	pv.Title = "Simple painter"

	pv.OnScreenReady = opLoop.Start
	opLoop.Receiver = &pv

	go func() {
		http.Handle("/", lang.HttpHandler(&opLoop, &parser))
		_ = http.ListenAndServe("localhost:17000", nil)
	}()

	pv.Main()
	opLoop.StopAndWait()
}

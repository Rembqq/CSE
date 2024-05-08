package ui

import (
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Visualizer struct {
	Title         string
	WindowWidth   int
	WindowHeight  int
	Debug         bool
	OnScreenReady func(s screen.Screen)

	w    screen.Window
	tx   chan screen.Texture
	done chan struct{}

	mousePos mouse.Button

	sz  size.Event
	pos image.Rectangle
}

func (pw *Visualizer) Main() {
	pw.tx = make(chan screen.Texture)
	pw.done = make(chan struct{})
	pw.pos.Max.X = 200
	pw.pos.Max.Y = 200
	driver.Main(pw.run)
}

func (pw *Visualizer) Update(t screen.Texture) {
	pw.tx <- t
}

func (pw *Visualizer) run(s screen.Screen) {
	w, err := s.NewWindow(&screen.NewWindowOptions{
		Width:  pw.WindowWidth,
		Height: pw.WindowHeight,
		Title:  pw.Title,
	})
	if err != nil {
		log.Fatal("Failed to initialize the app window:", err)
	}
	defer func() {
		w.Release()
		close(pw.done)
	}()

	if pw.OnScreenReady != nil {
		pw.OnScreenReady(s)
	}

	pw.w = w

	events := make(chan any)
	go func() {
		for {
			e := w.NextEvent()
			if pw.Debug {
				log.Printf("new event: %v", e)
			}
			if detectTerminate(e) {
				close(events)
				break
			}
			events <- e
		}
	}()

	var t screen.Texture

	for {
		select {
		case e, ok := <-events:
			if !ok {
				return
			}
			pw.handleEvent(e, t)

		case t = <-pw.tx:
			w.Send(paint.Event{})
		}
	}
}

func detectTerminate(e any) bool {
	switch e := e.(type) {
	case lifecycle.Event:
		if e.To == lifecycle.StageDead {
			return true // Window destroy initiated.
		}
	case key.Event:
		if e.Code == key.CodeEscape {
			return true // Esc pressed.
		}
	}
	return false
}

func (pw *Visualizer) handleEvent(e any, t screen.Texture) {
	switch e := e.(type) {

	case size.Event: // Оновлення даних про розмір вікна.
		pw.sz = e

	case error:
		log.Printf("ERROR: %s", e)

	case mouse.Event:
		if t == nil {
			if e.Direction == mouse.DirPress {
				if e.Button == mouse.ButtonLeft {
					x, y := 1-(800-e.X)/800, 1-(800-e.Y)/800 // за формулою знаходження відсотків
					pw.w.Fill(pw.sz.Bounds(), color.RGBA{R: 0, G: 200, B: 0, A: 200}, screen.Src)
					pw.XFigureDraw(x, y)
				}
			}
		}

	case paint.Event:
		// Малювання контенту вікна.
		if t == nil {
			pw.drawDefaultUI()
		} else {
			// Використання текстури отриманої через виклик Update.
			pw.w.Scale(pw.sz.Bounds(), t, t.Bounds(), draw.Src, nil)
		}
		pw.w.Publish()
	}
}

func (pw *Visualizer) drawDefaultUI() {
	pw.w.Fill(pw.sz.Bounds(), color.RGBA{R: 0, G: 200, B: 0, A: 200}, screen.Src)
	pw.XFigureDraw(0.5, 0.5) // малює хрестик по центу екрану
}

func (pw *Visualizer) XFigureDraw(x, y float32) { // малює хрестик по координатам що вказують на його центр
	cordinateX := int(x * 800)
	cordinateY := int(y * 800)
	xFigureWidth, xFigureHeight := 170, 70 // половини розмірів довжини та висоти прямокутника
	startX1, startY1, endX1, endY1 := cordinateX-xFigureWidth, cordinateY-xFigureHeight, cordinateX+xFigureWidth, cordinateY+xFigureHeight
	pw.RectFigureDraw(startX1, startY1, endX1, endY1, 200, 0, 0, 200)
	startX2, startY2, endX2, endY2 := cordinateX-xFigureHeight, cordinateY-xFigureWidth, cordinateX+xFigureHeight, cordinateY+xFigureWidth
	pw.RectFigureDraw(startX2, startY2, endX2, endY2, 200, 0, 0, 200)
}

func (pw *Visualizer) RectFigureDraw(x1, y1, x2, y2 int, r, g, b, a byte) {
	pw.pos.Min.X = x1 // лівий верхній кут, початкова координата Х
	pw.pos.Min.Y = y1 // лівий верхній кут, початкова координата У
	pw.pos.Max.X = x2 // правий нижній кут, кінцева координата Х
	pw.pos.Max.Y = y2 // правий нижній кут, кінцева координата У
	pw.w.Fill(pw.pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}

//for _, br := range imageutil.Border(pw.sz.Bounds(), 5) { // Малювання білої рамки.
//pw.w.Fill(br, color.White, draw.Src)
//}

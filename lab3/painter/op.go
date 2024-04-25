package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// малює червоний хрестик по координатам що вказують на його центр
func XFigureDraw(t screen.Texture /*, x, y float32*/) {
	cordinateX := int(0.5 * 800)
	cordinateY := int(0.5 * 800)
	xFigureWidth, xFigureHeight := 170, 70 // половини розмірів довжини та висоти прямокутника
	startX1, startY1, endX1, endY1 := cordinateX-xFigureWidth, cordinateY-xFigureHeight, cordinateX+xFigureWidth, cordinateY+xFigureHeight
	RectFigureDraw(t, startX1, startY1, endX1, endY1, 200, 0, 0, 200)
	startX2, startY2, endX2, endY2 := cordinateX-xFigureHeight, cordinateY-xFigureWidth, cordinateX+xFigureHeight, cordinateY+xFigureWidth
	RectFigureDraw(t, startX2, startY2, endX2, endY2, 200, 0, 0, 200)
}

func XFigureDraw2(t screen.Texture /*, x, y float32*/) {
	cordinateX := int(0.52 * 800)
	cordinateY := int(0.52 * 800)
	xFigureWidth, xFigureHeight := 170, 70 // половини розмірів довжини та висоти прямокутника
	startX1, startY1, endX1, endY1 := cordinateX-xFigureWidth, cordinateY-xFigureHeight, cordinateX+xFigureWidth, cordinateY+xFigureHeight
	RectFigureDraw(t, startX1, startY1, endX1, endY1, 200, 0, 0, 200)
	startX2, startY2, endX2, endY2 := cordinateX-xFigureHeight, cordinateY-xFigureWidth, cordinateX+xFigureHeight, cordinateY+xFigureWidth
	RectFigureDraw(t, startX2, startY2, endX2, endY2, 200, 0, 0, 200)
}

// малює чорний прямокутник по координатам що вказують
func BlackRect(t screen.Texture /*, x1, y1, x2, y2 int*/) {
	RectFigureDraw(t, 200, 200, 600, 600, 0, 0, 0, 0)
}

func RectFigureDraw(t screen.Texture, x1, y1, x2, y2 int, r, g, b, a byte) {
	var pos image.Rectangle
	pos.Min.X = x1 // лівий верхній кут, початкова координата Х
	pos.Min.Y = y1 // лівий верхній кут, початкова координата У
	pos.Max.X = x2 // правий нижній кут, кінцева координата Х
	pos.Max.Y = y2 // правий нижній кут, кінцева координата У
	t.Fill(pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}

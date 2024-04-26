package painter

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture, c Cord) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

type Cord struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

type CordinateList []Cord

func (ol OperationList) Do(t screen.Texture, c Cord) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t, c) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture, c Cord) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture, c Cord)

func (f OperationFunc) Do(t screen.Texture, c Cord) bool {
	f(t, c)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture, c Cord) {
	fmt.Println("white")
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture, c Cord) {
	fmt.Println("green")
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// малює червоний хрестик по координатам що вказують на його центр
func XFigureDraw(t screen.Texture, c Cord) {
	fmt.Println("figure", c.X1, c.Y1)
	cordinateX := int(c.X1 * 800)
	cordinateY := int(c.Y1 * 800)
	xFigureWidth, xFigureHeight := 170, 70 // половини розмірів довжини та висоти прямокутника
	startX1, startY1, endX1, endY1 := cordinateX-xFigureWidth, cordinateY-xFigureHeight, cordinateX+xFigureWidth, cordinateY+xFigureHeight
	RectFigureDraw(t, startX1, startY1, endX1, endY1, 200, 0, 0, 200)
	startX2, startY2, endX2, endY2 := cordinateX-xFigureHeight, cordinateY-xFigureWidth, cordinateX+xFigureHeight, cordinateY+xFigureWidth
	RectFigureDraw(t, startX2, startY2, endX2, endY2, 200, 0, 0, 200)
}

// малює чорний прямокутник по координатам що вказують
func BlackRect(t screen.Texture, c Cord) {
	fmt.Println("rect", c.X1, c.Y1, c.X2, c.Y2)
	RectFigureDraw(t, int(c.X1), int(c.Y1), int(c.X2), int(c.Y2), 0, 0, 0, 0)
}

func RectFigureDraw(t screen.Texture, x1, y1, x2, y2 int, r, g, b, a byte) {
	var pos image.Rectangle
	pos.Min.X = x1 // лівий верхній кут, початкова координата Х
	pos.Min.Y = y1 // лівий верхній кут, початкова координата У
	pos.Max.X = x2 // правий нижній кут, кінцева координата Х
	pos.Max.Y = y2 // правий нижній кут, кінцева координата У
	t.Fill(pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}

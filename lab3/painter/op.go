package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Do(t screen.Texture) (ready bool)
}

type Cord struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool {
	return true
}

type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

var CordList []Cord

func PushCord(c Cord) {
	CordList = append(CordList, c)
}

func PullCord() Cord {
	res := CordList[0]
	CordList[0] = Cord{X1: 0, Y1: 0, X2: 0, Y2: 0}
	CordList = CordList[1:]
	return res
}

func BlackFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

func WhiteFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.White, screen.Src)
}

func GreenFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

func XFigureDraw(t screen.Texture) {
	c := PullCord()
	cordinateX := int(c.X1 * 800)
	cordinateY := int(c.Y1 * 800)
	xFigureWidth, xFigureHeight := 170, 70
	startX1, startY1, endX1, endY1 := cordinateX-xFigureWidth, cordinateY-xFigureHeight, cordinateX+xFigureWidth, cordinateY+xFigureHeight
	RectFigureDraw(t, startX1, startY1, endX1, endY1, 200, 0, 0, 200)
	startX2, startY2, endX2, endY2 := cordinateX-xFigureHeight, cordinateY-xFigureWidth, cordinateX+xFigureHeight, cordinateY+xFigureWidth
	RectFigureDraw(t, startX2, startY2, endX2, endY2, 200, 0, 0, 200)
}

func BlackRect(t screen.Texture) {
	c := PullCord()
	RectFigureDraw(t, int(c.X1), int(c.Y1), int(c.X2), int(c.Y2), 0, 0, 0, 0)
}

func RectFigureDraw(t screen.Texture, x1, y1, x2, y2 int, r, g, b, a byte) {
	var pos image.Rectangle
	pos.Min.X = x1
	pos.Min.Y = y1
	pos.Max.X = x2
	pos.Max.Y = y2
	t.Fill(pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}

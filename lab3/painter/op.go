package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Do(t screen.Texture) (ready bool)
}

// структура координат
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
	PullCord()
	return true
}

type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// масив координат
var CordList []Cord

// запис в координат
func PushCord(c Cord) {
	CordList = append(CordList, c)
}

// отримання координат
func PullCord() Cord {
	res := CordList[0]
	CordList[0] = Cord{}
	CordList = CordList[1:]
	return res
}

// чорний фон
func BlackFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

// білий фон
func WhiteFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// зелений фон
func GreenFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// червона фігура х за центральними координатами
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

// чорний прамокутник
func BlackRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 0, 0, 0, 0)
}

// малює прямокутник
func RectFigureDraw(t screen.Texture, x1, y1, x2, y2 int, r, g, b, a byte) {
	var pos image.Rectangle
	pos.Min.X = x1
	pos.Min.Y = y1
	pos.Max.X = x2
	pos.Max.Y = y2
	t.Fill(pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}

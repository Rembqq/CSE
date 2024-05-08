package lang

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/Rembqq/CSE/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	// попередні операції фону
	PrevcolorWindow []painter.Operation
	// попередні операції прамокутника
	PrevrectBlack []painter.Operation
	// попередні операції фігури
	PrevfigureX []painter.Operation
	// попередні координати фону
	PrevcordWindow []painter.Cord
	// попередні координати прямокутника
	PrevcordRect []painter.Cord
	// попередні координати фігур
	PrevcordFigureX []painter.Cord
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, []painter.Cord, error) {
	// масив операцій
	var res []painter.Operation
	// масив операцій фону
	var colorWindow []painter.Operation
	// масив операцій прямокутника
	var rectBlack []painter.Operation
	// масив операцій фігур
	var figureX []painter.Operation
	// масив операцій оновлення
	var updatePicture []painter.Operation
	// масив операцій очищення
	var colorReset []painter.Operation
	// масив координат
	var cordIn []painter.Cord
	// масив координат фону
	var cordWindow []painter.Cord
	// масив координат прямокутника
	var cordRect []painter.Cord
	// масив координат фігур
	var cordFigureX []painter.Cord
	// масив координат оновлення
	var cordUpdate []painter.Cord
	// масив координат очищення
	var cordReset []painter.Cord

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		op, x1, y1, x2, y2 := p.ParseText(commandLine)
		switch op {
		case "white":
			colorWindow = append(colorWindow, painter.OperationFunc(painter.WhiteFill))
			cordWindow = append(cordWindow, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "green":
			colorWindow = append(colorWindow, painter.OperationFunc(painter.GreenFill))
			cordWindow = append(cordWindow, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "bgrect":
			rectBlack = append(rectBlack, painter.OperationFunc(painter.BlackRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "figure":
			figureX = append(figureX, painter.OperationFunc(painter.XFigureDraw))
			cordFigureX = append(cordFigureX, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "move":
			colorWindow = append(colorWindow, p.PrevcolorWindow...) // використання попередніх операцій
			rectBlack = append(rectBlack, p.PrevrectBlack...)
			figureX = append(figureX, p.PrevfigureX...)

			cordWindow = append(cordWindow, p.PrevcordWindow...) // використання попередніх координат
			cordRect = append(cordRect, p.PrevcordRect...)

			size := len(p.PrevcordFigureX)
			for i := 0; i < size; i++ { // зміна попередніх координат фігур
				cordFigureX = append(cordFigureX, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
			}
		case "reset":
			colorReset = append(colorReset, painter.OperationFunc(painter.BlackFill))
			cordReset = append(cordReset, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "update":
			updatePicture = append(updatePicture, painter.UpdateOp)
			cordUpdate = append(cordUpdate, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		}
	}

	res = append(res, colorWindow...)
	res = append(res, rectBlack...)
	res = append(res, figureX...)
	res = append(res, colorReset...)
	res = append(res, updatePicture...)

	cordIn = append(cordIn, cordWindow...)
	cordIn = append(cordIn, cordRect...)
	cordIn = append(cordIn, cordFigureX...)
	cordIn = append(cordIn, cordReset...)
	cordIn = append(cordIn, cordUpdate...)

	// зберігання попередніх операцій та координат
	if len(colorWindow) != 0 || len(rectBlack) != 0 || len(figureX) != 0 {
		p.PrevcolorWindow = colorWindow
		p.PrevrectBlack = rectBlack
		p.PrevfigureX = figureX

		p.PrevcordWindow = cordWindow
		p.PrevcordRect = cordRect
		p.PrevcordFigureX = cordFigureX
	}

	return res, cordIn, nil
}

func (p *Parser) ParseText(lineText string) (string, float64, float64, float64, float64) {
	command := strings.Fields(lineText)
	size := len(command)
	com := command[0]
	x1, y1, x2, y2 := float64(0), float64(0), float64(0), float64(0)
	if size >= 3 {
		if command[1] != "" {
			x1, _ = strconv.ParseFloat(command[1], 64)
		}
		if command[2] != "" {
			y1, _ = strconv.ParseFloat(command[2], 64)
		}
	}
	if size == 5 {
		if command[3] != "" {
			x2, _ = strconv.ParseFloat(command[3], 64)
		}
		if command[4] != "" {
			y2, _ = strconv.ParseFloat(command[4], 64)
		}
	}
	return com, x1, y1, x2, y2
}

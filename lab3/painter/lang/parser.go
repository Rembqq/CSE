package lang

import (
	"io"

	"github.com/Rembqq/CSE/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation           // масив операцій який передається на виконання
	var colorWindow []painter.Operation   // масив операцій  які перемальовують фон
	var figurePicture []painter.Operation // масив операцій які малюють фігури

	// TODO: Реалізувати парсинг команд.
	colorWindow = append(colorWindow, painter.OperationFunc(painter.WhiteFill))
	figurePicture = append(figurePicture, painter.OperationFunc(painter.BlackRect))
	figurePicture = append(figurePicture, painter.OperationFunc(painter.XFigureDraw))
	colorWindow = append(colorWindow, painter.OperationFunc(painter.GreenFill))
	figurePicture = append(figurePicture, painter.OperationFunc(painter.XFigureDraw2))
	res = append(res, colorWindow...)   // в масив який передається на виконнання, додаються операції для перемалювання фону
	res = append(res, figurePicture...) // потім в масив який передається на виконнання, додаються операції для малювання фігур
	res = append(res, painter.UpdateOp) // у кінець масиву передається операція оновлення текстури

	return res, nil
}

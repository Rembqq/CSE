package lang

import (
	"bufio"
	"io"
	"strings"

	"github.com/Rembqq/CSE/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation           // масив операцій який передається на виконання
	var colorWindow []painter.Operation   // масив операцій які перемальовують фон
	var figurePicture []painter.Operation // масив операцій які малюють фігури
	var updatePicture []painter.Operation // масив операцій які малюють фігури
	var resetPicture []painter.Operation

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		op := p.ParseText(commandLine) // parse the line to get Operation
		switch op {
		case "white":
			colorWindow = append(colorWindow, painter.OperationFunc(painter.WhiteFill))
		case "green":
			colorWindow = append(colorWindow, painter.OperationFunc(painter.GreenFill))
		case "bgrect":
			figurePicture = append(figurePicture, painter.OperationFunc(painter.BlackRect))
		case "figure":
			figurePicture = append(figurePicture, painter.OperationFunc(painter.XFigureDraw))
		case "move":
		case "reset":
			resetPicture = append(resetPicture, painter.OperationFunc(painter.BlackFill))
		case "update":
			updatePicture = append(updatePicture, painter.UpdateOp)
		}
	}

	res = append(res, colorWindow...)   // в масив який передається на виконнання, додаються операції для перемалювання фону
	res = append(res, figurePicture...) // потім в масив який передається на виконнання, додаються операції для малювання фігур
	res = append(res, resetPicture...)
	res = append(res, updatePicture...) // у кінець масиву передається операція оновлення текстури

	return res, nil
}

func (p *Parser) ParseText(lineText string) string {
	command := strings.Fields(lineText)
	com := command[0]
	return com
}

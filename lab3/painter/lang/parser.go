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
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, []painter.Cord, error) {
	var res []painter.Operation
	var colorWindow []painter.Operation
	var figurePicture []painter.Operation
	var updatePicture []painter.Operation

	var cordIn []painter.Cord
	var cordWindow []painter.Cord
	var cordFigure []painter.Cord
	var cordUpdate []painter.Cord

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
			figurePicture = append(figurePicture, painter.OperationFunc(painter.BlackRect))
			cordFigure = append(cordFigure, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "figure":
			figurePicture = append(figurePicture, painter.OperationFunc(painter.XFigureDraw))
			cordFigure = append(cordFigure, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "move":
		case "reset":
		case "update":
			updatePicture = append(updatePicture, painter.UpdateOp)
			cordUpdate = append(cordUpdate, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		}
	}

	res = append(res, colorWindow...)
	res = append(res, figurePicture...)
	res = append(res, updatePicture...)

	cordIn = append(cordIn, cordWindow...)
	cordIn = append(cordIn, cordFigure...)
	cordIn = append(cordIn, cordUpdate...)

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

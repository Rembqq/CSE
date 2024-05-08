package lang

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Rembqq/CSE/painter"
)

func TestParser(t *testing.T) {
	var p Parser

	for _, tc := range []struct {
		numt string
		cmd  string
		cord []painter.Cord
	}{
		{numt: "1", cmd: "update", cord: []painter.Cord{{X1: 0, Y1: 0, X2: 0, Y2: 0}}},
		{numt: "2", cmd: "bgrect 0.25 0.25 0.75 0.75\nfigure 0.5 0.5\nupdate", cord: []painter.Cord{{X1: 0.25, Y1: 0.25, X2: 0.75, Y2: 0.75}, {X1: 0.5, Y1: 0.5, X2: 0, Y2: 0}, {X1: 0, Y1: 0, X2: 0, Y2: 0}}},
	} {
		t.Run(tc.numt, func(t *testing.T) {
			in := strings.NewReader(tc.cmd)
			_, c, _ := p.Parse(in)
			for i := 0; i < len(c); i++ {
				if c[i] != tc.cord[i] {
					t.Errorf("Неправильний запис координат.")
					return
				} else {
					fmt.Println("Правильний запис", c[i], "=", tc.cord[i])
				}
			}

		})
	}
}

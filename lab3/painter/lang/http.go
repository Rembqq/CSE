package lang

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Rembqq/CSE/painter"
)

// HttpHandler конструює обробник HTTP запитів, який дані з запиту віддає у Parser, а потім відправляє отриманий список
// операцій у painter.Loop.
func HttpHandler(loop *painter.Loop, p *Parser) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var in io.Reader = r.Body
		if r.Method == http.MethodGet {
			in = strings.NewReader(r.URL.Query().Get("cmd"))
		}

		cmds, cords, err := p.Parse(in)
		if err != nil {
			log.Printf("Bad script: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(cords) != 0 {
			for i := 0; i <= len(cords)-1; i++ {
				loop.PostCord(cords[i])
			}
		}
		loop.Post(painter.OperationList(cmds))
		rw.WriteHeader(http.StatusOK)
	})
}

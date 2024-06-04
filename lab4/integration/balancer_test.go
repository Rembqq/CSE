package integration

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"gopkg.in/check.v1"

	"github.com/Rembqq/CSE/httptools"
	"github.com/Rembqq/CSE/signal"
)

const baseAddress = "http://localhost:8090" /*"http://balancer:8090"*/
var port = flag.Int("port", 8080, "server port")
var traceEnabled = flag.Bool("trace", true, "whether to include tracing information into responses")

var (
	i int = 0
)

func ImplemTestBalancer(t *testing.T) {
	if _, exists := os.LookupEnv("INTEGRATION_TEST"); !exists {
		t.Skip("Integration test is not enabled")
	}
	check.TestingT(t)
}

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func (s *MySuite) TestBalanser(c *check.C) {
	server := httptools.CreateServer(*port, main.h)
	server.Start()
	signal.WaitForTerminationSignal()

	frontend := httptools.CreateServer(*port, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientAddr := r.RemoteAddr
		serverIndex := main.hash(clientAddr) % len(main.serversPoolTrue)
		main.forward(main.serversPoolTrue[serverIndex], rw, r)
	}))

	log.Println("Starting load balancer...")
	log.Printf("Tracing support enabled: %t", *traceEnabled)
	frontend.Start()
	signal.WaitForTerminationSignal()

	client := new(http.Client)
	client.Timeout = 10 * time.Second

	for range time.Tick(1 * time.Second) {
		if i != 20 {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			if err != nil {
				c.Error(err)
			} else {
				c.Logf("response from [%s]", resp.Header.Get("lb-from"))
			}
			i++
		} else {
			return
		}
	}
}

func BenchmarkBalancer(b *testing.B) {
	client := new(http.Client)
	client.Timeout = 10 * time.Second
	b.Run("small", func(b *testing.B) {
		for i := 0; i <= 10; i++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			if err != nil {
				b.Error(err)
			} else {
				resp.Header.Get("lb-from")
			}
		}
	})
	b.Run("medium", func(b *testing.B) {
		for i := 0; i <= 50; i++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			if err != nil {
				b.Error(err)
			} else {
				resp.Header.Get("lb-from")
			}
		}
	})
	b.Run("large", func(b *testing.B) {
		for i := 0; i <= 100; i++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			if err != nil {
				b.Error(err)
			} else {
				resp.Header.Get("lb-from")
			}
		}
	})
}

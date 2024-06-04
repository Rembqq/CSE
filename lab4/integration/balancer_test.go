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

const baseAddress = "http://balancer:8090"

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
  client := new(http.Client)
	client.Timeout = 10 * time.Second
	for range time.Tick(2 * time.Second) {
		if i != 20 {
			  resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			  if err != nil {
				  b.Error(err)
			  } else {
				  resp.Header.Get("lb-from")
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

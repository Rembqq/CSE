package integration

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"gopkg.in/check.v1"
)

const baseAddress = "http://balancer:8090"
const baseAddress2 = "http://localhost:8090"

func TestBalancer(t *testing.T) {
	check.TestingT(t)
}

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func (s *MySuite) TestHealth(c *check.C) {
	if _, exists := os.LookupEnv("INTEGRATION_TEST"); !exists {
		c.Skip("Integration test is not enabled")
	}

	client := new(http.Client)
	client.Timeout = 10 * time.Second
	var i int = 0

	for range time.Tick(1 * time.Second) {
		if i < 20 {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
			if err != nil {
				c.Error(err)
			}
			c.Check(resp.Header.Get("lb-from")[:7], check.Not(check.Equals), "fwd 404")
			c.Logf("response from [%s]", resp.Header.Get("lb-from"))
			i++
		}
	}
}

func BenchmarkBalancer(b *testing.B) {
	client := new(http.Client)
	client.Timeout = 10 * time.Second
	b.Run("small", func(b *testing.B) {
		for k := 0; k < 10; k++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress2))
			if err != nil {
				b.Error(err)
			}
			resp.Header.Get("lb-from")
		}
	})
	b.Run("medium", func(b *testing.B) {
		for k := 0; k < 100; k++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress2))
			if err != nil {
				b.Error(err)
			}
			resp.Header.Get("lb-from")
		}
	})
	b.Run("large", func(b *testing.B) {
		for k := 0; k < 1000; k++ {
			resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress2))
			if err != nil {
				b.Error(err)
			}
			resp.Header.Get("lb-from")
		}
	})
}

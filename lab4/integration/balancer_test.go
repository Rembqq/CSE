package integration

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

const baseAddress = "http://balancer:8090"

var (
	client = http.Client{
		Timeout: 3 * time.Second,
	}
	i int = 0
)

func TestBalancer(t *testing.T) {
	if _, exists := os.LookupEnv("INTEGRATION_TEST"); !exists {
		t.Skip("Integration test is not enabled")
	}

	for range time.Tick(2 * time.Second) {
		if i != 100 {
			resp1, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/some-data", baseAddress), nil)
			if err != nil {
				t.Error(err)
			}

			resp2, err2 := client.Do(resp1)
			if err2 != nil {
				t.Error(err2)
				return
			}
			defer resp2.Body.Close()

			t.Logf("response from [%s]", resp2.Header.Get("lb-from"))
			i++
		} else {
			return
		}
	}
}

func BenchmarkBalancer(b *testing.B) {

}

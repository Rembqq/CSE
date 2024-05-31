package main

import (
	"github.com/Rembqq/CSE/httptools"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

//var (
//	serversPool = []string{
//		"server1:8080",
//		"server2:8080",
//		"server3:8080",
//	}
//)

type MySuite struct{}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestBalancer(c *C) {
	// TODO: Реалізуйте юніт-тест для балансувальникка.

	frontend := httptools.CreateServer(*port, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientAddr := r.RemoteAddr
		serverIndex := hash(clientAddr) % len(serversPool)
		expectedIndex := calculateExpectedIndex(clientAddr)
		c.Assert(serverIndex, Equals, expectedIndex)
	}))
	//frontend.Start()
}

func calculateExpectedIndex(clientAddr string) int {
	return hash(clientAddr) % len(serversPool)
}

func testScheme() string {
	if *https {
		return "https"
	}
	return "http"
}

//func testHealth(dst string) bool {
//	ctx, _ := context.WithTimeout(context.Background(), timeout)
//	req, _ := http.NewRequestWithContext(ctx, "GET",
//		fmt.Sprintf("%s://%s/health", testScheme(), dst), nil)
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return false
//	}
//	if resp.StatusCode != http.StatusOK {
//		return false
//	}
//	return true
//}
//
//func testForward(dst string, rw http.ResponseWriter, r *http.Request) error {
//	ctx, _ := context.WithTimeout(r.Context(), timeout)
//	fwdRequest := r.Clone(ctx)
//	fwdRequest.RequestURI = ""
//	fwdRequest.URL.Host = dst
//	fwdRequest.URL.Scheme = testScheme()
//	fwdRequest.Host = dst
//
//	resp, err := http.DefaultClient.Do(fwdRequest)
//	if err == nil {
//		for k, values := range resp.Header {
//			for _, value := range values {
//				rw.Header().Add(k, value)
//			}
//		}
//		if *traceEnabled {
//			rw.Header().Set("lb-from", dst)
//		}
//		log.Println("fwd", resp.StatusCode, resp.Request.URL)
//		rw.WriteHeader(resp.StatusCode)
//		defer resp.Body.Close()
//		_, err := io.Copy(rw, resp.Body)
//		if err != nil {
//			log.Printf("Failed to write response: %s", err)
//		}
//		return nil
//	} else {
//		log.Printf("Failed to get response from %s: %s", dst, err)
//		rw.WriteHeader(http.StatusServiceUnavailable)
//		return err
//	}
//}

var _ = Suite(&MySuite{})

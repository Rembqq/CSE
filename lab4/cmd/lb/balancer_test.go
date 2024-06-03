package main

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestHash(c *C) {
	input := "127.0.0.1:8080"
	expected := hash(input)
	c.Assert(expected, Equals, hash(input))

	// Checking for other inputs
	input2 := "192.168.1.1:8080"
	c.Assert(hash(input), Not(Equals), hash(input2))
}

func (s *MySuite) TestBalancer(c *C) {
	// Test different RemoteAddr for different serverIndex
	testCases := []struct {
		remoteAddr    string
		expectedIndex int
	}{
		{"127.0.0.1:8080", hash("127.0.0.1:8080") % len(serversPool)},
		{"192.168.1.1:8080", hash("192.168.1.1:8080") % len(serversPool)},
		{"10.0.0.1:8080", hash("10.0.0.1:8080") % len(serversPool)},
	}

	// tc - iterator of testCases (test case)
	for _, tc := range testCases {
		// Creates new GET request on http://localhost:8090
		req, err := http.NewRequest("GET", "http://localhost:8090", nil)
		// Checks for mistakes
		c.Assert(err, IsNil)
		// initializes client address as remoteAddr in tc
		req.RemoteAddr = tc.remoteAddr

		// http.ResponseWriter analogy in httptest package
		rw := httptest.NewRecorder()

		handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			clientAddr := r.RemoteAddr
			serverIndex := hash(clientAddr) % len(serversPool)
			c.Assert(serverIndex, Equals, tc.expectedIndex)
		})

		// imitation of processing an HTTP request by a server
		handler.ServeHTTP(rw, req)

		c.Assert(rw.Code, Equals, http.StatusOK)
		c.Assert(rw.Body.String(), Matches, ".*")
	}
}

func (s *MySuite) TestForward(c *C) {
	// Creating test server to imitate back-end server
	backendServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Back-end imitator message"))
	}))
	defer backendServer.Close()

	// Update serversPool to use a test server
	serversPool = []string{backendServer.Listener.Addr().String()}

	// Create HTTP-request for testing
	req, err := http.NewRequest("GET", "http://localhost:8090", nil)
	c.Assert(err, IsNil)

	// http.ResponseWriter analogy in httptest package
	rw := httptest.NewRecorder()

	// Call the forward function and check the result
	err = forward(serversPool[0], rw, req)
	c.Assert(err, IsNil)
	c.Assert(rw.Code, Equals, http.StatusOK)
	c.Assert(rw.Body.String(), Equals, "Back-end imitator message")
}

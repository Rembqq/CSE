package mainserv

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Handler1)
	return mux
}

func TestBalancer(t *testing.T) {
	check.TestingT(t)
}

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func (s *MySuite) TestHealth(c *check.C) {
	// створення масиву серверів
	serverTest := make([]*httptest.Server, 3)

	// створення  серверів
	serverTest[0] = httptest.NewServer(createMux())
	serverTest[1] = httptest.NewServer(createMux())
	serverTest[2] = httptest.NewServer(createMux())

	// перевірка здоров'я серверів
	for _, server := range serverTest {
		c.Check(health(server.URL[7:]), check.Equals, true)
	}

	// припиняємо роботу першого сервера
	server1 := serverTest[0]
	server2 := serverTest[1]
	server3 := serverTest[2]
	server1.Close()

	// перевірка здоров'я серверів
	c.Check(health(server1.URL[7:]), check.Equals, false)
	c.Check(health(server2.URL[7:]), check.Equals, true)
	c.Check(health(server3.URL[7:]), check.Equals, true)

	// припиняємо роботу всіх серверів
	for _, server := range serverTest {
		server.Close()
	}

	// перевірка здоров'я серверів
	c.Check(health(server1.URL[7:]), check.Equals, false)
	c.Check(health(server2.URL[7:]), check.Equals, false)
	c.Check(health(server3.URL[7:]), check.Equals, false)
}

func (s *MySuite) TestHash(c *check.C) {
	// перевіряємо роботу hash функції
	input := "127.0.0.1:8080"
	expected := hash(input)
	c.Check(expected, check.Equals, hash(input))

	// перевіряємо що хеш першого та другого користувача не співпадають
	input2 := "192.168.1.1:8080"
	c.Check(hash(input), check.Not(check.Equals), hash(input2))
}

func (s *MySuite) TestForward(c *check.C) {

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
	c.Check(err, check.IsNil)

	c.Assert(err, IsNil)

	// http.ResponseWriter analogy in httptest package
	rw := httptest.NewRecorder()

	// Call the forward function and check the result
	err = forward(serversPool[0], rw, req)
	c.Check(err, check.IsNil)
	c.Check(rw.Code, check.Equals, http.StatusOK)
	c.Check(rw.Body.String(), check.Equals, "Back-end imitator message")

	c.Assert(err, IsNil)
	c.Assert(rw.Code, Equals, http.StatusOK)
	c.Assert(rw.Body.String(), Equals, "Back-end imitator message")
}

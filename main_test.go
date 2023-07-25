package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Dummy server for testing
type dummyServer struct {
	addr    string
	isAlive bool
}

func (s *dummyServer) Address() string { return s.addr }

func (s *dummyServer) IsAlive() bool { return s.isAlive }

func (s *dummyServer) Serve(rw http.ResponseWriter, req *http.Request) {}

func TestGetNextAvailableServer(t *testing.T) {
	servers := []Server{
		&dummyServer{addr: "server1", isAlive: true},
		&dummyServer{addr: "server2", isAlive: false},
		&dummyServer{addr: "server3", isAlive: true},
	}

	lb := NewLoadBalancer("3000", servers)

	// Round-robin count starts from 0
	server1 := lb.getNextAvailableServer()
	if server1.Address() != "server1" {
		t.Errorf("Expected server1, got %s", server1.Address())
	}

	server2 := lb.getNextAvailableServer()
	if server2.Address() != "server3" {
		t.Errorf("Expected server3, got %s", server2.Address())
	}

	server3 := lb.getNextAvailableServer()
	if server3.Address() != "server1" {
		t.Errorf("Expected server1, got %s", server3.Address())
	}
}

func TestLoadBalancerServeProxy(t *testing.T) {
	servers := []Server{
		&dummyServer{addr: "server1", isAlive: true},
		&dummyServer{addr: "server2", isAlive: true},
	}

	lb := NewLoadBalancer("3000", servers)

	// Create a mock request and recorder for testing
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	lb.serveProxy(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// TODO: Add more assertions for request forwarding to the correct server
	// (This might be more complex to test due to httptest.NewRecorder limitations.)
}

// TODO: Write more test cases to cover additional scenarios, including edge cases.

func TestMain(m *testing.M) {
	// Test setup code, if any.
	// Run tests
	m.Run()
}

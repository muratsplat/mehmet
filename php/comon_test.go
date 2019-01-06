package php

import (
	"testing"
)

func TestResolveRemoteAddrIPv4(t *testing.T) {
	result := ResolveRemoteAddr("192.0.2.1:55083")
	expectedIp := "192.0.2.1"
	expectedPort := "55083"

	if result.Port != expectedPort {
		t.Fatalf("Expected: %s but got %s", expectedPort, result.Port)
	}

	if result.Ip != expectedIp {
		t.Fatalf("Expected: %s but got %s", expectedIp, result.Ip)
	}

}

func TestResolveRemoteAddr(t *testing.T) {
	result := ResolveRemoteAddr("[::1]:55083")
	expectedIp := "::1"
	expectedPort := "55083"

	if result.Port != expectedPort {
		t.Fatalf("Expected: %s but got %s", expectedPort, result.Port)
	}

	if result.Ip != expectedIp {
		t.Fatalf("Expected: %s but got %s", expectedIp, result.Ip)
	}

}

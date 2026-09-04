package main

import "testing"

func TestGetenv(t *testing.T) {
	t.Setenv("TEST_VALUE", "configured")
	if got := getenv("TEST_VALUE", "fallback"); got != "configured" {
		t.Fatalf("getenv() = %q, want configured", got)
	}
	if got := getenv("MISSING_VALUE", "fallback"); got != "fallback" {
		t.Fatalf("getenv() = %q, want fallback", got)
	}
}

package common

import "testing"

func TestAddressAndCIDRHelpers(t *testing.T) {
	if got, err := PublicValidator("203.0.113.1"); err != nil || got != "203.0.113.1" {
		t.Fatalf("PublicValidator = %q, %v", got, err)
	}
	if got, err := PrivateValidator("10.0.0.1"); err != nil || got != "10.0.0.1" {
		t.Fatalf("PrivateValidator = %q, %v", got, err)
	}
	if got, err := CIDRValidator("10.0.0.1/24"); err != nil || got != "10.0.0.0/24" {
		t.Fatalf("CIDRValidator = %q, %v", got, err)
	}
	if !IPcheck("10.0.0.5", "10.0.0.0/24") {
		t.Fatal("IPcheck rejected an address in the subnet")
	}
	if IPcheck("invalid", "invalid") {
		t.Fatal("IPcheck accepted invalid input")
	}
}

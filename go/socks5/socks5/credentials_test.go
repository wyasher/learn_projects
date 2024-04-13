package socks5

import "testing"

func TestMemoryCredentials(t *testing.T) {
	credentials := MemoryCredentials{
		"foo": "bar",
		"baz": "",
	}
	if !credentials.Valid("foo", "bar") {
		t.Fatalf("expect foo valid")
	}
	if !credentials.Valid("baz", "") {
		t.Fatalf("expect baz valid")
	}
	if credentials.Valid("ff", "baz") {
		t.Fatalf("expect ff invalid")
	}
}

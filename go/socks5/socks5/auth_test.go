package socks5

import (
	"bytes"
	"errors"
	"testing"
)

func TestNoAuth(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{1, NoAuth})
	var resp bytes.Buffer

	s, _ := New(&Config{})
	ctx, err := s.authenticate(&resp, req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ctx.Method != NoAuth {
		t.Fatal("invalid context method")
	}
	out := resp.Bytes()
	if !bytes.Equal(out, []byte{SocksVersion5, NoAuth}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestPasswordAuth_Valid(t *testing.T) {
	req := bytes.NewBuffer(nil)
	// 支持两种认证
	req.Write([]byte{2, NoAuth, UsernamePassword})
	//	+----+------+----------+------+----------+
	//	|VER | ULEN |  UNAME   | PLEN |  PASSWD  |
	//	+----+------+----------+------+----------+
	//	| 1  |  1   | 1 to 255 |  1   | 1 to 255 |
	//	+----+------+----------+------+----------+
	req.Write([]byte{1, 3, 'f', 'o', 'o', 3, 'b', 'a', 'r'})
	var resp bytes.Buffer

	cred := MemoryCredentials{
		"foo": "bar",
	}
	authenticator := UserPassAuthenticator{Credentials: cred}

	s, _ := New(&Config{AuthMethods: []Authenticator{&authenticator}})
	ctx, err := s.authenticate(&resp, req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ctx.Method != UsernamePassword {
		t.Fatal("invalid context method")
	}
	val, ok := ctx.Payload["Username"]
	if !ok {
		t.Fatal("missing key username in auth context's payload")
	}
	if val != "foo" {
		t.Fatal("invalid username in auth context's payload")
	}
	out := resp.Bytes()
	if !bytes.Equal(out, []byte{SocksVersion5, UsernamePassword, 1, authSuccess}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestUserPassAuth_Invalid(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{2, NoAuth, UsernamePassword})
	req.Write([]byte{1, 3, 'f', 'o', 'o', 3, 'b', 'a', 'z'})
	var resp bytes.Buffer

	cred := MemoryCredentials{
		"foo": "bar",
	}

	authenticator := UserPassAuthenticator{Credentials: cred}
	s, _ := New(&Config{AuthMethods: []Authenticator{&authenticator}})

	ctx, err := s.authenticate(&resp, req)
	if !errors.Is(err, UserAuthFailed) {
		t.Fatalf("err: %v", err)
	}
	if ctx != nil {
		t.Fatal("invalid context method")
	}
	out := resp.Bytes()
	if !bytes.Equal(out, []byte{SocksVersion5, UsernamePassword, 1, authFail}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestNoSupportedAuth(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{1, NoAuth})
	var resp bytes.Buffer
	cred := MemoryCredentials{
		"foo": "bar",
	}
	authenticator := UserPassAuthenticator{Credentials: cred}

	s, _ := New(&Config{AuthMethods: []Authenticator{&authenticator}})

	ctx, err := s.authenticate(&resp, req)

	if !errors.Is(err, NoSupportedAuth) {
		t.Fatalf("err: %v", err)
	}

	if ctx != nil {
		t.Fatal("Invalid Context Method")
	}

	out := resp.Bytes()
	if !bytes.Equal(out, []byte{SocksVersion5, NoAcceptable}) {
		t.Fatalf("bad: %v", out)
	}
}

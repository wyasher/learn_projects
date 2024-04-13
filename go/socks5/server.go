package main

import (
	"log"
	"os"
	"socks5/socks5"
)

func main() {
	// Create a SOCKS5 server
	cred := socks5.MemoryCredentials{
		"foo": "bar",
	}
	authenticator := socks5.UserPassAuthenticator{Credentials: cred}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{&authenticator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
		panic(err)
	}

}

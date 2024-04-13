package socks5

import (
	"context"
	"testing"
)

func TestPermitCommand(t *testing.T) {
	ctx := context.Background()
	p := PermitAll()
	if _, ok := p.Allow(ctx, &Request{Command: ConnectCommand}); !ok {
		t.Fatalf("Allow(ConnectCommand) should be true")
	}
	if _, ok := p.Allow(ctx, &Request{Command: UDPAssociateCommand}); !ok {
		t.Fatalf("Allow(UDPAssociateCommand) should be true")
	}
	if _, ok := p.Allow(ctx, &Request{Command: BindCommand}); !ok {
		t.Fatalf("Allow(BindCommand) should be true")
	}
}

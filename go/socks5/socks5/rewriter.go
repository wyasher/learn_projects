package socks5

import "context"

// AddressRewriter 用于地址重写
type AddressRewriter interface {
	Rewrite(ctx context.Context, request *Request) (context.Context, AddressSpec)
}

package socks5

import "context"

// RuleSet 自定义规则是否允许访问
type RuleSet interface {
	Allow(ctx context.Context, req *Request) (context.Context, bool)
}

func DefaultRuleSet() RuleSet {
	return &PermitCommand{
		true, false, false,
	}
}

// PermitAll 创建应该允许所有规则的类型
func PermitAll() RuleSet {
	return &PermitCommand{
		true, true, true,
	}
}

// PermitNone 创建应该拒绝所有规则的类型
func PermitNone() RuleSet {
	return &PermitCommand{
		false, false, false,
	}
}

type PermitCommand struct {
	EnableConnect      bool
	EnableBind         bool
	EnableUDPAssociate bool
}

func (p *PermitCommand) Allow(ctx context.Context, req *Request) (context.Context, bool) {
	switch req.Command {
	case ConnectCommand:
		return ctx, p.EnableConnect
	case BindCommand:
		return ctx, p.EnableBind
	case UDPAssociateCommand:
		return ctx, p.EnableUDPAssociate
	}
	return ctx, false
}

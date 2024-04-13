package spider

// RuleTree 采集规则树
type RuleTree struct {
	Root  func() ([]*Request, error) // 根节点(执行入口)
	Trunk map[string]*Rule           // 规则哈希表
}

// Rule 采集规则节点
type Rule struct {
	ItemFields []string
	ParseFunc  func(ctx *Context) (ParseResult, error)
}

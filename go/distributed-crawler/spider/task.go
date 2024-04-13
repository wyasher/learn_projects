package spider

import "sync"

type Property struct {
	Name     string `json:"name"` // 任务名称，应保证唯一性
	URL      string `json:"url"`
	Cookie   string `json:"cookie"`
	WaitTime int64  `json:"waitTime"` // 随机休眠时间，秒
	Reload   bool   `json:"reload"`   // 网站是否可以重复爬取
	MaxDepth int64  `json:"maxDepth"` // 最大深度防止无限递归
}

type LimitConfig struct {
	EventCount    int
	EventDuration int // 秒
	Bucket        int // 桶大小
}

type TaskConfig struct {
	Name     string
	Cookie   string
	WaitTime int64
	Reload   bool
	MaxDepth int64
	Fetcher  string
	Limits   []LimitConfig
}

// Task 任务实例
type Task struct {
	Visited     map[string]bool // 是否已经访问
	VisitedLock sync.Mutex

	Closed bool

	Rule RuleTree
	Options
}

func NewTask(opts ...Option) *Task {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}
	t := &Task{}
	t.Options = options
	return t
}

type Fetcher interface {
	Get(url *Request) ([]byte, error)
}

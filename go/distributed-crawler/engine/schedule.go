package engine

import (
	"crawler/spider"
	"go.uber.org/zap"
	"sync"
)

func init() {

}

type Scheduler interface {
	Schedule()
	Push(...*spider.Request)
	Pull() *spider.Request
}

type Schedule struct {
	requestCn   chan *spider.Request
	workerCn    chan *spider.Request
	priReqQueue []*spider.Request // 优先队列
	reqQueue    []*spider.Request
	Logger      *zap.Logger
}

type CrawlerStore struct {
	list []*spider.Task
	Hash map[string]*spider.Task
}

type Crawler struct {
	id          string
	out         chan spider.ParseResult
	Visited     map[string]bool
	VisitedLock sync.Mutex

	failures map[string]*spider.Request // 失败请求id -> 失败请求
	failureLock sync.Mutex

	resources map[string]*
}

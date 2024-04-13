package spider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"regexp"
	"time"
)

type Context struct {
	Body []byte
	Req  *Request
}

func (c *Context) GetRule(ruleName string) *Rule {
	return c.Req.Task.Rule.Trunk[ruleName]
}

func (c *Context) Output(data any) *DataCell {
	cell := &DataCell{
		Task: c.Req.Task,
	}
	cell.Data = make(map[string]any)
	cell.Data["Task"] = c.Req.Task.Name
	cell.Data["Rule"] = c.Req.RuleName
	cell.Data["Data"] = data
	cell.Data["URL"] = c.Req.URL
	cell.Data["Time"] = time.Now().Format("2006-01-02 15:04:05")
	return cell
}

func (c *Context) ParseJSReg(name string, reg string) ParseResult {
	re := regexp.MustCompile(reg)

	matches := re.FindAllSubmatch(c.Body, -1)
	result := ParseResult{}

	for _, m := range matches {
		u := string(m[1])
		result.Requests = append(
			result.Requests, &Request{
				Method:   "GET",
				Task:     c.Req.Task,
				URL:      u,
				Depth:    c.Req.Depth + 1,
				RuleName: name,
			})
	}

	return result
}

func (c *Context) OutputJS(reg string) ParseResult {
	re := regexp.MustCompile(reg)
	if ok := re.Match(c.Body); !ok {
		return ParseResult{
			Items: []any{},
		}
	}
	result := ParseResult{
		Items: []any{
			c.Req.URL,
		},
	}
	return result
}

type Request struct {
	Task     *Task
	URL      string
	Method   string
	Depth    int64 // 深度
	Priority int64 // 优先级
	RuleName string
	TmpData  *Temp
}

func (r *Request) Fetch() ([]byte, error) {
	if err := r.Task.Limit.Wait(context.Background()); err != nil {
		return nil, err
	}
	// 随机休眠，模拟人类行为
	sleepTime := rand.Int63n(r.Task.WaitTime * 1000)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	return r.Task.Fetcher.Get(r)
}

type ParseResult struct {
	Requests []*Request
	Items    []any
}

func (r *Request) Check() error {
	if r.Depth > r.Task.MaxDepth {
		return errors.New("max depth limit reached")
	}
	if r.Task.Closed {
		return errors.New("task has Closed")
	}
	return nil
}

func (r *Request) Unique() string {
	block := md5.Sum([]byte(r.URL + r.Method))
	return hex.EncodeToString(block[:])
}

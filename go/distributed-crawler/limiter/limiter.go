package limiter

import (
	"context"
	"golang.org/x/time/rate"
	"sort"
	"time"
)

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func Multi(limiters ...RateLimiter) *MultiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)

	return &MultiLimiter{limiters: limiters}
}

type MultiLimiter struct {
	limiters []RateLimiter
}

func (m *MultiLimiter) Wait(ctx context.Context) error {
	for _, l := range m.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (m *MultiLimiter) Limit() rate.Limit {
	return m.limiters[0].Limit()
}

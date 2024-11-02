package redis

import (
	"sync"
	"time"
)

type (
	Redis struct {
		m sync.Map
	}

	item struct {
		exp time.Time
		val any
	}
)

func NewRedis() *Redis {
	return &Redis{}
}

func (r *Redis) Get(key any) (any, bool) {
	v, ok := r.m.Load(key)
	if !ok {
		return nil, ok
	}

	if v.(item).exp.After(time.Now()) {
		return v.(item).val, ok
	}

	return nil, false
}

func (r *Redis) Set(key, value any, dur time.Duration) {
	r.m.Store(key, item{exp: time.Now().Add(dur), val: value})
}

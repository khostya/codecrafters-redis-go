package main

import "sync"

type Redis struct {
	m sync.Map
}

func NewRedis() *Redis {
	return &Redis{}
}

func (r *Redis) Get(key any) (any, bool) {
	return r.m.Load(key)
}

func (r *Redis) Set(key, value any) {
	r.m.Store(key, value)
}

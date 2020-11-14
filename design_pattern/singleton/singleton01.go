package main

import "sync"

var mu sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		defer mu.Unlock()
		mu.Lock()
		instance = &singleton{}
	}
	return instance
}

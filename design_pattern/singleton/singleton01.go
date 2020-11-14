package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

type singleton struct {
	cfg string
}

func (s *singleton) GetSysCfg() string {
	if s.cfg == "" {
		s.cfg = "System Config"
	}
	return s.cfg
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		defer mu.Unlock()
		mu.Lock()
		instance = &singleton{}
	}
	return instance
}

func main() {
	cfg := GetInstance().GetSysCfg()
	fmt.Println(cfg)
}

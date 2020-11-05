package internal

import "time"

type UTestConfig struct {
	Name string
	Func func(uint64, int) time.Duration
	Inputs interface{}
	Competitor func(uint64, int) time.Duration
}

type ITestConfig struct {
	Name string
	Func func(int64, int) time.Duration
	Inputs interface{}
	Competitor func(int64, int) time.Duration
}
